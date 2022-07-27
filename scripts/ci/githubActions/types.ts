export type CreateMatrix = {
  /**
   * The name of the branch of reference.
   */
  baseBranch: string;
  /**
   * `true` means we generated the matrix for the `clients` job, `false` for the specs.
   */
  forClients?: boolean;
};

type BaseMatrix = {
  /**
   * Path to the file/folder being handled.
   */
  path: string;
  /**
   * The computed cache key, used to restore files from the CI.
   */
  cacheKey: string;
  /**
   * The list of clients to run in the CI.
   */
  toRun: string;
};

export type ClientMatrix = BaseMatrix & {
  /**
   * The client language.
   */
  language: string;
  /**
   * The root of the test folder.
   */
  testsRootFolder: string;
  /**
   * The test output path to delete before running the CTS generation.
   */
  testsToDelete: string;
  /**
   * The test output path to store in the artifact.
   */
  testsToStore: string;
};

export type SpecMatrix = Pick<BaseMatrix, 'cacheKey' | 'toRun'> & {
  /**
   * The path of the bundled spec file.
   */
  bundledPath: string;
};

export type Matrix<TMatrix> = {
  client: TMatrix[];
};

export type ToRunMatrix = {
  path: string;
  toRun: string[];
  cacheToCompute: string[];
};
