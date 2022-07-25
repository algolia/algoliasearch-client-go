/* eslint-disable import/no-commonjs */
/* eslint-disable @typescript-eslint/no-var-requires */
const execa = require('execa');

async function run(command) {
  return (
    (
      await execa.command(command, {
        shell: 'bash',
        all: true,
      })
    ).all ?? ''
  );
}

module.exports = { run };
