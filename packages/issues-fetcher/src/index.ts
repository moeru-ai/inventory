import type { Model, ModelIdsByProvider, ProviderNames } from '@proj-airi/jem'
import * as fs from 'node:fs'
import path from 'node:path'
import * as process from 'node:process'
import { models } from '@proj-airi/jem'

import git from 'isomorphic-git'
import { Octokit } from 'octokit'
import { parseModelIssue } from './issue-parser.ts'
import { createRequire } from 'node:module'

const http = createRequire(import.meta.url)('isomorphic-git/http/node')
const gitUrl = new URL('https://github.com/moeru-ai/inventory.git')
gitUrl.password = process.env.GITHUB_TOKEN!
gitUrl.username = process.env.GITHUB_USERNAME!

const cwd = process.cwd()
const modelsFilePath = path.join(cwd, '..', 'jem', 'src', 'models.ts')

function generateModelsFileContent(models: Model<ProviderNames, ModelIdsByProvider<ProviderNames>>[]) {
  return `// Auto-generated file. Do not edit.

  import type { Model } from './types'

  export const models = ${JSON.stringify(models, null, 2)} as const satisfies Model[]
  `
}

async function main() {
  const issueId = process.env.ISSUE_ID
  const isModified = process.env.TRIGGER_ACTION === 'edited'

  if (!issueId) {
    throw new Error('ISSUE_ID is not set')
  }

  const client = new Octokit({ auth: process.env.GITHUB_TOKEN })

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

  const modelInfo = await parseModelIssue(issue.data.body)
  if (!modelInfo.modelId) {
    throw new Error('Model ID not found in the issue body')
  }
  if (!modelInfo.provider) {
    throw new Error('Provider not found in the issue body')
  }
  const existingModel = models.find(it => it.modelId === modelInfo.modelId)
  if (existingModel === modelInfo) {
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

  if (pullRequestComment) {
    prNumber = Number(pullRequestComment.body!.slice(prIssueCommentPrefix.length).trim())
    pr = await client.rest.pulls.get({
      owner: 'moeru-ai',
      repo: 'inventory',
      pull_number: prNumber,
    })
  }

  if (!pr) {
    await git.branch({
      fs,
      dir: cwd,
      ref: branchName,
      checkout: true,
      force: true
    })
  }
  else {
    await git.checkout({ fs, dir: cwd, ref: branchName })
  }

  const existingModels = models.filter(
    it => it.provider !== modelInfo.provider && it.modelId !== modelInfo.modelId,
  )
  existingModels.push(modelInfo as Model<ProviderNames, ModelIdsByProvider<ProviderNames>>)

  const newModelsFileContent = generateModelsFileContent(existingModels)
  await fs.promises.writeFile(modelsFilePath, newModelsFileContent)
  console.log(`Wrote to ${modelsFilePath}`)

  const rootDir = path.join(cwd, '..', '..')
  await git.add({ fs, dir: rootDir, filepath: path.relative(rootDir, modelsFilePath) })
  let commitMessage = isModified ? `chore: update ${modelInfo.modelId} in the inventory` : `feat: add ${modelInfo.modelId} to the inventory`
  commitMessage += `

  Co-authored-by: ${issue.data.user?.login} <${issue.data.user?.id}+${issue.data.user?.login}@users.noreply.github.com>`
  await git.commit({ fs, dir: rootDir, message: commitMessage, author: { name: 'github-actions[bot]', email: 'github-actions@github.com' } })
  console.log('Committed')
  await git.push({ fs, http, dir: rootDir, ref: branchName, remote: 'origin', url: gitUrl.toString() })
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

main().catch(console.error)
