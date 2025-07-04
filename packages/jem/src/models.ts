// Auto-generated file. Do not edit.

  import type { Model } from './types'

  export const models = [
  {
    "modelId": "MiniMax-Text-01",
    "provider": "minimax",
    "endpoints": [
      "chat-completion"
    ],
    "capabilities": [
      "tool-call",
      "streaming"
    ],
    "inputModalities": [
      "text"
    ],
    "outputModalities": [
      "text"
    ]
  },
  {
    "capabilities": [],
    "endpoints": [
      "image-generation"
    ],
    "inputModalities": [
      "text",
      "image"
    ],
    "outputModalities": [
      "image"
    ],
    "modelId": "gpt-image-1",
    "provider": "openai"
  }
] as const satisfies Model[]
  