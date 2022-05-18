import execa from 'execa';

import { gitCommit } from '../common';
import { getClientsConfigField } from '../config';

jest.mock('execa');

describe('gitCommit', () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  it('commits with message', () => {
    gitCommit({ message: 'chore: does something' });
    expect(execa).toHaveBeenCalledTimes(1);
    expect(execa).toHaveBeenCalledWith(
      'git',
      ['commit', '-m', 'chore: does something'],
      { cwd: expect.any(String) }
    );
  });

  it('commits with co-author', () => {
    // This reflects how it can be retrieved from git commands.
    const author = `Co-authored-by: them <them@algolia.com>
     `.trim();
    const coAuthors = `

      Co-authored-by: me <me@algolia.com>


      Co-authored-by: you <you@algolia.com>
      
      `
      .split('\n')
      .map((coAuthor) => coAuthor.trim())
      .filter(Boolean);

    gitCommit({
      message: 'chore: does something',
      coAuthors: [author, ...coAuthors],
    });
    expect(execa).toHaveBeenCalledTimes(1);
    expect(execa).toHaveBeenCalledWith(
      'git',
      [
        'commit',
        '-m',
        'chore: does something\n\n\nCo-authored-by: them <them@algolia.com>\nCo-authored-by: me <me@algolia.com>\nCo-authored-by: you <you@algolia.com>',
      ],
      { cwd: expect.any(String) }
    );
  });
});

describe('config', () => {
  describe('getClientsConfigField', () => {
    it('throws if the field is not found', () => {
      expect(() => {
        getClientsConfigField('javascript', 'packageVersion');
      }).toThrowErrorMatchingInlineSnapshot(
        `"Unable to find 'packageVersion' for 'javascript'"`
      );

      expect(() => {
        getClientsConfigField('java', 'utilsPackageVersion');
      }).toThrowErrorMatchingInlineSnapshot(
        `"Unable to find 'utilsPackageVersion' for 'java'"`
      );
    });

    it('find the field if it exists', () => {
      expect(getClientsConfigField('java', ['tests', 'extension'])).toEqual(
        '.test.java'
      );
    });
  });
});
