import type { AdvancedSyntaxFeatures } from './advancedSyntaxFeatures';
import type { AlternativesAsExact } from './alternativesAsExact';
import type { ExactOnSingleWordQuery } from './exactOnSingleWordQuery';
import type { QueryType } from './queryType';
import type { RemoveWordsIfNoResults } from './removeWordsIfNoResults';
import type { TypoTolerance } from './typoTolerance';

export type IndexSettingsAsSearchParams = {
  /**
   * The complete list of attributes used for searching.
   */
  searchableAttributes?: string[];
  /**
   * The complete list of attributes that will be used for faceting.
   */
  attributesForFaceting?: string[];
  /**
   * List of attributes that can\'t be retrieved at query time.
   */
  unretrievableAttributes?: string[];
  /**
   * This parameter controls which attributes to retrieve and which not to retrieve.
   */
  attributesToRetrieve?: string[];
  /**
   * Restricts a given query to look in only a subset of your searchable attributes.
   */
  restrictSearchableAttributes?: string[];
  /**
   * Controls how Algolia should sort your results.
   */
  ranking?: string[];
  /**
   * Specifies the custom ranking criterion.
   */
  customRanking?: string[];
  /**
   * Controls the relevancy threshold below which less relevant results aren\'t included in the results.
   */
  relevancyStrictness?: number;
  /**
   * List of attributes to highlight.
   */
  attributesToHighlight?: string[];
  /**
   * List of attributes to snippet, with an optional maximum number of words to snippet.
   */
  attributesToSnippet?: string[];
  /**
   * The HTML string to insert before the highlighted parts in all highlight and snippet results.
   */
  highlightPreTag?: string;
  /**
   * The HTML string to insert after the highlighted parts in all highlight and snippet results.
   */
  highlightPostTag?: string;
  /**
   * String used as an ellipsis indicator when a snippet is truncated.
   */
  snippetEllipsisText?: string;
  /**
   * Restrict highlighting and snippeting to items that matched the query.
   */
  restrictHighlightAndSnippetArrays?: boolean;
  /**
   * Set the number of hits per page.
   */
  hitsPerPage?: number;
  /**
   * Minimum number of characters a word in the query string must contain to accept matches with 1 typo.
   */
  minWordSizefor1Typo?: number;
  /**
   * Minimum number of characters a word in the query string must contain to accept matches with 2 typos.
   */
  minWordSizefor2Typos?: number;
  typoTolerance?: TypoTolerance;
  /**
   * Whether to allow typos on numbers (\"numeric tokens\") in the query string.
   */
  allowTyposOnNumericTokens?: boolean;
  /**
   * List of attributes on which you want to disable typo tolerance.
   */
  disableTypoToleranceOnAttributes?: string[];
  /**
   * Control which separators are indexed.
   */
  separatorsToIndex?: string;
  /**
   * Treats singular, plurals, and other forms of declensions as matching terms.
   */
  ignorePlurals?: string;
  /**
   * Removes stop (common) words from the query before executing it.
   */
  removeStopWords?: string;
  /**
   * List of characters that the engine shouldn\'t automatically normalize.
   */
  keepDiacriticsOnCharacters?: string;
  /**
   * Sets the languages to be used by language-specific settings and functionalities such as ignorePlurals, removeStopWords, and CJK word-detection.
   */
  queryLanguages?: string[];
  /**
   * Splits compound words into their composing atoms in the query.
   */
  decompoundQuery?: boolean;
  /**
   * Whether Rules should be globally enabled.
   */
  enableRules?: boolean;
  /**
   * Enable the Personalization feature.
   */
  enablePersonalization?: boolean;
  queryType?: QueryType;
  removeWordsIfNoResults?: RemoveWordsIfNoResults;
  /**
   * Enables the advanced query syntax.
   */
  advancedSyntax?: boolean;
  /**
   * A list of words that should be considered as optional when found in the query.
   */
  optionalWords?: string[];
  /**
   * List of attributes on which you want to disable the exact ranking criterion.
   */
  disableExactOnAttributes?: string[];
  exactOnSingleWordQuery?: ExactOnSingleWordQuery;
  /**
   * List of alternatives that should be considered an exact match by the exact ranking criterion.
   */
  alternativesAsExact?: AlternativesAsExact[];
  /**
   * Allows you to specify which advanced syntax features are active when ‘advancedSyntax\' is enabled.
   */
  advancedSyntaxFeatures?: AdvancedSyntaxFeatures[];
  /**
   * Enables de-duplication or grouping of results.
   */
  distinct?: number;
  /**
   * Whether to take into account an index\'s synonyms for a particular search.
   */
  synonyms?: boolean;
  /**
   * Whether to highlight and snippet the original word that matches the synonym or the synonym itself.
   */
  replaceSynonymsInHighlight?: boolean;
  /**
   * Precision of the proximity ranking criterion.
   */
  minProximity?: number;
  /**
   * Choose which fields to return in the API response. This parameters applies to search and browse queries.
   */
  responseFields?: string[];
  /**
   * Maximum number of facet hits to return during a search for facet values. For performance reasons, the maximum allowed number of returned values is 100.
   */
  maxFacetHits?: number;
  /**
   * When attribute is ranked above proximity in your ranking formula, proximity is used to select which searchable attribute is matched in the attribute ranking stage.
   */
  attributeCriteriaComputedByMinProximity?: boolean;
  /**
   * Content defining how the search interface should be rendered. Can be set via the settings for a default value and can be overridden via rules.
   */
  renderingContent?: Record<string, any>;
};
