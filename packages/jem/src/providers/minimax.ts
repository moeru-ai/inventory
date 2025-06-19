// Auto-generated file. Do not edit.

import type { Provider } from "../types/index";

export const minimax: Provider = {
  name: "minimax",
  apiBaseURL: "https://api.minimaxi.com/v1",
  endpoints: {
    "chat-completion": "/text/chatcompletion_v2",
    
  },
  models: [
    {
      modelId: "MiniMax-Text-01",
      provider: "{[tool-call streaming] [chat-completion] [text] MiniMax-Text-01 [text] minimax}",
      endpoints: [
        "chat-completion",
        
      ],
      capabilities: [
        "tool-call",
        "streaming",
        
      ],
      inputModalities: [
        "text",
        
      ],
      outputModalities: [
        "text",
        
      ],
    },
    
  ],
}
