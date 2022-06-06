import type { QuerySuggestionsIndexParam } from './querySuggestionsIndexParam';

/**
 * Properties for the `del` method.
 */
export type DelProps = {
  /**
   * The path of the API endpoint to target, anything after the /1 needs to be specified.
   */
  path: string;
  /**
   * Query parameters to be applied to the current query.
   */
  parameters?: Record<string, any>;
};

/**
 * Properties for the `deleteConfig` method.
 */
export type DeleteConfigProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
};

/**
 * Properties for the `get` method.
 */
export type GetProps = {
  /**
   * The path of the API endpoint to target, anything after the /1 needs to be specified.
   */
  path: string;
  /**
   * Query parameters to be applied to the current query.
   */
  parameters?: Record<string, any>;
};

/**
 * Properties for the `getConfig` method.
 */
export type GetConfigProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
};

/**
 * Properties for the `getConfigStatus` method.
 */
export type GetConfigStatusProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
};

/**
 * Properties for the `getLogFile` method.
 */
export type GetLogFileProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
};

/**
 * Properties for the `post` method.
 */
export type PostProps = {
  /**
   * The path of the API endpoint to target, anything after the /1 needs to be specified.
   */
  path: string;
  /**
   * Query parameters to be applied to the current query.
   */
  parameters?: Record<string, any>;
  /**
   * The parameters to send with the custom request.
   */
  body?: Record<string, any>;
};

/**
 * Properties for the `put` method.
 */
export type PutProps = {
  /**
   * The path of the API endpoint to target, anything after the /1 needs to be specified.
   */
  path: string;
  /**
   * Query parameters to be applied to the current query.
   */
  parameters?: Record<string, any>;
  /**
   * The parameters to send with the custom request.
   */
  body?: Record<string, any>;
};

/**
 * Properties for the `updateConfig` method.
 */
export type UpdateConfigProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
  querySuggestionsIndexParam: QuerySuggestionsIndexParam;
};
