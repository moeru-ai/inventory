// Auto-generated file. Do not edit.

import type { Provider } from './types.ts'

export const providers = [
  { name: 'anthropic', apiBaseURL: 'https://api.anthropic.com/v1/', endpoints: { 'chat-completion': '/chat/completions' } },
  { name: 'azure', apiBaseURL: ({ resourceName }) => `https://${resourceName}.services.ai.azure.com/models/`, endpoints: { 'chat-completion': '/chat/completions' } },
  { name: 'cerebras', apiBaseURL: 'https://api.cerebras.ai/v1/', endpoints: { 'chat-completion': '/chat/completions' } },
  { name: 'deepinfra', apiBaseURL: 'https://api.deepinfra.com/v1/openai/', endpoints: { 'chat-completion': '/chat/completions' } },
  { name: 'deepseek', apiBaseURL: 'https://api.deepseek.com/', endpoints: { 'chat-completion': '/chat/completions' } },
  { name: 'fatherless-ai', apiBaseURL: 'https://api.featherless.ai/v1/', endpoints: { 'chat-completion': '/chat/completions' } },
  { name: 'fireworks', apiBaseURL: 'https://api.fireworks.ai/inference/v1/', endpoints: { 'chat-completion': '/chat/completions' } },
  { name: 'google-generative-ai', apiBaseURL: 'https://generativelanguage.googleapis.com/v1beta/openai/', endpoints: { 'chat-completion': '/chat/completions' } },
  { name: 'groq', apiBaseURL: 'https://api.groq.com/openai/v1/', endpoints: { 'chat-completion': '/chat/completions' } },
  { name: 'minimax', apiBaseURL: 'https://api.minimaxi.chat/v1/', endpoints: { 'chat-completion': '/chat/completions' } },
  { name: 'mistral', apiBaseURL: 'https://api.mistral.ai/v1/', endpoints: { 'chat-completion': '/chat/completions' } },
  { name: 'moonshot', apiBaseURL: 'https://api.moonshot.cn/v1/', endpoints: { 'chat-completion': '/chat/completions' } },
  { name: 'novita', apiBaseURL: 'https://api.novita.ai/v3/openai/', endpoints: { 'chat-completion': '/chat/completions' } },
  { name: 'openai', apiBaseURL: 'https://api.openai.com/v1/', endpoints: { 'chat-completion': '/chat/completions' } },
  { name: 'openrouter', apiBaseURL: 'https://openrouter.ai/api/v1/', endpoints: { 'chat-completion': '/chat/completions' } },
  { name: 'perplexity', apiBaseURL: 'https://api.perplexity.ai/', endpoints: { 'chat-completion': '/chat/completions' } },
  { name: 'qwen', apiBaseURL: 'https://dashscope.aliyuncs.com/compatible-mode/v1/', endpoints: { 'chat-completion': '/chat/completions' } },
  { name: 'silicon-flow', apiBaseURL: 'https://api.siliconflow.cn/v1/', endpoints: { 'chat-completion': '/chat/completions' } },
  { name: 'stepfun', apiBaseURL: 'https://api.stepfun.com/v1/', endpoints: { 'chat-completion': '/chat/completions' } },
  { name: 'tencent-hunyuan', apiBaseURL: 'https://api.hunyuan.cloud.tencent.com/v1/', endpoints: { 'chat-completion': '/chat/completions' } },
  { name: 'together-ai', apiBaseURL: 'https://api.together.xyz/v1/', endpoints: { 'chat-completion': '/chat/completions' } },
  { name: 'workers-ai', apiBaseURL: ({ accountId }) => `https://api.cloudflare.com/client/v4/accounts/${accountId}/ai/v1/`, endpoints: { 'chat-completion': '/chat/completions' } },
  { name: 'xai', apiBaseURL: 'https://api.x.ai/v1/', endpoints: { 'chat-completion': '/chat/completions' } },
  { name: 'zhipu', apiBaseURL: 'https://open.bigmodel.cn/api/paas/v4/', endpoints: { 'chat-completion': '/chat/completions' } },
] as const satisfies Provider<string[]>[]
