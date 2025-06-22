# Issues Fetcher

This tool fetches issues from `https://github.com/moeru-ai/inventory/issues` and parse form structure to generate a pull request for the issue.

## Model Issue Form Structure

An issue form structure is like this:

```markdown
This issue is about to add/modify model `o4-mini` from `openai` to the inventory.

## Provider

```text
openai
```

## Model Name

```text
o4-mini
```

## Model Capabilities

- [x] streaming
- [x] reasoning
- [x] tool-call

## Model Input Modalities

- [x] text
- [x] image
- [ ] audio
- [ ] video
- [ ] vector

## Model Output Modalities

- [x] text
- [ ] image
- [ ] audio
- [ ] video
- [ ] vector

## Model Endpoints

- [x] chat-completion
- [x] completion
- [ ] embedding
- [ ] image-generation
- [ ] audio-speech
- [ ] audio-music

```

It will be parsed to a Golang struct like this:

```go
import "github.com/moeru-ai/inventory/types"

var model = types.Model{
  Provider: "openai",
  ModelName: "o4-mini",
  Capabilities: []types.Capability{
    types.CapabilityStreaming,
    types.CapabilityReasoning,
    types.CapabilityToolCall,
  },
  InputModalities: []types.Modality{
    types.ModalityText,
    types.ModalityImage,
  },
  OutputModalities: []types.Modality{
    types.ModalityText,
  },
  Endpoints: []types.EndpointType{
    types.EndpointTypeChatCompletion,
    types.EndpointTypeCompletion,
  },
}
```

## Q&A

### Why we don't need a provider issue form structure?

If a model is from a unknown provider, CI pipeline will fail so we will add the provider manually.
