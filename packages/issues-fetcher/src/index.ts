import type { CommonRequestOptions } from 'xsai'
import type { Model } from './issue-parser.ts'
import * as fs from 'node:fs'
import { createRequire } from 'node:module'
import path from 'node:path'

import { cwd, env, exit } from 'node:process'
import { models } from '@moeru-ai/jem'
import { readableStreamToAsyncIterator } from '@moeru/std'
import * as providers from '@xsai-ext/providers-cloud'
import { execa } from 'execa'
import git from 'isomorphic-git'
import { Octokit } from 'octokit'
import { generateText, streamText, tool } from 'xsai'
import { z } from 'zod'
import { parseModelIssue } from './issue-parser.ts'

const http = createRequire(import.meta.url)('isomorphic-git/http/node')
const gitUrl = new URL('https://github.com/moeru-ai/inventory.git')
gitUrl.password = env.GITHUB_TOKEN!
gitUrl.username = env.GITHUB_USERNAME!

const rootDir = path.join(cwd(), '..', '..')
const modelsFilePath = path.join(rootDir, 'packages', 'jem', 'src', 'models.ts')

function generateModelsFileContent(models: Model[]) {
  return `// Auto-generated file. Do not edit.

  import type { Model } from './types.ts'

  export const models = ${JSON.stringify(models, null, 2)} as const satisfies Model[]
  `
}

function getProviderApiKey(provider: string): string {
  // Handle special cases for provider name to env key mapping
  let envKey: string
  const normalizedProvider = provider.toLowerCase()
  if (normalizedProvider === 'google' || normalizedProvider === 'google-generative-ai') {
    envKey = 'GOOGLE_API_KEY'
  }
  else {
    envKey = `${normalizedProvider.toUpperCase().replace(/-/g, '_')}_API_KEY`
  }

  const apiKey = env[envKey]
  if (!apiKey) {
    throw new Error(`${envKey} is not set`)
  }
  return apiKey
}

function createProviderInstance(provider: string) {
  const apiKey = getProviderApiKey(provider)

  // Map provider names to creation functions
  switch (provider.toLowerCase()) {
    case 'openai':
      return providers.createOpenAI(apiKey)
    case 'minimax':
      return providers.createMinimax(apiKey)
    case 'minimaxi':
      return providers.createMinimaxi(apiKey)
    case 'google':
    case 'google-generative-ai':
      return providers.createGoogleGenerativeAI(apiKey)
    case 'anthropic':
      return providers.createAnthropic(apiKey)
    case 'groq':
      return providers.createGroq(apiKey)
    case 'mistral':
      return providers.createMistral(apiKey)
    case 'deepseek':
      return providers.createDeepSeek(apiKey)
    case 'qwen':
      return providers.createQwen(apiKey)
    case 'moonshot':
      return providers.createMoonshot(apiKey)
    case 'zhipu':
      return providers.createZhipu(apiKey)
    case 'stepfun':
      return providers.createStepfun(apiKey)
    case 'silicon-flow':
      return providers.createSiliconFlow(apiKey)
    case 'together-ai':
      return providers.createTogetherAI(apiKey)
    case 'fireworks':
      return providers.createFireworks(apiKey)
    case 'perplexity':
      return providers.createPerplexity(apiKey)
    case 'openrouter':
      return providers.createOpenRouter(apiKey)
    case 'xai':
      return providers.createXAI(apiKey)
    case 'novita':
      return providers.createNovita(apiKey)
    case 'deepinfra':
      return providers.createDeepInfra(apiKey)
    case 'cerebras':
      return providers.createCerebras(apiKey)
    // TODO: Add support for fatherless-ai
    case 'fatherless-ai':
    case 'featherless-ai':
    default:
      throw new Error(`Unsupported provider: ${provider}`)
  }
}

async function detectTextGeneration(chat: CommonRequestOptions): Promise<boolean> {
  try {
    await generateText({
      ...chat,
      messages: [
        {
          content: 'Say "hello"',
          role: 'user',
        },
      ],
    })
    return true
  }
  catch {
    return false
  }
}

