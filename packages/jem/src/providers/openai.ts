// Auto-generated file. Do not edit.

import type { Provider, Model } from "../types/index";

export const openai = {
  name: "openai",
  apiBaseURL: "https://api.openai.com/v1",
  endpoints: {"chat-completion": "/chat/completions",},
  models: [{
      modelId: "gpt-4o",
      provider: "openai",
      endpoints: ["chat-completion",],
      capabilities: ["tool-call","streaming",],
      inputModalities: ["text",],
      outputModalities: ["text",],
    },] as const satisfies Model[],
} as const satisfies Provider
