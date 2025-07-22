// Auto-generated file. Do not edit.

import type { Model } from './types.ts'

export const models = [
  {
    modelId: 'MiniMax-Text-01',
    provider: 'minimax',
    endpoints: [
      'chat-completion',
    ],
    capabilities: [
      'tool-call',
      'streaming',
    ],
    inputModalities: [
      'text',
    ],
    outputModalities: [
      'text',
    ],
  },
  {
    modelId: 'MiniMax-Text-01',
    provider: 'minimaxi',
    endpoints: [
      'chat-completion',
    ],
    capabilities: [
      'tool-call',
      'streaming',
    ],
    inputModalities: [
      'text',
    ],
    outputModalities: [
      'text',
    ],
  },
  {
    modelId: 'gpt-4o',
    provider: 'openai',
    endpoints: [
      'chat-completion',
    ],
    capabilities: [
      'tool-call',
      'streaming',
    ],
    inputModalities: [
      'text',
    ],
    outputModalities: [
      'text',
    ],
  },
  {
    capabilities: [],
    endpoints: [],
    inputModalities: [
      'text',
    ],
    outputModalities: [
      'audio',
    ],
    modelId: 'gemini-2.5-flash-preview-tts',
    provider: 'google',
  },
] as const satisfies Model[]
