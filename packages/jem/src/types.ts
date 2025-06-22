export interface Provider {
  name: string
  apiBaseURL: string
  endpoints: Record<string, string>
}

export interface Model {
  capabilities: string[]
  endpoints: string[]
  inputModalities: string[]
  modelId: string
  outputModalities: string[]
  provider: string
}