async function detectStreamGeneration(chat: CommonRequestOptions): Promise<boolean> {
  try {
    const stream = await streamText({
      ...chat,
      messages: [
        {
          content: 'Say "hello"',
          role: 'user',
        },
      ],
    })
    // Try to read at least one chunk
    const iterator = readableStreamToAsyncIterator(stream.textStream, async it => it.toString())
    const firstChunk = await iterator.next()
    return !firstChunk.done
  }
  catch {
    return false
  }
}

async function detectToolCall(chat: CommonRequestOptions): Promise<boolean> {
  try {
    const response = await generateText({
      ...chat,
      messages: [
        {
          content: 'What is 1 + 1?',
          role: 'user',
        },
      ],
      tools: [
        await tool({
          name: 'add',
          description: 'Add two numbers',
          parameters: z.object({
            a: z.number(),
            b: z.number(),
          }),
          execute: async (args) => {
            return args.a + args.b
          },
        }),
      ],
    })
    // Check if tool was called
    return response.toolCalls && response.toolCalls.length > 0
  }
  catch {
    return false
  }
}

async function detectModelCapabilities(provider: string, modelId: string): Promise<Partial<Model>> {
  console.log(`Detecting capabilities for ${provider}/${modelId}...`)

  const providerInstance = createProviderInstance(provider)
  const chat = providerInstance.chat(modelId)

  const detectedModel: Partial<Model> = {
    modelId,
    provider,
    capabilities: [],
    endpoints: ['chat-completion'],
    inputModalities: ['text'],
    outputModalities: [],
  }

  // Detect text output (most models support this)
  const supportsText = await detectTextGeneration(chat)
  if (supportsText) {
    detectedModel.outputModalities = ['text']
    console.log('  ✓ Text generation supported')
  }

  // Detect streaming
  if (supportsText) {
    const supportsStreaming = await detectStreamGeneration(chat)
    if (supportsStreaming) {
      detectedModel.capabilities = detectedModel.capabilities || []
      detectedModel.capabilities.push('streaming')
      console.log('  ✓ Streaming supported')
    }
  }

  // Detect tool call
  if (supportsText) {
    const supportsToolCall = await detectToolCall(chat)
    if (supportsToolCall) {
      detectedModel.capabilities = detectedModel.capabilities || []
      detectedModel.capabilities.push('tool-call')
      console.log('  ✓ Tool call supported')
    }
  }

  // Default to text if no output modality was detected
  if (!detectedModel.outputModalities || detectedModel.outputModalities.length === 0) {
    detectedModel.outputModalities = ['text']
  }

  console.log(`Detection complete: ${JSON.stringify(detectedModel, null, 2)}`)
  return detectedModel
}

