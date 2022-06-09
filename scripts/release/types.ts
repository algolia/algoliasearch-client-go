import type { ReleaseType } from 'semver';

import type { Language } from '../types';

export type Version = {
  current: string;
  releaseType: ReleaseType | null;
  next: string | null;
  skipRelease?: boolean;
  noCommit?: boolean;
};

export type Versions = {
  [lang: string]: Version;
};

export type VersionsBeforeBump = {
  [lang: string]: Omit<Version, 'next' | 'releaseType'>;
};

export type Scope = Language | 'clients' | 'specs';

export type PassedCommit = {
  hash: string;
  type: string;
  /**
   * A commit can be scoped to a language. When scoped to `clients` or `specs`, it impacts all clients.
   */
  scope: Scope;
  message: string;
  raw: string;
};

export type Commit =
  | PassedCommit
  | { error: 'generation-commit' }
  | { error: 'missing-language-scope' }
  | { error: 'unknown-language-scope' };

export type VersionsToRelease = {
  [lang in Language]?: {
    current: string;
    next: string;
    releaseType: ReleaseType;
  };
};

export type Changelog = {
  [lang in Language]?: string;
};

export type BeforeClientCommitCommand = (params: {
  dir: string;
}) => Promise<void>;
