package models

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	v1 "github.com/moeru-ai/inventory/apis/inventoryapi/v1"
	"github.com/moeru-ai/unspeech/pkg/utils"
	"github.com/samber/lo"
)

func EmbeddingVoyageAIVoyage3Large() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()

	body := lo.Must(json.Marshal(map[string]any{
		"model": "voyage-3-large",
		"input": "Hello, world!",
	}))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.voyageai.com/v1/embeddings", bytes.NewBuffer(body))
	if err != nil {
		slog.Error("failed", slog.Any("error", err))
		return
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("VOYAGE_API_KEY"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error("failed", slog.Any("error", err))
		return
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	if res.StatusCode >= 400 && res.StatusCode <= 599 {
		switch {
		case strings.HasPrefix(res.Header.Get("Content-Type"), "application/json"):
			slog.Error("failed", slog.Any("error", utils.NewJSONResponseError(res.StatusCode, res.Body).OrEmpty().Error()))
		case strings.HasPrefix(res.Header.Get("Content-Type"), "text/"):
			slog.Error("failed", slog.Any("error", utils.NewTextResponseError(res.StatusCode, res.Body).OrEmpty().Error()))
		default:
			slog.Warn("unknown upstream error with unknown Content-Type",
				slog.Int("status", res.StatusCode),
				slog.String("content_type", res.Header.Get("Content-Type")),
				slog.String("content_length", res.Header.Get("Content-Length")),
			)

			slog.Error("failed", slog.Any("error", res.Status))
		}

		return
	}

	var m map[string]any

	if json.NewDecoder(res.Body).Decode(&m) != nil {
		slog.Error("failed", slog.Any("error", err))
		return
	}

	embedding := utils.GetByJSONPath[[]float64](m, "{ .data[0].embedding }")

	model, ok := lo.Find(commonTasksEmbeddingModels, func(item *v1.GetModelsModelItem) bool {
		return item.GetId() == "voyage-3-large"
	})
	if ok {
		model.GetEmbedding().Dimensions = int64(len(embedding))
	} else {
		commonTasksEmbeddingModels = append(commonTasksEmbeddingModels, &v1.GetModelsModelItem{
			Id:                  "voyage-3-large",
			Name:                "Voyage 3 Large",
			Description:         "",
			Deprecated:          false,
			ProviderId:          "voyage.ai",
			ProviderName:        "Voyage",
			ProviderDescription: "",
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
}
