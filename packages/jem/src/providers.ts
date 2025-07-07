// Auto-generated file. Do not edit.

import type { Provider } from './types.ts'

export const providers = [{
  name: 'openai',
  apiBaseURL: 'https://api.openai.com/v1',
  endpoints: { 'chat-completion': '/chat/completions' },
}, {
  name: 'minimax',
  apiBaseURL: 'https://api.minimaxi.com/v1',
  endpoints: { 'chat-completion': '/text/chatcompletion_v2' },
}] as const satisfies Provider[]
