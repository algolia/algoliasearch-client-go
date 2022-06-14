<?php

$env = require_once('../loadEnv.php');

use Algolia\AlgoliaSearch\Api\InsightsClient;

$client = InsightsClient::create($env['ALGOLIA_APPLICATION_ID'], $env['ALGOLIA_ADMIN_KEY']);
$indexName = $env['SEARCH_INDEX'];

$twoDaysAgoMs = (time() - (2 * 24 * 60 * 60)) * 1000;

$event = [
    'eventType' => 'click',
    'eventName' => 'foo',
    'index' => 'sending_events',
    'userToken' => 'bar',
    'objectIDs' => ['one', 'two'],
    'timestamp' => $twoDaysAgoMs,
];

var_dump(
    $client->pushEvents([$event])
);
