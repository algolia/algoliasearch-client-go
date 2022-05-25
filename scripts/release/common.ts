import path from 'path';

import config from '../../config/release.config.json';
import { run } from '../common';
import { getGitHubUrl } from '../config';
import type { Language } from '../types';

export const RELEASED_TAG = config.releasedTag;
export const TEAM_SLUG = config.teamSlug;

export function getTargetBranch(language: string): string {
  return config.targetBranch[language] || config.defaultTargetBranch;
}

export function getGitAuthor(): { name: string; email: string } {
  return config.gitAuthor;
}

export async function configureGitHubAuthor(cwd?: string): Promise<void> {
  const { name, email } = getGitAuthor();

  await run(`git config user.name "${name}"`, { cwd });
  await run(`git config user.email "${email}"`, { cwd });
}

export async function cloneRepository({
  lang,
  githubToken,
  tempDir,
}: {
  lang: Language;
  githubToken: string;
  tempDir: string;
}): Promise<{ tempGitDir: string }> {
  const targetBranch = getTargetBranch(lang);

  const gitHubUrl = getGitHubUrl(lang, { token: githubToken });
  const tempGitDir = path.resolve(tempDir, lang);
  await run(`rm -rf ${tempGitDir}`);
  await run(
    `git clone --depth 1 --branch ${targetBranch} ${gitHubUrl} ${tempGitDir}`
  );

  return {
    tempGitDir,
  };
}
