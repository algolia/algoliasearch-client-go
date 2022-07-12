// This file is generated, manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.

import type { LogQuery } from './logQuery';

export type Log = {
  /**
   * Timestamp in ISO-8601 format.
   */
  timestamp: string;
  /**
   * HTTP method of the performed request.
   */
  method: string;
  /**
   * HTTP response code.
   */
  answer_code: string;
  /**
   * Request body. Truncated after 1000 characters.
   */
  query_body: string;
  /**
   * Answer body. Truncated after 1000 characters.
   */
  answer: string;
  /**
   * Request URL.
   */
  url: string;
  /**
   * IP of the client which performed the request.
   */
  ip: string;
  /**
   * Request Headers (API Key is obfuscated).
   */
  query_headers: string;
  /**
   * SHA1 signature of the log entry.
   */
  sha1: string;
  /**
   * Number of API calls.
   */
  nb_api_calls: string;
  /**
   * Processing time for the query. It doesn\'t include network time.
   */
  processing_time_ms: string;
  /**
   * Index targeted by the query.
   */
  index?: string;
  /**
   * Query parameters sent with the request.
   */
  query_params?: string;
  /**
   * Number of hits returned for the query.
   */
  query_nb_hits?: string;
  /**
   * Array of all performed queries for the given request.
   */
  inner_queries?: LogQuery[];
};
