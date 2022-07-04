import type { BaseSearchParamsWithoutQuery } from './baseSearchParamsWithoutQuery';
import type { SearchParamsQuery } from './searchParamsQuery';

export type BaseSearchParams = BaseSearchParamsWithoutQuery & SearchParamsQuery;
