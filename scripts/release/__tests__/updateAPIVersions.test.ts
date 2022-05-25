import { getVersionChangesText } from '../createReleasePR';
import TEXT from '../text';
import { getVersionsToRelease } from '../updateAPIVersions';

describe('updateAPIversions', () => {
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

  it('correctly reads clients version and their next release type', () => {
    // This test is a glue between createReleasePR and updateAPIVersions.
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
