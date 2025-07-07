import type { Model } from './issue-parser.ts'
import remarkParse from 'remark-parse'
import { unified } from 'unified'
import { describe, expect, it } from 'vitest'
import { extractCheckedListItems, parseModelIssue } from './issue-parser.ts'

describe('extractCheckedListItems', () => {
  it('should extract checked list items', () => {
    const list = `
- [x] item 1
- [ ] item 2
- [x] item 3
`

    const tree = unified().use(remarkParse).parse(list)
    const listNode = tree.children[0]
    if (listNode.type !== 'list') {
      throw new Error('List node not found')
    }

    const items = extractCheckedListItems(listNode)
    expect(items).toEqual(['item 1', 'item 3'])
  })
})

describe('parseModelIssue', () => {
  it('should parse the issue', () => {
    const issue = `
### Provider

openai

### Model ID

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
- [ ] audio-music`

    const parsedIssue = parseModelIssue(issue)
    expect(parsedIssue).toEqual({
      capabilities: ['streaming', 'reasoning', 'tool-call'],
      endpoints: ['chat-completion', 'completion'],
      inputModalities: ['text', 'image'],
      modelId: 'o4-mini',
      outputModalities: ['text'],
      provider: 'openai',
    } satisfies Model)
  })
})
