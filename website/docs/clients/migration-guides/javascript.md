---
title: JavaScript
---

| Previous             | Latest                                 | Description                                                                                                                                                                                                               |
| -------------------- | :------------------------------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `@algolia`           | `@experimental-api-clients-automation` | **During the beta phase**, the clients are available under the NPM `@experimental-api-clients-automation` namespace, you can find a full list [here](https://www.npmjs.com/org/experimental-api-clients-automation).      |
| `algoliasearch/lite` | `algoliasearch-lite`                   | The lite version of the client now have [its own package](https://www.npmjs.com/package/@experimental-api-clients-automation/algoliasearch-lite).                                                                         |
| `search`             | `searchClient`                         | Exported clients are suffixed by `Client`.                                                                                                                                                                                |
| `destroy`            | **removed**                            | This method has not been implemented in the new clients, if you feel the need for it, [please open an issue](https://github.com/algolia/api-clients-automation/issues/new?assignees=&labels=&template=Feature_request.md) |

### Usage

To get started, first install the `algoliasearch` client.

```bash
yarn add @experimental-api-clients-automation/algoliasearch
# or
npm install @experimental-api-clients-automation/algoliasearch
```

You can now uninstall the previously added client.

> Make sure to update all your imports.

```bash
yarn remove algoliasearch
# or
npm uninstall algoliasearch
```

You can continue this guide on [our installation page](/docs/clients/installation).

### Importing algoliasearch using ES Modules

```diff
- import algoliasearch from 'algoliasearch/lite';
+ import { algoliasearchLiteClient } from '@experimental-api-clients-automation/algoliasearch-lite';

- import algoliasearch from 'algoliasearch';
+ import { algoliasearch } from '@experimental-api-clients-automation/algoliasearch';
```

### Methods targeting an `indexName`

Prior to the `initIndex` removal stated in the [common breaking changes](/docs/clients/migration-guides/#common-breaking-changes), all methods previously available at the `initIndex` level requires the `indexName` to be sent with the query.

```js
import { algoliasearch } from '@experimental-api-clients-automation/algoliasearch';

const client = algoliasearch('<YOUR_APP_ID>', '<YOUR_API_KEY>');

// only query string
const searchResults = await client.search({
  requests: [
    {
      indexName: '<YOUR_INDEX_NAME>',
      query: '<YOUR_QUERY>',
    },
  ],
});

// with params
const searchResults2 = await client.search({
  requests: [
    {
      indexName: '<YOUR_INDEX_NAME>',
      query: '<YOUR_QUERY>',
      attributesToRetrieve: ['firstname', 'lastname'],
      hitsPerPage: 50,
    },
  ],
});
```
