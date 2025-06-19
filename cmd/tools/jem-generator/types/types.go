package types

import "net/http"

type Capability string // enum

const (
	// LLM capabilities
	CapabilityToolCall  Capability = "tool-call"
	CapabilityReasoning Capability = "reasoning"
	CapabilityStreaming Capability = "streaming"
)

type Modality string // enum

const (
	ModalityText   Modality = "text"
	ModalityAudio  Modality = "audio"
	ModalityImage  Modality = "image"
	ModalityVideo  Modality = "video"
	ModalityVector Modality = "vector"
)

type EndpointType string // enum

const (
	EndpointTypeChatCompletion  EndpointType = "chat-completion"
	EndpointTypeCompletion      EndpointType = "completion"
	EndpointTypeEmbedding       EndpointType = "embedding"
	EndpointTypeImageGeneration EndpointType = "image-generation"
	EndpointTypeAudioSpeech     EndpointType = "audio-speech"
	EndpointTypeAudioMusic      EndpointType = "audio-music"
)

type Provider struct {
	APIBaseURL        string                             `json:"api_base_url"`
	Name              string                             `json:"name"`
	Endpoints         map[EndpointType]string            `json:"endpoints"`
	ParseResponseFunc map[EndpointType]ParseResponseFunc `json:"parse_response_func"`
	Models            []Model                            `json:"models"`
}

type Model struct {
	Capabilities     []Capability   `json:"capabilities"`
	Endpoints        []EndpointType `json:"endpoints"`
	InputModalities  []Modality     `json:"input_modalities"`
	ModelID          string         `json:"model_id"`
	OutputModalities []Modality     `json:"output_modalities"`
	Provider         string         `json:"provider"`
}

type Catalog struct {
	GeneratedAt string     `json:"generated_at"` // ISO 8601 timestamp
	Models      []Model    `json:"models"`
	Providers   []Provider `json:"providers"`
	Version     string     `json:"version"` // semantic version
}

type ParseResponseFunc func(resp *http.Response, requestBody map[string]any) error
