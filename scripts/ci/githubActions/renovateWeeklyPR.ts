/* eslint-disable no-console */
// script coming from crawler <3
import type { Octokit } from '@octokit/rest';

import { getOctokit, OWNER, REPO, wait } from '../../common';

const BRANCH = 'chore/renovateBaseBranch';
const BRANCH_BASE = 'main';
const EMPTY_COMMIT_MSG = 'Automatic empty commit';

async function getRef(
  octokit: Octokit,
  branch: string
): Promise<string | false> {
  try {
    const ref = await octokit.git.getRef({
      owner: OWNER,
      repo: REPO,
      ref: `heads/${branch}`,
    });
    return ref.data.object.sha;
  } catch (err) {
    if (!(err instanceof Error) || (err as any).status !== 404) {
      throw err;
    }
  }
  return false;
}

async function createBranch(octokit: Octokit, sha: string): Promise<any> {
  const create = await octokit.git.createRef({
    owner: OWNER,
    repo: REPO,
    ref: `refs/heads/${BRANCH}`,
    sha,
  });
  return create;
}

async function deleteRef(octokit: Octokit): Promise<any> {
  console.log(`Deleting ref for ${BRANCH}`);
  const ref = await octokit.git.deleteRef({
    owner: OWNER,
    repo: REPO,
    ref: `heads/${BRANCH}`,
  });
  return ref;
}

async function updateRef(octokit: Octokit, sha: string): Promise<any> {
  console.log(`Changing ref for ${BRANCH} to`, sha);
  const ref = await octokit.git.updateRef({
    owner: OWNER,
    repo: REPO,
    ref: `heads/${BRANCH}`,
    sha,
  });
  return ref;
}

async function getCommit(octokit: Octokit, sha: string): Promise<any> {
  const commit = await octokit.git.getCommit({
    owner: OWNER,
    repo: REPO,
    commit_sha: sha,
  });
  return commit.data;
}

function isCommitAnEmptyCommit(commit: any): boolean {
  return commit.message.search(EMPTY_COMMIT_MSG) >= 0;
}

async function createEmptyCommit(
  octokit: Octokit,
  refCommit: any
): Promise<any> {
  console.log('Creating empty commit');
  const commit = await octokit.git.createCommit({
    owner: OWNER,
    repo: REPO,
    message: EMPTY_COMMIT_MSG,
    tree: refCommit.tree.sha,
    parents: [refCommit.sha],
  });
  return commit.data;
}

async function createPR(octokit: Octokit): Promise<any> {
  // Next monday
  const date = new Date();
  date.setDate(date.getDate() + 3);

  const title = `chore(scripts): dependencies ${
    date.toISOString().split('T')[0]
  }`;
  const { data } = await octokit.pulls.create({
    repo: REPO,
    owner: OWNER,
    title,
    body: `Weekly dependencies update.
Contributes to #532
    `,
    head: BRANCH,
    base: BRANCH_BASE,
  });
  return data;
}

async function resetBranch(
  octokit: Octokit,
  refBase: string,
  exists: boolean
): Promise<void> {
  if (exists) {
    console.log('Deleting branch');
    await deleteRef(octokit);
    await wait(5000);
  }

  console.log('Creating branch');

  await createBranch(octokit, refBase);

  const commit = await getCommit(octokit, refBase);

  const empty = await createEmptyCommit(octokit, commit);
  await updateRef(octokit, empty.sha);
}

(async (): Promise<void> => {
  try {
    const octokit = getOctokit();

    const refBase = await getRef(octokit, BRANCH_BASE);
    const refTarget = await getRef(octokit, BRANCH);
    console.log(BRANCH_BASE, 'is at', refBase);
    console.log(BRANCH, 'is at', refTarget);

    if (!refBase) {
      console.error('no sha for base branch');
      return;
    }

    if (refTarget) {
      console.log('Branch exists');
      const commit = await getCommit(octokit, refTarget);

      if (isCommitAnEmptyCommit(commit)) {
        console.log('Empty commit exists');
        return;
      }
    }

    await resetBranch(octokit, refBase, Boolean(refTarget));

    console.log('Creating pull request');
    await createPR(octokit);
  } catch (err) {
    console.error(err);
  }
})();
