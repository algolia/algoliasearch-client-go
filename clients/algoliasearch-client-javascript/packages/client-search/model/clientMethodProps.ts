// This file is generated, manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.

import type { CreateIterablePromise } from '@algolia/client-common';

import type { ApiKey } from './apiKey';
import type { AssignUserIdParams } from './assignUserIdParams';
import type { AttributeOrBuiltInOperation } from './attributeOrBuiltInOperation';
import type { BatchAssignUserIdsParams } from './batchAssignUserIdsParams';
import type { BatchDictionaryEntriesParams } from './batchDictionaryEntriesParams';
import type { BatchWriteParams } from './batchWriteParams';
import type { BrowseRequest } from './browseRequest';
import type { DictionaryType } from './dictionaryType';
import type { GetTaskResponse } from './getTaskResponse';
import type { IndexSettings } from './indexSettings';
import type { Key } from './key';
import type { LogType } from './logType';
import type { OperationIndexParams } from './operationIndexParams';
import type { Rule } from './rule';
import type { SearchDictionaryEntriesParams } from './searchDictionaryEntriesParams';
import type { SearchForFacetValuesRequest } from './searchForFacetValuesRequest';
import type { SearchForFacetsOptions } from './searchForFacetsOptions';
import type { SearchForHitsOptions } from './searchForHitsOptions';
import type { SearchParams } from './searchParams';
import type { SearchParamsObject } from './searchParamsObject';
import type { SearchRulesParams } from './searchRulesParams';
import type { SearchSynonymsParams } from './searchSynonymsParams';
import type { Source } from './source';
import type { SynonymHit } from './synonymHit';
import type { SynonymType } from './synonymType';

/**
 * Properties for the `addOrUpdateObject` method.
 */
export type AddOrUpdateObjectProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
  /**
   * Unique identifier of an object.
   */
  objectID: string;
  /**
   * The Algolia object.
   */
  body: Record<string, any>;
};

/**
 * Properties for the `assignUserId` method.
 */
export type AssignUserIdProps = {
  /**
   * UserID to assign.
   */
  xAlgoliaUserID: string;
  assignUserIdParams: AssignUserIdParams;
};

/**
 * Properties for the `batch` method.
 */
export type BatchProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
  batchWriteParams: BatchWriteParams;
};

/**
 * Properties for the `batchAssignUserIds` method.
 */
export type BatchAssignUserIdsProps = {
  /**
   * UserID to assign.
   */
  xAlgoliaUserID: string;
  batchAssignUserIdsParams: BatchAssignUserIdsParams;
};

/**
 * Properties for the `batchDictionaryEntries` method.
 */
export type BatchDictionaryEntriesProps = {
  /**
   * The dictionary to search in.
   */
  dictionaryName: DictionaryType;
  batchDictionaryEntriesParams: BatchDictionaryEntriesParams;
};

/**
 * Properties for the `browse` method.
 */
export type BrowseProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
  browseRequest?: BrowseRequest;
};

/**
 * Properties for the `clearAllSynonyms` method.
 */
export type ClearAllSynonymsProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
  /**
   * When true, changes are also propagated to replicas of the given indexName.
   */
  forwardToReplicas?: boolean;
};

/**
 * Properties for the `clearObjects` method.
 */
export type ClearObjectsProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
};

/**
 * Properties for the `clearRules` method.
 */
export type ClearRulesProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
  /**
   * When true, changes are also propagated to replicas of the given indexName.
   */
  forwardToReplicas?: boolean;
};

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
 * Properties for the `deleteApiKey` method.
 */
export type DeleteApiKeyProps = {
  /**
   * API Key string.
   */
  key: string;
};

/**
 * Properties for the `deleteBy` method.
 */
export type DeleteByProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
  searchParams: SearchParams;
};

/**
 * Properties for the `deleteIndex` method.
 */
export type DeleteIndexProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
};

/**
 * Properties for the `deleteObject` method.
 */
export type DeleteObjectProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
  /**
   * Unique identifier of an object.
   */
  objectID: string;
};

/**
 * Properties for the `deleteRule` method.
 */
export type DeleteRuleProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
  /**
   * Unique identifier of an object.
   */
  objectID: string;
  /**
   * When true, changes are also propagated to replicas of the given indexName.
   */
  forwardToReplicas?: boolean;
};

/**
 * Properties for the `deleteSource` method.
 */
export type DeleteSourceProps = {
  /**
   * The IP range of the source.
   */
  source: string;
};

/**
 * Properties for the `deleteSynonym` method.
 */
export type DeleteSynonymProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
  /**
   * Unique identifier of an object.
   */
  objectID: string;
  /**
   * When true, changes are also propagated to replicas of the given indexName.
   */
  forwardToReplicas?: boolean;
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
 * Properties for the `getApiKey` method.
 */
