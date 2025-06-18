package minimax

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/moeru-ai/inventory/cmd/tools/jem-generator/types"
	"github.com/nekomeowww/xo/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var miniMaxLogger, _ = logger.NewLogger(
	logger.WithLevel(zapcore.DebugLevel),
	logger.WithCallFrameSkip(1),
	logger.WithNamespace("minimax"),
)

type ChatCompletionResponse struct {
	BaseResp struct {
		StatusCode int    `json:"status_code"`
		StatusMsg  string `json:"status_msg"`
	} `json:"base_resp"`
}

func checkChatCompletionResponse(responseBodyStr string) error {
	var parsedResponse ChatCompletionResponse
	if err := json.Unmarshal([]byte(responseBodyStr), &parsedResponse); err != nil {
		return err
	}

	if parsedResponse.BaseResp.StatusCode != 0 {
		return fmt.Errorf("status code: %d, status msg: %s", parsedResponse.BaseResp.StatusCode, parsedResponse.BaseResp.StatusMsg)
	}

	return nil
}

var Provider = types.Provider{
	APIBaseURL: "https://api.minimaxi.com/v1",
	Name:       "minimax",
	Endpoints: map[types.EndpointType]string{
		types.EndpointTypeChatCompletion: "/text/chatcompletion_v2",
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
			miniMaxLogger.Debug("response body", zap.String("response_body", responseBodyStr))
			if requestBody["stream"] == true {
				responseSlices := strings.Split(responseBodyStr, "data: ")
				for _, responseSlice := range responseSlices {
					if responseSlice == "" {
						continue
					}
					err := checkChatCompletionResponse(responseSlice)
					if err != nil {
						return err
					}
				}
				return nil
			}

			err = checkChatCompletionResponse(responseBodyStr)
			if err != nil {
				return err
			}

			return nil
		},
	},
}

var Models = []types.Model{
	{
		ModelID:  "MiniMax-Text-01",
		Provider: Provider.Name,
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
