<?php

$env = require_once('../loadEnv.php');

use Algolia\AlgoliaSearch\Api\AnalyticsClient;

$client = AnalyticsClient::create($env['ALGOLIA_APPLICATION_ID'], $env['ALGOLIA_ANALYTICS_KEY']);
$indexName = $env['ANALYTICS_INDEX'];

var_dump(
    $client->getTopFilterForAttribute(
        'myAttribute1,myAttribute2',
        $indexName
    )
);
