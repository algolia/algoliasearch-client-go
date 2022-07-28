// This file is generated, manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.

import type { Variant } from './variant';

export type ABTest = {
  /**
   * The A/B test ID.
   */
  abTestID: number;

  /**
   * A/B test significance based on click data. Should be > 0.95 to be considered significant (no matter which variant is winning).
   */
  clickSignificance: number;

  /**
   * A/B test significance based on conversion data. Should be > 0.95 to be considered significant (no matter which variant is winning).
   */
  conversionSignificance: number;

  /**
   * End date for the A/B test expressed as YYYY-MM-DDThh:mm:ssZ.
   */
  endAt: string;

  /**
   * Update date for the A/B test expressed as YYYY-MM-DDThh:mm:ssZ.
   */
  updatedAt: string;

  /**
   * Creation date for the A/B test expressed as YYYY-MM-DDThh:mm:ssZ.
   */
  createdAt: string;

  /**
   * A/B test name.
   */
  name: string;

  /**
   * Status of the A/B test.
   */
  status: string;

  /**
   * List of A/B test variant.
   */
  variants: Variant[];
};
