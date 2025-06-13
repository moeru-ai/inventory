package models

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	v1 "github.com/moeru-ai/inventory/apis/inventoryapi/v1"
	"github.com/moeru-ai/unspeech/pkg/utils"
	"github.com/samber/lo"
)

type OpenAIModel struct {
	ID          string
	Name        string
	Description string
	MaxTokens   int
}

var openAIModels = []OpenAIModel{
	{
		ID:          "text-embedding-3-small",
		Name:        "Text Embedding 3 Small",
		Description: "Most capable small embedding model for both english and non-english tasks",
		// https://platform.openai.com/docs/guides/embeddings
		MaxTokens: 8_191, //nolint:mnd
	},
	{
		ID:          "text-embedding-3-large",
		Name:        "Text Embedding 3 Large",
		Description: "Most capable large embedding model for both english and non-english tasks",
		// https://platform.openai.com/docs/guides/embeddings
		MaxTokens: 8_191, //nolint:mnd
	},
	{
		ID:          "text-embedding-ada-002",
		Name:        "Text Embedding Ada 002",
		Description: "Previous generation embedding model",
		// https://platform.openai.com/docs/guides/embeddings
		MaxTokens: 8_191, //nolint:mnd
	},
}

func getOpenAIModel(modelID string) (OpenAIModel, bool) {
	return lo.Find(openAIModels, func(m OpenAIModel) bool {
		return m.ID == modelID
	})
}

func EmbeddingOpenAI() {
	for _, model := range openAIModels {
		err := embeddingOpenAI(model.ID, "Hello, world!")
		if err != nil {
			slog.Error("failed to embed model", slog.Any("error", err))
		}
	}
}

func embeddingOpenAI(modelID string, input string) error {
	model, found := getOpenAIModel(modelID)
	if !found {
		return fmt.Errorf("unsupported model ID: %s", modelID)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()

	body := lo.Must(json.Marshal(map[string]any{
		"model": model.ID,
		"input": input,
	}))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.openai.com/v1/embeddings", bytes.NewBuffer(body))
	if err != nil {
		slog.Error("failed to create request", slog.Any("error", err))
		return err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error("failed to make request", slog.Any("error", err))
		return err
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	if res.StatusCode >= 400 && res.StatusCode <= 599 {
		switch {
		case strings.HasPrefix(res.Header.Get("Content-Type"), "application/json"):
			err := utils.NewJSONResponseError(res.StatusCode, res.Body).OrEmpty()
			slog.Error("failed with JSON error", slog.Any("error", err))

			return err
		case strings.HasPrefix(res.Header.Get("Content-Type"), "text/"):
			err := utils.NewTextResponseError(res.StatusCode, res.Body).OrEmpty()
			slog.Error("failed with text error", slog.Any("error", err))

			return err
		default:
			slog.Warn("unknown upstream error with unknown Content-Type",
				slog.Int("status", res.StatusCode),
				slog.String("content_type", res.Header.Get("Content-Type")),
				slog.String("content_length", res.Header.Get("Content-Length")),
			)

			return fmt.Errorf("unknown error: %s", res.Status)
		}
	}

	var m map[string]any
	if err := json.NewDecoder(res.Body).Decode(&m); err != nil {
		slog.Error("failed to decode response", slog.Any("error", err))
		return err
	}

	embedding := utils.GetByJSONPath[[]float64](m, "{ .data[0].embedding }")

	existingModel, ok := lo.Find(commonTasksEmbeddingModels, func(item *v1.GetModelsModelItem) bool {
		return item.GetId() == model.ID
	})
	if ok {
		existingModel.GetEmbedding().Dimensions = int64(len(embedding))
	} else {
		commonTasksEmbeddingModels = append(commonTasksEmbeddingModels, &v1.GetModelsModelItem{
			Id:           model.ID,
			Name:         model.Name,
			Description:  model.Description,
			Deprecated:   false,
			ProviderId:   "openai.com",
			ProviderName: "OpenAI",
			ProviderDescription: "" +
				"OpenAI is an AI research and deployment company. Our mission is to " +
				"ensure that artificial general intelligence benefits all of " +
				"humanity." +
				"",
			Provider: &v1.GetModelsModelItem_Cloud{
				Cloud: &v1.GetModelsModelItemProviderCloud{},
			},
			ModelType: &v1.GetModelsModelItem_Embedding{
				Embedding: &v1.GetModelsModelItemEmbedding{
					Dimensions: int64(len(embedding)),
				},
			},
		})
	}

	return nil
}
