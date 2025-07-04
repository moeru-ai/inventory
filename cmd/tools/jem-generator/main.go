package main

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"net/http"
	"os"
	"slices"

	"github.com/moeru-ai/inventory/cmd/tools/jem-generator/generator"
	"github.com/moeru-ai/inventory/cmd/tools/jem-generator/providers"
	"github.com/moeru-ai/inventory/cmd/tools/jem-generator/types"
	"github.com/nekomeowww/xo/logger"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"
)

var apiKeyMap = map[string]string{
	providers.ProviderNameOpenai:  os.Getenv("OPENAI_API_KEY"),
	providers.ProviderNameMinimax: os.Getenv("MINIMAX_API_KEY"),
}

var errorGroup = new(errgroup.Group)

var mainLogger, _ = logger.NewLogger(
	logger.WithLevel(zapcore.DebugLevel),
	logger.WithCallFrameSkip(1),
	logger.WithNamespace("jem-generator"),
)

const fileMode = 0755

func createRequestBody(model types.Model) map[string]any {
	requestBody := map[string]any{
		"model": model.ModelID,
		"messages": []map[string]any{
			{
				"role":    "system",
				"content": "You are a cute catgirl.",
			},
			{
				"role":    "user",
				"content": "Hello, how are you?",
			},
		},
	}
	if slices.Index(model.Capabilities, types.CapabilityStreaming) != -1 {
		requestBody["stream"] = true
	}

	return requestBody
}

func checkChatCompletion(model types.Model, provider types.Provider, ctx context.Context) {
	mainLogger.With(
		zap.String("model", model.ModelID),
		zap.String("provider", provider.Name),
	).Info("Detecting chat completion.")

	apiKey := apiKeyMap[provider.Name]

	requestBody := createRequestBody(model)

	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		mainLogger.With(
			zap.String("model", model.ModelID),
			zap.String("provider", provider.Name),
			zap.Error(err),
		).Error("Error marshalling request body.")

		return
	}

	request, err := http.NewRequest(
		http.MethodPost,
		provider.APIBaseURL+provider.Endpoints[types.EndpointTypeChatCompletion],
		bytes.NewBuffer(requestBodyBytes),
	)
	if err != nil {
		mainLogger.With(
			zap.String("model", model.ModelID),
			zap.String("provider", provider.Name),
			zap.Error(err),
		).Error("Error creating request.")

		return
	}

	request = request.WithContext(ctx)
	request.Header.Set("Authorization", "Bearer "+apiKey)
	request.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		mainLogger.With(
			zap.String("model", model.ModelID),
			zap.String("provider", provider.Name),
		).Error("Error sending request.",
			zap.Error(err),
		)

		return
	}
	defer response.Body.Close()

	err = provider.ParseResponseFunc[types.EndpointTypeChatCompletion](response, requestBody)
	if err != nil {
		mainLogger.With(
			zap.String("model", model.ModelID),
			zap.String("provider", provider.Name),
			zap.Int("status_code", response.StatusCode),
			zap.Error(err),
		).Error("Error checking chat completion.")

		return
	}

	mainLogger.With(
		zap.String("model", model.ModelID),
		zap.String("provider", provider.Name),
	).Info("Done.")
}

var endpointsDetectionMap = map[types.EndpointType]func(model types.Model, provider types.Provider, ctx context.Context){
	types.EndpointTypeChatCompletion: checkChatCompletion,
}

func detectEndpoints(model types.Model, provider types.Provider, ctx context.Context) {
	for _, endpoint := range model.Endpoints {
		fn := endpointsDetectionMap[endpoint]

		errorGroup.Go(func() error {
			fn(model, provider, ctx)
			return nil
		})
	}
}

func main() {
	// load configurations
	for _, provider := range providers.Providers {
		apiKey := apiKeyMap[provider.Name]
		if apiKey == "" {
			mainLogger.With(
				zap.String("provider", provider.Name),
			).Fatal("No API key found for provider.")
		}
	}

	outputDir := os.Getenv("OUTPUT_DIR")
	if outputDir == "" {
		mainLogger.Error("Output directory is not set.")
		return
	}

	err := os.MkdirAll(outputDir, fileMode)
	if err != nil {
		mainLogger.With(
			zap.String("output_dir", outputDir),
		).Fatal("Error creating output directory.", zap.Error(err))
	}

	// detect capabilities
	for _, model := range providers.Models {
		provider, ok := lo.Find(providers.Providers, func(item types.Provider) bool {
			return item.Name == model.Provider
		})
		if !ok {
			mainLogger.With(
				zap.String("model", model.ModelID),
			).Fatal("Provider not found.")
		}

		detectEndpoints(model, provider, context.TODO())
	}

	err = errorGroup.Wait()
	if err != nil {
		mainLogger.Fatal("Error detecting endpoints. Stopping...", zap.Error(err))
	}

	// generate
	generator, err := generator.New(providers.Providers, providers.Models, outputDir)
	if err != nil {
		mainLogger.Fatal("Error creating generator.", zap.Error(err))
	}

	err = generator.Generate()
	if err != nil {
		mainLogger.Fatal("Error generating.", zap.Error(err))
	}

	mainLogger.Info("Done.")
}
