package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"slices"
	"sync"

	"github.com/moeru-ai/inventory/cmd/tools/jem-generator/providers"
	"github.com/moeru-ai/inventory/cmd/tools/jem-generator/types"
	"github.com/nekomeowww/xo/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var providersMap = map[string]types.Provider{
	providers.ProviderOpenAI.Name: providers.ProviderOpenAI,
}

var apiKeyMap = map[string]string{
	providers.ProviderOpenAI.Name: os.Getenv("OPENAI_API_KEY"),
}

var wg = sync.WaitGroup{}

var mainLogger, _ = logger.NewLogger(
	logger.WithLevel(zapcore.InfoLevel),
	logger.WithCallFrameSkip(1),
	logger.WithNamespace("jem-generator"),
)

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

	mainLogger.With(
		zap.String("model", model.ModelID),
		zap.String("provider", provider.Name),
		zap.Int("status_code", response.StatusCode),
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

	// detect capabilities
	models := providers.ModelOpenAI

	for _, model := range models {
		detectEndpoints(model, context.TODO())
	}

	// TODO: generate npm package

	wg.Wait()
}
