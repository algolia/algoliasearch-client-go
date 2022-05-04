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
   * Name of the client.
   */
  name: string;
  /**
   * Path to the file/folder being handled.
   */
  path: string;
  /**
   * The computed cache key, used to restore files from the CI.
   */
  cacheKey: string;
};

export type ClientMatrix = BaseMatrix & {
  /**
   * The client language.
   */
  language: string;

  /**
   * The client name plus `Config` appended. With the casing corresponding to the language.
   */
  configName: string;
  /**
   * The client name plus `Client` appended. With the casing corresponding to the language.
   */
  apiName: string;

  /**
   * Path to the `API` file/folder of the client, based on the language.
   */
  apiPath: string;
  /**
   * Path to the `Model` file/folder of the client, based on the language.
   */
  modelPath: string;
};

export type SpecMatrix = BaseMatrix & {
  /**
   * The path of the bundled spec file.
   */
  bundledPath: string;
};

export type Matrix<TMatrix> = {
  client: TMatrix[];
};
