package generator

import (
	"testing"

	"github.com/moeru-ai/inventory/cmd/tools/jem-generator/types"
	"github.com/stretchr/testify/assert"
)

func TestGenerator(t *testing.T) {
	a := assert.New(t)

	t.Run("generate provider", func(t *testing.T) {
		expected := `// Auto-generated file. Do not edit.

import type { Provider } from "./types";

export const providers = [{
    name: "test",
    apiBaseURL: "https://api.test.com",
    endpoints: {"audio-music": "/audio/music","audio-speech": "/audio/speech","chat-completion": "/chat/completions","completion": "/completions","embedding": "/embeddings","image-generation": "/images/generations",},
  },] as const satisfies Provider[]
`

		generator := Generator{
			providers: []types.Provider{
				{
					Name:       "test",
					APIBaseURL: "https://api.test.com",
					Endpoints: map[types.EndpointType]string{
						types.EndpointTypeAudioMusic:      "/audio/music",
						types.EndpointTypeAudioSpeech:     "/audio/speech",
						types.EndpointTypeChatCompletion:  "/chat/completions",
						types.EndpointTypeCompletion:      "/completions",
						types.EndpointTypeEmbedding:       "/embeddings",
						types.EndpointTypeImageGeneration: "/images/generations",
					},
				},
			},
		}
		result, err := generator.generateProviders()
		a.NoError(err)
		a.Equal(expected, string(result))
	})

	t.Run("generate models", func(t *testing.T) {
		expected := `// Auto-generated file. Do not edit.

import type { Model } from "./types";

export const models = [{
    modelId: "test",
    provider: "test",
    endpoints: ["audio-music","audio-speech","chat-completion","completion","embedding","image-generation",],
    capabilities: ["reasoning","streaming","tool-call",],
    inputModalities: ["audio","image","text","video","vector",],
    outputModalities: ["audio","image","text","video","vector",],
  },] as const satisfies Model[]
`

		generator := Generator{
			models: []types.Model{
				{
					ModelID:  "test",
					Provider: "test",
					Endpoints: []types.EndpointType{
						types.EndpointTypeAudioMusic,
						types.EndpointTypeAudioSpeech,
						types.EndpointTypeChatCompletion,
						types.EndpointTypeCompletion,
						types.EndpointTypeEmbedding,
						types.EndpointTypeImageGeneration,
					},
					Capabilities: []types.Capability{
						types.CapabilityReasoning,
						types.CapabilityStreaming,
						types.CapabilityToolCall,
					},
					InputModalities: []types.Modality{
						types.ModalityAudio,
						types.ModalityImage,
						types.ModalityText,
						types.ModalityVideo,
						types.ModalityVector,
					},
					OutputModalities: []types.Modality{
						types.ModalityAudio,
						types.ModalityImage,
						types.ModalityText,
						types.ModalityVideo,
						types.ModalityVector,
					},
				},
			},
		}
		result, err := generator.generateModels()
		a.NoError(err)
		a.Equal(expected, string(result))
	})
}