export type GetApiKeyProps = {
  /**
   * API Key string.
   */
  key: string;
};

/**
 * Properties for the `getLogs` method.
 */
export type GetLogsProps = {
  /**
   * First entry to retrieve (zero-based). Log entries are sorted by decreasing date, therefore 0 designates the most recent log entry.
   */
  offset?: number;
  /**
   * Maximum number of entries to retrieve. The maximum allowed value is 1000.
   */
  length?: number;
  /**
   * Index for which log entries should be retrieved. When omitted, log entries are retrieved across all indices.
   */
  indexName?: string;
  /**
   * Type of log entries to retrieve. When omitted, all log entries are retrieved.
   */
  type?: LogType;
};

/**
 * Properties for the `getObject` method.
 */
export type GetObjectProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
  /**
   * Unique identifier of an object.
   */
  objectID: string;
  /**
   * List of attributes to retrieve. If not specified, all retrievable attributes are returned.
   */
  attributesToRetrieve?: string[];
};

/**
 * Properties for the `getRule` method.
 */
export type GetRuleProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
  /**
   * Unique identifier of an object.
   */
  objectID: string;
};

/**
 * Properties for the `getSettings` method.
 */
export type GetSettingsProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
};

/**
 * Properties for the `getSynonym` method.
 */
export type GetSynonymProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
  /**
   * Unique identifier of an object.
   */
  objectID: string;
};

/**
 * Properties for the `getTask` method.
 */
export type GetTaskProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
  /**
   * Unique identifier of an task. Numeric value (up to 64bits).
   */
  taskID: number;
};

/**
 * Properties for the `getUserId` method.
 */
export type GetUserIdProps = {
  /**
   * UserID to assign.
   */
  userID: string;
};

/**
 * Properties for the `hasPendingMappings` method.
 */
export type HasPendingMappingsProps = {
  /**
   * If the clusters pending mapping state should be on the response.
   */
  getClusters?: boolean;
};

/**
 * Properties for the `listIndices` method.
 */
export type ListIndicesProps = {
  /**
   * Requested page (zero-based). When specified, will retrieve a specific page; the page size is implicitly set to 100. When null, will retrieve all indices (no pagination).
   */
  page?: number;
};

/**
 * Properties for the `listUserIds` method.
 */
export type ListUserIdsProps = {
  /**
   * Requested page (zero-based). When specified, will retrieve a specific page; the page size is implicitly set to 100. When null, will retrieve all indices (no pagination).
   */
  page?: number;
  /**
   * Maximum number of objects to retrieve.
   */
  hitsPerPage?: number;
};

/**
 * Properties for the `operationIndex` method.
 */
export type OperationIndexProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
  operationIndexParams: OperationIndexParams;
};

/**
 * Properties for the `partialUpdateObject` method.
 */
export type PartialUpdateObjectProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
  /**
   * Unique identifier of an object.
   */
  objectID: string;
  /**
   * List of attributes to update.
   */
  attributeOrBuiltInOperation: Array<
    Record<string, AttributeOrBuiltInOperation>
  >;
  /**
   * Creates the record if it does not exist yet.
   */
  createIfNotExists?: boolean;
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
 * Properties for the `removeUserId` method.
 */
export type RemoveUserIdProps = {
  /**
   * UserID to assign.
   */
  userID: string;
};

/**
 * Properties for the `replaceSources` method.
 */
export type ReplaceSourcesProps = {
  /**
   * The sources to allow.
   */
  source: Source[];
};

/**
 * Properties for the `restoreApiKey` method.
 */
export type RestoreApiKeyProps = {
  /**
   * API Key string.
   */
  key: string;
};

/**
 * Properties for the `saveObject` method.
 */
export type SaveObjectProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
  /**
   * The Algolia record.
   */
  body: Record<string, any>;
};

/**
 * Properties for the `saveRule` method.
 */
export type SaveRuleProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
  /**
   * Unique identifier of an object.
   */
  objectID: string;
  rule: Rule;
  /**
   * When true, changes are also propagated to replicas of the given indexName.
   */
  forwardToReplicas?: boolean;
};

/**
 * Properties for the `saveRules` method.
 */
export type SaveRulesProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
  rule: Rule[];
  /**
   * When true, changes are also propagated to replicas of the given indexName.
   */
  forwardToReplicas?: boolean;
  /**
   * When true, existing Rules are cleared before adding this batch. When false, existing Rules are kept.
   */
  clearExistingRules?: boolean;
};

/**
 * Properties for the `saveSynonym` method.
 */
export type SaveSynonymProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
  /**
   * Unique identifier of an object.
   */
  objectID: string;
  synonymHit: SynonymHit;
  /**
   * When true, changes are also propagated to replicas of the given indexName.
   */
  forwardToReplicas?: boolean;
};

