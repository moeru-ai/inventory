package providers

import "github.com/moeru-ai/inventory/cmd/tools/jem-generator/types"

var ProviderOpenAI = types.Provider{
	APIBaseURL: "https://api.openai.com/v1",
	Name:       "openai",
	Endpoints: map[types.EndpointType]string{
		types.EndpointTypeChatCompletion: "/chat/completions",
	},
}

var ModelOpenAI = []types.Model{
	{
		ModelID:  "gpt-4o",
		Provider: ProviderOpenAI.Name,
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
