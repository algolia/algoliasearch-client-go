import { getVersionChangesText } from '../create-release-issue';
import { getVersionsToRelease } from '../process-release';
import TEXT from '../text';

describe('process release', () => {
  it('gets versions to release', () => {
    const versions = getVersionsToRelease(`
    ## Version Changes
    
    - javascript: 1.0.0 -> **\`minor\` _(e.g. 1.1.0)_**
    - ~java: 3.0.0 -> **\`patch\` _(e.g. 3.0.1)_**~
      - No \`feat\` or \`fix\` commit, thus unchecked by default.
    - php: 2.0.0 -> **\`patch\` _(e.g. 2.0.1)_**
    `);

    expect(Object.keys(versions)).toEqual(['javascript', 'php']);
    expect(versions.javascript?.current).toEqual('1.0.0');
    expect(versions.javascript?.releaseType).toEqual('minor');
    expect(versions.php?.current).toEqual('2.0.0');
    expect(versions.php?.releaseType).toEqual('patch');
  });

  it('parses issue body correctly', () => {
    // This test is a glue between create-release-issue and process-release.
    const issueBody = [
      TEXT.versionChangeHeader,
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
      }),
    ].join('\n');

    const versions = getVersionsToRelease(issueBody);
    expect(versions).toEqual({
      javascript: expect.objectContaining({
        current: '0.0.1',
        releaseType: 'patch',
      }),
      php: expect.objectContaining({
        current: '0.0.1',
        releaseType: 'minor',
      }),
    });
  });
});
