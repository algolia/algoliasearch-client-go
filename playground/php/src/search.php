<?php

$env = require_once('../loadEnv.php');

use Algolia\AlgoliaSearch\Api\SearchClient;

$client = SearchClient::create(
    $env['ALGOLIA_APPLICATION_ID'],
    $env['ALGOLIA_ADMIN_KEY']
);
$indexName = $env['SEARCH_INDEX'];


$response = $client->saveObject(
    $indexName,
    ['objectID' => "111", 'name' => $env['SEARCH_QUERY']],
);

var_dump($response);

$client->waitForTask($indexName, $response['taskID']);

var_dump(
    $client->search([
        'requests' => [
            ['indexName' => $indexName, 'query' => $env['SEARCH_QUERY']],
        ],
    ])
);