async function main() {
  const issueId = env.ISSUE_ID
  const isModified = env.TRIGGER_ACTION === 'edited'

  if (!issueId) {
    throw new Error('ISSUE_ID is not set')
  }

  const client = new Octokit({ auth: env.GITHUB_TOKEN })

  const issue = await client.rest.issues.get({
    issue_number: Number(issueId),
    owner: 'moeru-ai',
    repo: 'inventory',
  })
  if (!issue.data.title.includes('Model Collection')) {
    throw new Error('Not a model collection issue')
  }
  if (!issue.data.body) {
    throw new Error('Issue body is empty')
  }

  const parsedModelInfo = await parseModelIssue(issue.data.body)
  if (!parsedModelInfo.modelId) {
    throw new Error('Model ID not found in the issue body')
  }
  if (!parsedModelInfo.provider) {
    throw new Error('Provider not found in the issue body')
  }

  // Auto-detect model capabilities
  const detectedCapabilities = await detectModelCapabilities(parsedModelInfo.provider, parsedModelInfo.modelId)

  // Merge parsed info with detected capabilities
  if (!detectedCapabilities.modelId || !detectedCapabilities.provider) {
    throw new Error('Failed to detect model capabilities')
  }

  const modelInfo: Model = {
    modelId: detectedCapabilities.modelId,
    provider: detectedCapabilities.provider,
    capabilities: detectedCapabilities.capabilities || [],
    endpoints: detectedCapabilities.endpoints || ['chat-completion'],
    inputModalities: detectedCapabilities.inputModalities || ['text'],
    outputModalities: detectedCapabilities.outputModalities || ['text'],
  }

  const existingModel = models.find(it => it.modelId === modelInfo.modelId && it.provider === modelInfo.provider)
  if (existingModel && JSON.stringify(existingModel) === JSON.stringify(modelInfo)) {
    throw new Error('Existing model is the same as the new model')
  }

  const comments = await client.rest.issues.listComments({
    issue_number: Number(issueId),
    owner: 'moeru-ai',
    repo: 'inventory',
  })

  const prIssueCommentPrefix = 'Thank you for your contribution! Pull request created: #'
  const pullRequestComment = comments.data.find(it => it.body?.startsWith(prIssueCommentPrefix))
  const branchName = `model-collection/issue-${issueId}`
  let prNumber: number | undefined
  let pr: Awaited<ReturnType<typeof client.rest.pulls.get>> | undefined

  if (pullRequestComment?.body) {
    prNumber = Number(pullRequestComment.body.slice(prIssueCommentPrefix.length).trim())
    const existingPr = await client.rest.pulls.get({
      owner: 'moeru-ai',
      repo: 'inventory',
      pull_number: prNumber,
    })
    console.log(`Pull request found: #${prNumber}`)

    if (!existingPr.data.closed_at) {
      pr = existingPr
    }
    else {
      console.log(`Pull request closed, need to create a new one: #${prNumber}`)
    }
  }

  if (!pr) {
    await git.branch({
      fs,
      dir: rootDir,
      ref: branchName,
      checkout: true,
      force: true,
    })
    console.log(`Created a new branch: ${branchName}`)
  }
  else {
    await git.fetch({ fs, dir: rootDir, http })
    await git.checkout({ fs, dir: rootDir, ref: branchName })
    console.log(`Checked out the branch: ${branchName}`)
  }

  const existingModels = models.filter(
    it => it.provider !== modelInfo.provider || it.modelId !== modelInfo.modelId,
  ) as Model[]
  existingModels.push(modelInfo)

  const newModelsFileContent = generateModelsFileContent(existingModels)
  await fs.promises.writeFile(modelsFilePath, newModelsFileContent)
  console.log(`Wrote to ${modelsFilePath}`)

  console.log('Running ESLint...')
  const { stdout: eslintOutput, stderr: eslintError, exitCode } = await execa('pnpm', ['-F', '@moeru-ai/jem', 'run', 'lint:fix'])
  console.log(eslintOutput)

  if (exitCode !== 0) {
    throw new Error(`ESLint failed: ${eslintError}`)
  }

  await git.add({ fs, dir: rootDir, filepath: path.relative(rootDir, modelsFilePath) })
  let commitMessage = isModified ? `chore: update ${modelInfo.modelId} in the inventory` : `feat: add ${modelInfo.modelId} to the inventory`
  commitMessage += `

  Co-authored-by: ${issue.data.user?.login} <${issue.data.user?.id}+${issue.data.user?.login}@users.noreply.github.com>`
  await git.commit({ fs, dir: rootDir, message: commitMessage, author: { name: 'github-actions[bot]', email: 'github-actions@github.com' } })
  console.log('Committed')
  await git.push({ fs, http, dir: rootDir, ref: branchName, remote: 'origin', url: gitUrl.toString(), force: true, remoteRef: `refs/heads/${branchName}` })
  console.log(`Pushed to origin/${branchName}`)

  const prTitle = isModified ? `chore: update ${modelInfo.modelId} in the inventory` : `feat: add ${modelInfo.modelId} to the inventory`
  const prBody = `This PR is auto generated, associated with the issue #${issueId}.`

  if (!pr) {
    const pr = await client.rest.pulls.create({
      owner: 'moeru-ai',
      repo: 'inventory',
      head: branchName,
      base: 'main',
      title: prTitle,
      body: prBody,
    })
    prNumber = pr.data.number
    console.log(`Pull request created: #${pr.data.number}`) // TODO: replace with @guiiai/logg

    await client.rest.issues.createComment({
      issue_number: Number(issueId),
      owner: 'moeru-ai',
      repo: 'inventory',
      body: `${prIssueCommentPrefix}${prNumber}`,
    })

    console.log(`Commented on the issue: #${issueId} with the pull request link.`)
  }
}

main().catch((err) => {
  console.error(err)
  exit(1)
})
