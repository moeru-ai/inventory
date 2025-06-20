package main

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"net/http"
	"os"
	"slices"
	"sync"

	"github.com/moeru-ai/inventory/cmd/tools/jem-generator/generator"
	"github.com/moeru-ai/inventory/cmd/tools/jem-generator/providers/minimax"
	"github.com/moeru-ai/inventory/cmd/tools/jem-generator/providers/openai"
	"github.com/moeru-ai/inventory/cmd/tools/jem-generator/types"
	"github.com/nekomeowww/xo/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var providers = []types.Provider{
	openai.Provider,
	minimax.Provider,
}

var apiKeyMap = map[string]string{
	openai.Provider.Name:  os.Getenv("OPENAI_API_KEY"),
	minimax.Provider.Name: os.Getenv("MINIMAX_API_KEY"),
}

var wg = sync.WaitGroup{}

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
	defer wg.Done()

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
		wg.Add(1)

		fn := endpointsDetectionMap[endpoint]
		go fn(model, provider, ctx)
	}
}

func main() {
	// load configurations
	for _, provider := range providers {
		apiKey := apiKeyMap[provider.Name]
		if apiKey == "" {
			mainLogger.With(
				zap.String("provider", provider.Name),
			).Error("No API key found for provider.")

			return
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
		).Error("Error creating output directory.", zap.Error(err))

		return
	}

	// detect capabilities

	for _, provider := range providers {
		for _, model := range provider.Models {
			detectEndpoints(model, provider, context.TODO())
		}
	}

	wg.Wait()

	// generate
	generator, err := generator.New(providers, outputDir)
	if err != nil {
		mainLogger.Error("Error creating generator.", zap.Error(err))
		return
	}

	err = generator.Generate()
	if err != nil {
		mainLogger.Error("Error generating.", zap.Error(err))
		return
	}

	mainLogger.Info("Done.")
}
