import { OWNER, REPO } from '../../release/common';

const REPO_URL = `https://github.com/${OWNER}/${REPO}`;

export default {
  notification: {
    header: '🔨 The codegen job will run at the end of the CI.',
    body: '_Make sure your last commit does not contains generated code, it will be automatically pushed by our CI._',
  },
  noGen: {
    header: '✗ No code generated.',
    body: `_If you believe this is an issue on our side, please [open an issue](${REPO_URL}/issues/new?template=Bug_report.md)._`,
  },
  cleanup: {
    header: '✗ The generated branch has been deleted.',
    body: `If the PR has been merged, you can check the generated code on the [\`generated/main\` branch](${REPO_URL}/tree/generated/main).`,
  },
  codegen: {
    header: '✔️ Code generated!',
    body(
      branch: string,
      commit: string,
      eventNumber: number,
      generatedCommit: string
    ): string {
      return `🔨 Triggered by commit [${commit}](${REPO_URL}/pull/${eventNumber}/commits/${commit}).
🔍 Browse the generated code on branch [${branch}](${REPO_URL}/tree/${branch}): [${generatedCommit}](${REPO_URL}/commit/${generatedCommit}).`;
    },
  },
};
