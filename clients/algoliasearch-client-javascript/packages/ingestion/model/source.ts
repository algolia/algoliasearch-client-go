// This file is generated, manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.

import type { SourceInput } from './sourceInput';
import type { SourceType } from './sourceType';

export type Source = {
  /**
   * The source UUID.
   */
  sourceID: string;

  type: SourceType;

  name: string;

  input: SourceInput;

  /**
   * The authentication UUID.
   */
  authenticationID?: string;

  /**
   * Date of creation (RFC3339 format).
   */
  createdAt: string;

  /**
   * Date of last update (RFC3339 format).
   */
  updatedAt?: string;
};
