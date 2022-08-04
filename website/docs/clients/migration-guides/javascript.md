---
title: JavaScript
---

| Previous  | Latest         | Description                                                                                                                                                                                                               |
| --------- | :------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `search`  | `searchClient` | Exported clients are suffixed by `Client`.                                                                                                                                                                                |
| `destroy` | **removed**    | This method has not been implemented in the new clients, if you feel the need for it, [please open an issue](https://github.com/algolia/api-clients-automation/issues/new?assignees=&labels=&template=Feature_request.md) |

### Usage

To get started, first install the `algoliasearch` client.

```bash
yarn add algoliasearch@alpha
# or
npm install algoliasearch@alpha
```

You can continue this guide on [our installation page](/docs/clients/installation).

### Methods targeting an `indexName`

Prior to the `initIndex` removal stated in the [common breaking changes](/docs/clients/migration-guides/#common-breaking-changes), all methods previously available at the `initIndex` level requires the `indexName` to be sent with the query.

```js
import { algoliasearch } from 'algoliasearch';

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
