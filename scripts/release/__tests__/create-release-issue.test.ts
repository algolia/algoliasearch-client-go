import generationCommitText from '../../ci/codegen/text';
import {
  parseCommit,
  getVersionChangesText,
  getSkippedCommitsText,
  decideReleaseStrategy,
  readVersions,
} from '../create-release-issue';

describe('create release issue', () => {
  it('reads versions of the current language', () => {
    expect(readVersions()).toEqual({
      java: {
        current: expect.any(String),
      },
      javascript: { current: expect.any(String) },
      php: { current: expect.any(String) },
    });
  });

  describe('parseCommit', () => {
    it('parses commit', () => {
      expect(parseCommit(`b2501882 fix(javascript): fix the thing`)).toEqual({
        hash: 'b2501882',
        scope: 'javascript',
        message: 'fix the thing',
        raw: 'b2501882 fix(javascript): fix the thing',
        type: 'fix',
      });
    });

    it('considers `specs` as a lang commit', () => {
      expect(parseCommit(`b2501882 fix(specs): fix the thing`)).toEqual({
        hash: 'b2501882',
        scope: 'specs',
        message: 'fix the thing',
        raw: 'b2501882 fix(specs): fix the thing',
        type: 'fix',
      });
    });

    it('returns error when language scope is missing', () => {
      expect(parseCommit(`b2501882 fix: fix the thing`)).toEqual({
        error: 'missing-language-scope',
      });
    });

    it('returns error when language scope is unknown', () => {
      expect(parseCommit(`b2501882 fix(basic): fix the thing`)).toEqual({
        error: 'unknown-language-scope',
      });
    });

    it('returns error when it is a generated commit', () => {
      expect(
        parseCommit(
          `49662518 ${generationCommitText.commitStartMessage} ABCDEF`
        )
      ).toEqual({
        error: 'generation-commit',
      });
    });

    it('returns error when it is a generated commit, even with other casing', () => {
      expect(
        parseCommit(
          `49662518 ${generationCommitText.commitStartMessage.toLocaleUpperCase()} ABCDEF`
        )
      ).toEqual({
        error: 'generation-commit',
      });
    });
  });

  describe('getVersionChangesText', () => {
    it('generates text for version changes', () => {
      expect(
        getVersionChangesText({
          javascript: {
            current: '0.0.1',
            releaseType: 'patch',
          },

          php: {
            current: '0.0.1',
            releaseType: 'patch',
          },

          java: {
            current: '0.0.1',
            releaseType: 'patch',
          },
        })
      ).toMatchInlineSnapshot(`
      "- javascript: 0.0.1 -> **\`patch\` _(e.g. 0.0.2)_**
      - java: 0.0.1 -> **\`patch\` _(e.g. 0.0.2)_**
      - php: 0.0.1 -> **\`patch\` _(e.g. 0.0.2)_**"
    `);
    });

    it('generates text for version changes with a language with no commit', () => {
      expect(
        getVersionChangesText({
          javascript: {
            current: '0.0.1',
            releaseType: 'patch',
          },

          php: {
            current: '0.0.1',
            releaseType: null,
            noCommit: true,
          },

          java: {
            current: '0.0.1',
            releaseType: 'patch',
          },
        })
      ).toMatchInlineSnapshot(`
      "- javascript: 0.0.1 -> **\`patch\` _(e.g. 0.0.2)_**
      - java: 0.0.1 -> **\`patch\` _(e.g. 0.0.2)_**
      - ~php: 0.0.1 (no commit)~"
    `);
    });

    it('generates text for version changes with a language to skip', () => {
      expect(
        getVersionChangesText({
          javascript: {
            current: '0.0.1',
            releaseType: 'patch',
          },

          php: {
            current: '0.0.1',
            releaseType: 'minor',
          },

          java: {
            current: '0.0.1',
            releaseType: 'patch',
            skipRelease: true,
          },
        })
      ).toMatchInlineSnapshot(`
      "- javascript: 0.0.1 -> **\`patch\` _(e.g. 0.0.2)_**
      - ~java: 0.0.1 -> **\`patch\` _(e.g. 0.0.2)_**~
        - No \`feat\` or \`fix\` commit, thus unchecked by default.
      - php: 0.0.1 -> **\`minor\` _(e.g. 0.1.0)_**"
    `);
    });
  });

  describe('decideReleaseStrategy', () => {
    it('bumps major version for BREAKING CHANGE', () => {
      const versions = decideReleaseStrategy({
        versions: {
          javascript: {
            current: '0.0.1',
          },
          java: {
            current: '0.0.1',
          },
          php: {
            current: '0.0.1',
          },
        },
        commits: [
          {
            hash: 'b2501882',
            type: 'feat',
            scope: 'javascript',
            message: 'update the API (BREAKING CHANGE)',
            raw: 'b2501882 feat(javascript): update the API (BREAKING CHANGE)',
          },
        ],
      });

      expect(versions.javascript.releaseType).toEqual('major');
    });

    it('bumps minor version for feat', () => {
      const versions = decideReleaseStrategy({
        versions: {
          javascript: {
            current: '0.0.1',
          },
          java: {
            current: '0.0.1',
          },
          php: {
            current: '0.0.1',
          },
        },
        commits: [
          {
            hash: 'b2501882',
            type: 'feat',
            scope: 'php',
            message: 'update the API',
            raw: 'b2501882 feat(php): update the API',
          },
        ],
      });

      expect(versions.php.releaseType).toEqual('minor');
    });

    it('bumps patch version for fix', () => {
      const versions = decideReleaseStrategy({
        versions: {
          javascript: {
            current: '0.0.1',
          },
          java: {
            current: '0.0.1',
          },
          php: {
            current: '0.0.1',
          },
        },
        commits: [
          {
            hash: 'b2501882',
            type: 'fix',
            scope: 'java',
            message: 'fix some bug',
            raw: 'b2501882 fix(java): fix some bug',
          },
        ],
      });

      expect(versions.java.releaseType).toEqual('patch');
    });

    it('marks noCommit for languages without any commit', () => {
      const versions = decideReleaseStrategy({
        versions: {
          javascript: {
            current: '0.0.1',
          },
          java: {
            current: '0.0.1',
          },
          php: {
            current: '0.0.1',
          },
        },
        commits: [
          {
            hash: 'b2501882',
            type: 'fix',
            scope: 'java',
            message: 'fix some bug',
            raw: 'b2501882 fix(java): fix some bug',
          },
        ],
      });

      expect(versions.javascript.noCommit).toEqual(true);
      expect(versions.php.noCommit).toEqual(true);
      expect(versions.java.noCommit).toBeUndefined();
    });

    it('releases every languages if a `specs` commit is present', () => {
      const versions = decideReleaseStrategy({
        versions: {
          javascript: {
            current: '0.0.1',
          },
          java: {
            current: '0.0.1',
          },
          php: {
            current: '0.0.1',
          },
        },
        commits: [
          {
            hash: 'b2501882',
            type: 'fix',
            scope: 'specs',
            message: 'fix some descriptions',
            raw: 'b2501882 fix(specs): fix some descriptions',
          },
        ],
      });

      expect(versions.javascript.noCommit).toBeUndefined();
      expect(versions.javascript.releaseType).toEqual('patch');
      expect(versions.php.noCommit).toBeUndefined();
      expect(versions.php.releaseType).toEqual('patch');
      expect(versions.java.noCommit).toBeUndefined();
      expect(versions.java.releaseType).toEqual('patch');
    });

    it('bumps for `specs` feat with only language `fix` commits', () => {
      const versions = decideReleaseStrategy({
        versions: {
          javascript: {
            current: '0.0.1',
          },
          java: {
            current: '0.0.1',
          },
          php: {
            current: '0.0.1',
          },
        },
        commits: [
          {
            hash: 'b2501882',
            type: 'fix',
            scope: 'php',
            message: 'fix some descriptions',
            raw: 'b2501882 feat(php): fix some descriptions',
          },
          {
            hash: 'b2501882',
            type: 'feat',
            scope: 'specs',
            message: 'add some descriptions',
            raw: 'b2501882 feat(specs): add some descriptions',
          },
        ],
      });

      expect(versions.javascript.noCommit).toBeUndefined();
      expect(versions.javascript.releaseType).toEqual('minor');
      expect(versions.php.noCommit).toBeUndefined();
      expect(versions.php.releaseType).toEqual('minor');
      expect(versions.java.noCommit).toBeUndefined();
      expect(versions.java.releaseType).toEqual('minor');
    });

    it('marks skipRelease for patch upgrade without fix commit', () => {
      const versions = decideReleaseStrategy({
        versions: {
          javascript: {
            current: '0.0.1',
          },
          java: {
            current: '0.0.1',
          },
          php: {
            current: '0.0.1',
          },
        },
        commits: [
          {
            hash: 'b2501882',
            type: 'chore',
            scope: 'javascript',
            message: 'update devDevpendencies',
            raw: 'b2501882 chore(javascript): update devDevpendencies',
          },
        ],
      });
      expect(versions.javascript.skipRelease).toEqual(true);
      expect(versions.java.skipRelease).toBeUndefined();
      expect(versions.php.skipRelease).toBeUndefined();
    });
  });

  it('generates text for skipped commits', () => {
    expect(
      getSkippedCommitsText({
        commitsWithoutLanguageScope: [],
        commitsWithUnknownLanguageScope: [],
      })
    ).toMatchInlineSnapshot(`"_(None)_"`);

    expect(
      getSkippedCommitsText({
        commitsWithoutLanguageScope: [
          'abcdefg fix: something',
          'abcdefg fix: somethin2',
        ],

        commitsWithUnknownLanguageScope: [
          'abcdef2 fix(pascal): what',
          'abcdef2 fix(pascal): what is that',
        ],
      })
    ).toMatchInlineSnapshot(`
      "</p>
      <p>It doesn't mean these commits are being excluded from the release. It means they're not taken into account when the release process figured out the next version number, and updated the changelog.</p>

      <details>
        <summary>
          <i>Commits without language scope:</i>
        </summary>

        - abcdefg fix: something
      - abcdefg fix: somethin2
      </details>

      <details>
        <summary>
          <i>Commits with unknown language scope:</i>
        </summary>

        - abcdef2 fix(pascal): what
      - abcdef2 fix(pascal): what is that
      </details>"
    `);
  });
});
