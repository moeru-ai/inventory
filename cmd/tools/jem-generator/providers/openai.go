package providers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/moeru-ai/inventory/cmd/tools/jem-generator/types"
	"github.com/nekomeowww/xo/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var openaiLogger, _ = logger.NewLogger(
	logger.WithLevel(zapcore.DebugLevel),
	logger.WithCallFrameSkip(1),
	logger.WithNamespace("openai"),
)

const ProviderNameOpenai = "openai"

var providerOpenai = types.Provider{
	APIBaseURL: "https://api.openai.com/v1",
	Name:       ProviderNameOpenai,
	Endpoints: map[types.EndpointType]string{
		types.EndpointTypeChatCompletion: "/chat/completions",
	},
	ParseResponseFunc: map[types.EndpointType]types.ParseResponseFunc{
		types.EndpointTypeChatCompletion: func(resp *http.Response, requestBody map[string]any) error {
			if resp.StatusCode != http.StatusOK {
				return fmt.Errorf("status code: %d", resp.StatusCode)
			}

			responseBody, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			responseBodyStr := string(responseBody)
			openaiLogger.Debug("response body", zap.String("response_body", responseBodyStr))

			return nil
		},
	},
}

var modelsOpenai = []types.Model{
	{
		ModelID:  "gpt-4o",
		Provider: ProviderNameOpenai,
		Endpoints: []types.EndpointType{
			types.EndpointTypeChatCompletion,
		},
		Capabilities: []types.Capability{
			types.CapabilityToolCall,
			types.CapabilityStreaming,
		},
		InputModalities: []types.Modality{
			types.ModalityText,
		},
		OutputModalities: []types.Modality{
			types.ModalityText,
		},
	},
}
