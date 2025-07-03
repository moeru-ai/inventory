import { Octokit } from 'octokit'

const client = new Octokit({ auth: process.env.GITHUB_TOKEN })

client.rest.issues.get()
