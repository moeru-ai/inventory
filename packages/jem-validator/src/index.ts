import type { CommonRequestOptions } from 'xsai'
import { env, exit } from 'node:process'
import { models } from '@moeru-ai/jem'
import { readableStreamToAsyncIterator } from '@moeru/std'
import * as providers from '@xsai-ext/providers-cloud'
import { generateText, streamText, tool } from 'xsai'
import { z } from 'zod'

async function checkTextGeneration(chat: CommonRequestOptions, modelId: string) {
  console.log(`${modelId} - Text Generation:`)

  const generateTextResponse = await generateText({
    ...chat,
    messages: [
      {
        content: 'You are a cute cat girl.',
        role: 'system',
      },
      {
        content: `What is ${Math.floor(Math.random() * 100)} + ${Math.floor(Math.random() * 100)}?`,
        role: 'user',
      },
    ],
  })

  console.log(generateTextResponse.text)
  console.log('='.repeat(50))
}

async function checkStreamGeneration(chat: CommonRequestOptions, modelId: string) {
  console.log(`${modelId} - Stream Generation:`)

  const stream = await streamText({
    ...chat,
    messages: [
      {
        content: 'You are a cute cat girl.',
        role: 'system',
      },
    ],
  })

  const chunks: string[] = []
  for await (const chunk of readableStreamToAsyncIterator(stream.textStream, async it => it.toString())) {
    chunks.push(chunk)
  }


  console.log(chunks)
  console.log('='.repeat(50))
}

async function checkToolCallGeneration(chat: CommonRequestOptions, modelId: string) {
  console.log(`${modelId} - Tool Call Generation:`)

  const toolCallResponse = await generateText({
    ...chat,
    messages: [
      {
        content: 'You are a cute cat girl.',
        role: 'system',
      },
      {
        content: `What is ${Math.floor(Math.random() * 100)} + ${Math.floor(Math.random() * 100)}?`,
        role: 'user',
      },
    ],
    tools: [
      await tool({
        name: 'plus',
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

  console.log(toolCallResponse.text)
  console.log('='.repeat(50))
}

async function main() {
  if (!env.OPENAI_API_KEY) {
    throw new Error('OPENAI_API_KEY is not set')
  }
  if (!env.MINIMAX_API_KEY) {
    throw new Error('MINIMAX_API_KEY is not set')
  }
  if (!env.MINIMAXI_API_KEY) {
    throw new Error('MINIMAX_API_KEY is not set')
  }
  if (!env.GOOGLE_API_KEY) {
    throw new Error('GOOGLE_API_KEY is not set')
  }


  const openai = providers.createOpenAI(env.OPENAI_API_KEY)
  const minimax = providers.createMinimax(env.MINIMAX_API_KEY)
  const minimaxi = providers.createMinimaxi(env.MINIMAXI_API_KEY)
  const google = providers.createGoogleGenerativeAI(env.GOOGLE_API_KEY)

  for (const model of models) {
    let chat: CommonRequestOptions | undefined
    switch (model.provider) {
      case 'openai':
        chat = openai.chat(model.modelId)
        break
      case 'minimax':
        chat = minimax.chat(model.modelId)
        break
      case 'minimaxi':
        chat = minimaxi.chat(model.modelId)
        break
      case 'google':
        chat = google.chat(model.modelId)
        break
    }

    if (!chat) {
      throw new Error(`Unsupported provider: ${model.provider}`)
    }

    for (const modality of model.outputModalities) {
      switch (modality) {
        case 'audio':
          // await checkAudioGeneration(chat, model.modelId) TODO: Audio check
          break
        case 'text':
          await checkTextGeneration(chat, model.modelId)
          break
        default:
          throw new Error(`Unsupported modality: ${modality}`)
      }
    }

    for (const capability of model.capabilities) {
      switch (capability) {
        case 'streaming':
          await checkStreamGeneration(chat, model.modelId)
          break
        case 'tool-call':
          await checkToolCallGeneration(chat, model.modelId)
          break
        default:
          throw new Error(`Unsupported capability: ${capability}`)
      }
    }
  }
}

main().catch((error) => {
  console.error(error)
  exit(1)
})
