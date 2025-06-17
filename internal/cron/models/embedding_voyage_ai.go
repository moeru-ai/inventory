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

type VoyageAIModel struct {
	ID          string
	Name        string
	Description string
	MaxTokens   int
}

var voyageAIModels = []VoyageAIModel{
	{
		ID:          "voyage-3.5-lite",
		Name:        "Voyage 3.5 Lite",
		Description: "Lightweight embedding model",
		// https://docs.voyageai.com/reference/embeddings-api
		MaxTokens: 1_000_000, //nolint:mnd
	},
	{
		ID:          "voyage-3.5",
		Name:        "Voyage 3.5",
		Description: "Standard embedding model",
		// https://docs.voyageai.com/reference/embeddings-api
		MaxTokens: 320_000, //nolint:mnd
	},
	{
		ID:          "voyage-2",
		Name:        "Voyage 2",
		Description: "Standard embedding model",
		// https://docs.voyageai.com/reference/embeddings-api
		MaxTokens: 320_000, //nolint:mnd
	},
	{
		ID:          "voyage-3-large",
		Name:        "Voyage 3 Large",
		Description: "Large embedding model",
		// https://docs.voyageai.com/reference/embeddings-api
		MaxTokens: 120_000, //nolint:mnd
	},
	{
		ID:          "voyage-code-3",
		Name:        "Voyage Code 3",
		Description: "Code-specialized embedding model",
		// https://docs.voyageai.com/reference/embeddings-api
		MaxTokens: 120_000, //nolint:mnd
	},
	{
		ID:          "voyage-finance-2",
		Name:        "Voyage Finance 2",
		Description: "Finance-specialized embedding model",
		// https://docs.voyageai.com/reference/embeddings-api
		MaxTokens: 120_000, //nolint:mnd
	},
	{
		ID:          "voyage-law-2",
		Name:        "Voyage Law 2",
		Description: "Law-specialized embedding model",
		// https://docs.voyageai.com/reference/embeddings-api
		MaxTokens: 120_000, //nolint:mnd
	},
}

func getVoyageAIModel(modelID string) (VoyageAIModel, bool) {
	return lo.Find(voyageAIModels, func(m VoyageAIModel) bool {
		return m.ID == modelID
	})
}

func EmbeddingVoyageAI() {
	for _, model := range voyageAIModels {
		err := embeddingVoyageAI(model.ID, "Hello, world!")
		if err != nil {
			slog.Error("failed to embed model", slog.Any("error", err))
		}
	}
}

func embeddingVoyageAI(modelID string, input string) error {
	model, found := getVoyageAIModel(modelID)
	if !found {
		return fmt.Errorf("unsupported model ID: %s", modelID)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()

	body := lo.Must(json.Marshal(map[string]any{
		"model": model.ID,
		"input": input,
	}))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.voyageai.com/v1/embeddings", bytes.NewBuffer(body))
	if err != nil {
		slog.Error("failed to create request", slog.Any("error", err))
		return err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("VOYAGE_API_KEY"))

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

	existingModel, ok := commonTasks.Find(model.ID)
	if ok {
		existingModel.GetEmbedding().Dimensions = int64(len(embedding))
	} else {
		commonTasks.Add(&v1.GetModelsModelItem{
			Id:           model.ID,
			Name:         model.Name,
			Description:  model.Description,
			Deprecated:   false,
			ProviderId:   "voyage.ai",
			ProviderName: "Voyage",
			ProviderDescription: "" +
				"Voyage AI delivers state-of-the-art embedding and reranking models that " +
				"supercharge intelligent retrieval for enterprises, driving forward " +
				"retrieval-augmented generation (RAG) and reliable LLM applications. " +
				"Our solutions are designed to optimize the way businesses access and " +
				"utilize information, making retrieval faster, more accurate, and " +
				"scalable." +
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
