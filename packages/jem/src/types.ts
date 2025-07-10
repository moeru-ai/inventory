export interface Provider<T extends string[]> {
  name: string
  apiBaseURL: string | ((parameters: { [key in T[number]]: string }) => string)
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
