package main

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"sync"

	"github.com/moeru-ai/inventory/cmd/tools/jem-generator/providers/minimax"
	"github.com/moeru-ai/inventory/cmd/tools/jem-generator/providers/openai"
	"github.com/moeru-ai/inventory/cmd/tools/jem-generator/types"
	"github.com/nekomeowww/xo/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var providersMap = map[string]types.Provider{
	openai.Provider.Name:  openai.Provider,
	minimax.Provider.Name: minimax.Provider,
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

var (
	//go:embed templates/providers.tmpl
	templateStrProviders string
	//go:embed templates/index.tmpl
	templateStrIndex string
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

func detectEndpoints(model types.Model, ctx context.Context) {
	provider := providersMap[model.Provider]

	for _, endpoint := range model.Endpoints {
		wg.Add(1)

		fn := endpointsDetectionMap[endpoint]
		go fn(model, provider, ctx)
	}
}

func main() {
	// load configurations
	for _, provider := range providersMap {
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

	for _, provider := range providersMap {
		for _, model := range provider.Models {
			detectEndpoints(model, context.TODO())
		}
	}

	wg.Wait()

	// generate npm package
	templateProviders := template.Must(template.New("providers").Parse(templateStrProviders))

	for _, provider := range providersMap {
		fileContent := bytes.NewBuffer(nil)
		err := templateProviders.Execute(fileContent, provider)

		if err != nil {
			mainLogger.With(
				zap.String("provider", provider.Name),
			).Error("Error executing template.", zap.Error(err))

			return
		}

		var providerFilePath string
		if filepath.IsAbs(outputDir) {
			providerFilePath = filepath.Join(outputDir, provider.Name+".ts")
		} else {
			wd, err := os.Getwd()
			if err != nil {
				mainLogger.Error("Error getting working directory.", zap.Error(err))
				return
			}

			providerFilePath = filepath.Join(wd, outputDir, provider.Name+".ts")
		}

		err = os.WriteFile(providerFilePath, fileContent.Bytes(), fileMode)

		if err != nil {
			mainLogger.With(
				zap.String("file_path", providerFilePath),
			).Error("Error writing file.", zap.Error(err))

			return
		}

		mainLogger.With(
			zap.String("file_path", providerFilePath),
		).Info("File written.")
	}
}
