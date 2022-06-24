#!/usr/bin/env node
/* eslint-disable no-console */
/* eslint-disable import/no-commonjs */
/* eslint-disable @typescript-eslint/no-var-requires */
const chalk = require('chalk');
const execa = require('execa');
const micromatch = require('micromatch');

const clientConfig = require('../../../config/clients.config.json');
const GENERATED_FILE_PATTERNS =
  require('../../../config/generation.config').patterns;

const run = async (command, { cwd } = {}) => {
  return (
    (await execa.command(command, { shell: 'bash', all: true, cwd })).all ?? ''
  );
};

function getPatterns() {
  const patterns = GENERATED_FILE_PATTERNS;
  for (const [language, { tests }] of Object.entries(clientConfig)) {
    patterns.push(`tests/output/${language}/${tests.outputFolder}/client/**`);
    patterns.push(`tests/output/${language}/${tests.outputFolder}/methods/**`);
  }
  return patterns;
}

async function preCommit() {
  // when merging, we want to stage all the files
  try {
    await run('git merge HEAD');
  } catch (e) {
    if (e.exitCode === 128) {
      console.log(
        'Skipping the pre-commit check because a merge is in progress'
      );
      return;
    }
  }

  const stagedFiles = (await run('git diff --name-only --cached')).split('\n');

  const toUnstage = micromatch.match(stagedFiles, getPatterns());
  if (toUnstage.length === 0) {
    return;
  }

  toUnstage.forEach((file) =>
    console.log(
      chalk.black.bgYellow('[INFO]'),
      `Generated file found, unstaging: ${file}`
    )
  );
  await run(`git restore --staged ${toUnstage.join(' ')}`);
}

if (require.main === module && process.env.CI !== 'true') {
  preCommit();
}

module.exports = { getPatterns };
