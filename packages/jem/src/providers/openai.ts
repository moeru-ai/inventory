// Auto-generated file. Do not edit.

import type { Provider } from "../types/index";

export const openai: Provider = {
  name: "openai",
  apiBaseURL: "https://api.openai.com/v1",
  endpoints: {
    "chat-completion": "/chat/completions",
    
  },
  models: [
    {
      modelId: "gpt-4o",
      provider: "{[tool-call streaming] [chat-completion] [text] gpt-4o [text] openai}",
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