/**
 * Properties for the `saveSynonyms` method.
 */
export type SaveSynonymsProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
  synonymHit: SynonymHit[];
  /**
   * When true, changes are also propagated to replicas of the given indexName.
   */
  forwardToReplicas?: boolean;
  /**
   * Replace all synonyms of the index with the ones sent with this request.
   */
  replaceExistingSynonyms?: boolean;
};

/**
 * In v4, the search parameters are wrapped in a `params` parameter.
 *
 * @deprecated The `search` method now accepts flat `searchParams` at the root of the method.
 */
type LegacySearchParams = {
  params?: SearchParamsObject;
};

/**
 * In v4, the search parameters are wrapped in a `params` parameter.
 *
 * @deprecated The `search` method now accepts flat `searchParams` at the root of the method.
 */
type LegacySearchForFacets = LegacySearchParams & SearchForFacetsOptions;

/**
 * In v4, the search parameters are wrapped in a `params` parameter.
 *
 * @deprecated The `search` method now accepts flat `searchParams` at the root of the method.
 */
type LegacySearchForHits = LegacySearchParams & SearchForHitsOptions;

type LegacySearchQuery = LegacySearchForFacets | LegacySearchForHits;

/**
 * Search method signature compatible with the `algoliasearch` v4 package. When using this signature, extra computation will be required to make it match the new signature.
 *
 * @deprecated This signature will be removed from the next major version, we recommend using the `SearchMethodParams` type for performances and future proof reasons.
 */
export type LegacySearchMethodProps = LegacySearchQuery[];

/**
 * Properties for the `searchDictionaryEntries` method.
 */
export type SearchDictionaryEntriesProps = {
  /**
   * The dictionary to search in.
   */
  dictionaryName: DictionaryType;
  searchDictionaryEntriesParams: SearchDictionaryEntriesParams;
};

/**
 * Properties for the `searchForFacetValues` method.
 */
export type SearchForFacetValuesProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
  /**
   * The facet name.
   */
  facetName: string;
  searchForFacetValuesRequest?: SearchForFacetValuesRequest;
};

/**
 * Properties for the `searchRules` method.
 */
export type SearchRulesProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
  searchRulesParams?: SearchRulesParams;
};

/**
 * Properties for the `searchSingleIndex` method.
 */
export type SearchSingleIndexProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
  searchParams?: SearchParams;
};

/**
 * Properties for the `searchSynonyms` method.
 */
export type SearchSynonymsProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
  /**
   * Only search for specific types of synonyms.
   */
  type?: SynonymType;
  /**
   * Requested page (zero-based). When specified, will retrieve a specific page; the page size is implicitly set to 100. When null, will retrieve all indices (no pagination).
   */
  page?: number;
  /**
   * Maximum number of objects to retrieve.
   */
  hitsPerPage?: number;
  /**
   * The body of the the `searchSynonyms` method.
   */
  searchSynonymsParams?: SearchSynonymsParams;
};

/**
 * Properties for the `setSettings` method.
 */
export type SetSettingsProps = {
  /**
   * The index in which to perform the request.
   */
  indexName: string;
  indexSettings: IndexSettings;
  /**
   * When true, changes are also propagated to replicas of the given indexName.
   */
  forwardToReplicas?: boolean;
};

/**
 * Properties for the `updateApiKey` method.
 */
export type UpdateApiKeyProps = {
  /**
   * API Key string.
   */
  key: string;
  apiKey: ApiKey;
};

/**
 * The `browseObjects`, `browseRules`, `browseSynonyms` options.
 */
export type BrowseOptions<T> = Partial<
  Pick<CreateIterablePromise<T>, 'validate'>
> &
  Required<Pick<CreateIterablePromise<T>, 'aggregator'>>;

type WaitForOptions<T> = Omit<
  CreateIterablePromise<T>,
  'func' | 'timeout' | 'validate'
> & {
  /**
   * The maximum number of retries. 50 by default.
   */
  maxRetries: number;

  /**
   * The function to decide how long to wait between retries.
   */
  timeout: (retryCount: number) => number;
};

export type WaitForTaskOptions = WaitForOptions<GetTaskResponse> & {
  /**
   * The `indexName` where the operation was performed.
   */
  indexName: string;
  /**
   * The `taskID` returned by the method response.
   */
  taskID: number;
};

export type WaitForApiKeyOptions = WaitForOptions<Key> & {
  /**
   * The API Key.
   */
  key: string;
} & (
    | {
        /**
         * The operation that has been performed, used to compute the stop condition.
         */
        operation: 'add' | 'delete';
        apiKey?: never;
      }
    | {
        /**
         * The operation that has been performed, used to compute the stop condition.
         */
        operation: 'update';
        /**
         * The updated fields, used to compute the stop condition.
         */
        apiKey: Partial<ApiKey>;
      }
  );
