// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

export type SourceCommercetools = {
  storeKeys?: string[];

  /**
   * Array of locales that must match the following pattern: ^[a-z]{2}(-[A-Z]{2})?$. For example [\"fr-FR\", \"en\"].
   */
  locales?: string[];

  url: string;

  projectKey: string;
};