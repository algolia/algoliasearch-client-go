// This file is generated, manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.

import type { MatchLevel } from './matchLevel';

/**
 * Show highlighted section and words matched on a query.
 */
export type HighlightResultOption = {
  /**
   * Markup text with occurrences highlighted.
   */
  value: string;

  matchLevel: MatchLevel;

  /**
   * List of words from the query that matched the object.
   */
  matchedWords: string[];

  /**
   * Whether the entire attribute value is highlighted.
   */
  fullyHighlighted?: boolean;
};