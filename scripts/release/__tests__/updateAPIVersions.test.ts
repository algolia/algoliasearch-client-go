import { getVersionsToRelease } from '../updateAPIVersions';

describe('updateAPIversions', () => {
  it('gets versions to release', () => {
    const versions = getVersionsToRelease({
      javascript: {
        current: '1.0.0',
        releaseType: 'minor',
      },
      php: {
        current: '2.0.0',
        releaseType: 'patch',
      },
      java: {
        current: '3.0.0',
        releaseType: null,
        noCommit: true,
        skipRelease: true,
      },
    });

    expect(Object.keys(versions)).toEqual(['javascript', 'php']);
    expect(versions.java).toBeUndefined();
    expect(versions.javascript?.current).toEqual('1.0.0');
    expect(versions.javascript?.releaseType).toEqual('minor');
    expect(versions.php?.current).toEqual('2.0.0');
    expect(versions.php?.releaseType).toEqual('patch');
  });

  it('correctly reads clients version and their next release type', () => {
    const versions = getVersionsToRelease({
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
    });
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
