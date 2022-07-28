// This file is generated, manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.

import type { HighlightResult } from './highlightResult';
import type { RankingInfo } from './rankingInfo';
import type { SnippetResult } from './snippetResult';

/**
 * A single hit.
 */
export type Hit<T = Record<string, any>> = T & {
  /**
   * Unique identifier of the object.
   */
  objectID: string;

  /**
   * Show highlighted section and words matched on a query.
   */
  _highlightResult?: Record<string, HighlightResult>;

  /**
   * Snippeted attributes show parts of the matched attributes. Only returned when attributesToSnippet is non-empty.
   */
  _snippetResult?: Record<string, SnippetResult>;

  _rankingInfo?: RankingInfo;

  _distinctSeqID?: number;
};
