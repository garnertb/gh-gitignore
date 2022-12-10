#!/usr/bin/env node

const { Octokit } = require("@octokit/rest");
const fetch = require('node-fetch');
const fs = require('fs').promises;

const octokit = new Octokit({
  auth: process.env.GITHUB_TOKEN,
});

const parseOwnershipFileResponse = (response) => {
  const {
    data
  } = response
  const content = Buffer.from(data.content, 'base64').toString()
  return content
}


const getFiles = async () => {
  const payload = await fetch('https://api.github.com/repos/github/gitignore/git/trees/main?recursive=1')
  const data = await payload.json()
  await fs.writeFile('gitignore.json', JSON.stringify(data, null, 2))
  return data
}


(async ( ) => {
  await getFiles()
  console.log('Updated gitignore.json')
})()