import type { SearchTypeFacet } from './searchTypeFacet';

export type SearchForFacetsOptions = {
  /**
   * The `facet` name.
   */
  facet: string;
  /**
   * The Algolia index name.
   */
  indexName: string;
  /**
   * Text to search inside the facet\'s values.
   */
  facetQuery?: string;
  /**
   * Maximum number of facet hits to return during a search for facet values. For performance reasons, the maximum allowed number of returned values is 100.
   */
  maxFacetHits?: number;
  type: SearchTypeFacet;
};