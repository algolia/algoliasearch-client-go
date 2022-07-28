// This file is generated, manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.

import type { FacetsStats } from './facetsStats';
import type { RenderingContent } from './renderingContent';

export type BaseSearchResponse = {
  /**
   * If a search encounters an index that is being A/B tested, abTestID reports the ongoing A/B test ID.
   */
  abTestID?: number;

  /**
   * If a search encounters an index that is being A/B tested, abTestVariantID reports the variant ID of the index used (starting at 1).
   */
  abTestVariantID?: number;

  /**
   * The computed geo location.
   */
  aroundLatLng?: string;

  /**
   * The automatically computed radius. For legacy reasons, this parameter is a string and not an integer.
   */
  automaticRadius?: string;

  /**
   * Whether the facet count is exhaustive or approximate.
   */
  exhaustiveFacetsCount?: boolean;

  /**
   * Indicate if the nbHits count was exhaustive or approximate.
   */
  exhaustiveNbHits: boolean;

  /**
   * Indicate if the typo-tolerance search was exhaustive or approximate (only included when typo-tolerance is enabled).
   */
  exhaustiveTypo: boolean;

  /**
   * A mapping of each facet name to the corresponding facet counts.
   */
  facets?: Record<string, Record<string, number>>;

  /**
   * Statistics for numerical facets.
   */
  facets_stats?: Record<string, FacetsStats>;

  /**
   * Set the number of hits per page.
   */
  hitsPerPage: number;

  /**
   * Index name used for the query.
   */
  index?: string;

  /**
   * Index name used for the query. In the case of an A/B test, the targeted index isn\'t always the index used by the query.
   */
  indexUsed?: string;

  /**
   * Used to return warnings about the query.
   */
  message?: string;

  /**
   * Number of hits that the search query matched.
   */
  nbHits: number;

  /**
   * Number of pages available for the current query.
   */
  nbPages: number;

  /**
   * The number of hits selected and sorted by the relevant sort algorithm.
   */
  nbSortedHits?: number;

  /**
   * Specify the page to retrieve.
   */
  page: number;

  /**
   * A url-encoded string of all search parameters.
   */
  params: string;

  /**
   * The query string that will be searched, after normalization.
   */
  parsedQuery?: string;

  /**
   * Time the server took to process the request, in milliseconds.
   */
  processingTimeMS: number;

  /**
   * The text to search in the index.
   */
  query: string;

  /**
   * A markup text indicating which parts of the original query have been removed in order to retrieve a non-empty result set.
   */
  queryAfterRemoval?: string;

  /**
   * Actual host name of the server that processed the request.
   */
  serverUsed?: string;

  /**
   * Lets you store custom data in your indices.
   */
  userData?: Record<string, any>;

  renderingContent?: RenderingContent;
};
