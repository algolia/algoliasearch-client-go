#!/usr/bin/env node
/* eslint-disable @typescript-eslint/no-var-requires */
/* eslint-disable import/no-commonjs */
const { exit } = require('process');

const ora = require('ora-classic');

const { run } = require('./utils');

async function formatGenerators() {
  const diff = (await run('git diff --name-only --cached -- generators')).split(
    '\n'
  );
  if (diff.length === 0 || diff[0].trim() === '') {
    return;
  }
  const spinner = ora('Linting generators').start();
  try {
    await run('yarn docker:no-tty format java generators');
    await run(`git add ${diff.join(' ')}`);
  } catch (e) {
    // eslint-disable-next-line no-console
    console.log(e);
    spinner.fail('Failed to format generators');
    exit(1);
  }
  spinner.succeed();
}

if (require.main === module && process.env.CI !== 'true') {
  formatGenerators();
}
