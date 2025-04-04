package models

import (
	v1 "github.com/moeru-ai/inventory/apis/inventoryapi/v1"
)

var (
	commonTasksEmbeddingModels = []*v1.GetModelsModelItem{
		{
			Id:           "voyage-3-large",
			ProviderId:   "voyage.ai",
			ProviderName: "Voyage",
			ModelType: &v1.GetModelsModelItem_Embedding{
				Embedding: &v1.GetModelsModelItemEmbedding{
					Dimensions: 0,
				},
			},
		},
	}
)

func Models() []*v1.GetModelsModelItem {
	return commonTasksEmbeddingModels
}
