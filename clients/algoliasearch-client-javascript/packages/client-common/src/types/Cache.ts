export type Cache = {
  /**
   * Gets the value of the given `key`.
   */
  get: <TValue>(
    key: Record<string, any> | string,
    defaultValue: () => Promise<TValue>,
    events?: CacheEvents<TValue>
  ) => Promise<TValue>;

  /**
   * Sets the given value with the given `key`.
   */
  set: <TValue>(
    key: Record<string, any> | string,
    value: TValue
  ) => Promise<TValue>;

  /**
   * Deletes the given `key`.
   */
  delete: (key: Record<string, any> | string) => Promise<void>;

  /**
   * Clears the cache.
   */
  clear: () => Promise<void>;
};

export type CacheEvents<TValue> = {
  /**
   * The callback when the given `key` is missing from the cache.
   */
  miss: (value: TValue) => Promise<any>;
};

export type MemoryCacheOptions = {
  /**
   * If keys and values should be serialized using `JSON.stringify`.
   */
  serializable?: boolean;
};

export type BrowserLocalStorageOptions = {
  /**
   * The cache key.
   */
  key: string;

  /**
   * The native local storage implementation.
   */
  localStorage?: Storage;
};

export type FallbackableCacheOptions = {
  /**
   * List of caches order by priority.
   */
  caches: Cache[];
};
