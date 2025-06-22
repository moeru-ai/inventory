// Auto-generated file. Do not edit.

import type { Model } from './types'

export const models = [{
  modelId: 'MiniMax-Text-01',
  provider: 'minimax',
  endpoints: ['chat-completion'],
  capabilities: ['tool-call', 'streaming'],
  inputModalities: ['text'],
  outputModalities: ['text'],
}, {
  modelId: 'gpt-4o',
  provider: 'openai',
  endpoints: ['chat-completion'],
  capabilities: ['tool-call', 'streaming'],
  inputModalities: ['text'],
  outputModalities: ['text'],
}] as const satisfies Model[]
