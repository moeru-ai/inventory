import type { List } from 'mdast'
import { toString } from 'mdast-util-to-string'
import remarkParse from 'remark-parse'
import { unified } from 'unified'
import { visit } from 'unist-util-visit'

// Move this to a shared location.
export interface Model {
  capabilities: string[]
  endpoints: string[]
  inputModalities: string[]
  modelId: string
  outputModalities: string[]
  provider: string
}

// Helper to extract checked items from a list
export function extractCheckedListItems(listNode: List) {
  return listNode.children
    .map(item => toString(item))
    .filter(text => text.startsWith('[x]'))
    .map(text => text.replace(/^\[x\]\s*/i, '').trim())
}

export function parseModelIssue(markdown: string): Partial<Model> {
  // Parse the markdown into an mdast (Markdown AST)
  const tree = unified().use(remarkParse).parse(markdown)

  // This object will hold the extracted model data
  const model: Partial<Model> = {
    modelId: '',
    provider: '',
  }

  // State variables to track which section we're in
  let currentSection = ''

  // Traverse the AST to extract relevant fields
  visit(tree, (node) => {
    // Section headings (### ...)
    if (node.type === 'heading' && node.depth === 3) {
      currentSection = toString(node).trim()
    }

    // Code blocks for Provider and Model ID
    if (node.type === 'paragraph' && (currentSection === 'Provider' || currentSection === 'Model ID')) {
      if (currentSection === 'Provider') {
        model.provider = toString(node).trim()
      }
      if (currentSection === 'Model ID') {
        model.modelId = toString(node).trim()
      }
    }
  })

  return model
}
