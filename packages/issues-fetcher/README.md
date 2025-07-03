# Issues Fetcher

This tool fetches issues from `https://github.com/moeru-ai/inventory/issues` and parse form structure to generate a pull request for the issue.

## Model Issue Form Structure

An issue form structure is like this:

```markdown
This issue is about to add model `o4-mini` from `openai` to the inventory.

### Provider

openai

### Model Name

o4-mini

### Model Capabilities

- [x] streaming
- [x] reasoning
- [x] tool-call

### Model Input Modalities

- [x] text
- [x] image
- [ ] audio
- [ ] video
- [ ] vector

### Model Output Modalities

- [x] text
- [ ] image
- [ ] audio
- [ ] video
- [ ] vector

### Model Endpoints

- [x] chat-completion
- [x] completion
- [ ] embedding
- [ ] image-generation
- [ ] audio-speech
- [ ] audio-music

### Additional Information

Additional information about the model.
```

It will be parsed to a TypeScript object like this:

```typescript
import type { Model } from '@moeru-ai/inventory/types'

const model = {
  Provider: 'openai',
  ModelName: 'o4-mini',
  Capabilities: ['streaming', 'reasoning', 'tool-call'],
  InputModalities: ['text', 'image'],
  OutputModalities: ['text'],
  Endpoints: ['chat-completion', 'completion'],
} as const satisfies Model
```

## Q&A

### Why we don't need a provider issue form structure?

If a model is from a unknown provider, CI pipeline will fail so we will add the provider manually.
