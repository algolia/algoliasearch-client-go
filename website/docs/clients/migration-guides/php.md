---
title: PHP
---

| Previous             | Latest                                 | Description                                                                                                                                                                                                               |
| -------------------- | :------------------------------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `"algolia/algoliasearch-client-php": "^3.2"`           | `"algolia/algoliasearch-client-php": "^4.0@alpha"`  | **During the beta phase**, the clients are available under the package 4.x.x-alpha , you can find a full list [here](https://packagist.org/packages/algolia/algoliasearch-client-php).      |
| `Algolia\AlgoliaSearch`             | `Algolia\AlgoliaSearch\Api`                         | Exported clients have now the namespace suffixed by `Api`.                                                                                                                                                                                |

### Usage

To get started, first uninstall the previously added clients.

```bash
composer remove algolia/algoliasearch-client-php
```

You can now install the `Algoliasearch` clients.

```bash
composer require algolia/algoliasearch-client-php "^4.0@alpha"
```

You can continue this guide on [our installation page](/docs/clients/installation).

### Methods targeting an `indexName`

Prior to the `initIndex` removal stated in the [common breaking changes](/docs/clients/migration-guides/#common-breaking-changes), all methods previously available at the `initIndex` level requires the `indexName` to be sent with the query.

```php
use Algolia\AlgoliaSearch\Api\SearchClient;

$client = SearchClient::create(
    '<YOUR_APP_ID>',
    '<YOUR_API_KEY>'
);

// only query string
$searchResults = $client->search([
    'requests' => [
        ['indexName' => '<YOUR_INDEX_NAME>', 'query' =>'<YOUR_QUERY>'],
    ],
]);

$searchResults2 = $client->search([
    'requests' => [
        [
            'indexName' => '<YOUR_INDEX_NAME>', 
            'query' => '<YOUR_QUERY>',
            'attributesToRetrieve' => ['firstname', 'lastname'],
            'hitsPerPage' => 50,
        ],
    ],
]);
```
