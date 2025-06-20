// Auto-generated file. Do not edit.

import type { Provider, Model } from "../types/index";

export const minimax = {
  name: "minimax",
  apiBaseURL: "https://api.minimaxi.com/v1",
  endpoints: {"chat-completion": "/text/chatcompletion_v2",},
  models: [{
      modelId: "MiniMax-Text-01",
      provider: "minimax",
      endpoints: ["chat-completion",],
      capabilities: ["tool-call","streaming",],
      inputModalities: ["text",],
      outputModalities: ["text",],
    },] as const satisfies Model[],
} as const satisfies Provider
