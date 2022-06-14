<?php

$env = require_once('../loadEnv.php');

use Algolia\AlgoliaSearch\Api\QuerySuggestionsClient;

$client = QuerySuggestionsClient::create($env['ALGOLIA_APPLICATION_ID'], $env['QUERY_SUGGESTIONS_KEY']);

var_dump($client->getAllConfigs());
