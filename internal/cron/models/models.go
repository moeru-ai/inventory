package models

import "time"

type Model struct {
	ID              string        `json:"id"`
	Provider        string        `json:"provider"`
	Dimensions      int           `json:"dimensions"`
	ResponseTimeP99 time.Duration `json:"response_time_p99"`
}

var (
	commonTasksEmbeddingModels = []*Model{
		{
			ID:              "voyage-3-large",
			Provider:        "Voyage",
			Dimensions:      0,
			ResponseTimeP99: 0,
		},
	}
)
