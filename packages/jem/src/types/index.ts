export interface Provider {
  name: string;
  apiBaseURL: string;
  endpoints: Record<string, string>;
  models: Model[];
}

export interface Model {
  capabilities: string[];
  endpoints: string[];
  inputModalities: string[];
  modelId: string;
  outputModalities: string[];
  provider: string;
}

// Enums, for type safety and autocomplete

export enum Capability {
  ToolCall = "tool-call",
  Reasoning = "reasoning",
  Streaming = "streaming",
}

export enum Modality {
  Text = "text",
  Image = "image",
  Audio = "audio",
  Video = "video",
  Vector = "vector",
}

export enum EndpointType {
  ChatCompletion = "chat-completion",
  Completion = "completion",
  Embedding = "embedding",
  ImageGeneration = "image-generation",
  AudioSpeech = "audio-speech",
  AudioMusic = "audio-music",
}
