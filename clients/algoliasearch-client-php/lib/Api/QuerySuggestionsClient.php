<?php

// This file is generated, manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.

namespace Algolia\AlgoliaSearch\Api;

use Algolia\AlgoliaSearch\Algolia;
use Algolia\AlgoliaSearch\Configuration\QuerySuggestionsConfig;
use Algolia\AlgoliaSearch\ObjectSerializer;
use Algolia\AlgoliaSearch\RetryStrategy\ApiWrapper;
use Algolia\AlgoliaSearch\RetryStrategy\ApiWrapperInterface;
use Algolia\AlgoliaSearch\RetryStrategy\ClusterHosts;

/**
 * QuerySuggestionsClient Class Doc Comment
 *
 * @category Class
 * @package  Algolia\AlgoliaSearch
 */
class QuerySuggestionsClient
{
    /**
     * @var ApiWrapperInterface
     */
    protected $api;

    /**
     * @var QuerySuggestionsConfig
     */
    protected $config;

    /**
     * @param QuerySuggestionsConfig $config
     * @param ApiWrapperInterface $apiWrapper
     */
    public function __construct(
        ApiWrapperInterface $apiWrapper,
        QuerySuggestionsConfig $config
    ) {
        $this->config = $config;
        $this->api = $apiWrapper;
    }

    /**
     * Instantiate the client with basic credentials and region
     *
     * @param string $appId  Application ID
     * @param string $apiKey Algolia API Key
     * @param string $region Region
     */
    public static function create($appId = null, $apiKey = null, $region = null)
    {
        $allowedRegions = ['eu', 'us'];

        if (
            $region === null ||
            ($region !== null && !in_array($region, $allowedRegions, true))
        ) {
            throw new AlgoliaException(
                '`region` is required and must be one of the following: ' .
                    implode(', ', $allowedRegions)
            );
        }

        $config = QuerySuggestionsConfig::create($appId, $apiKey, $region);

        return static::createWithConfig($config);
    }

    /**
     * Instantiate the client with configuration
     *
     * @param QuerySuggestionsConfig $config Configuration
     */
    public static function createWithConfig(QuerySuggestionsConfig $config)
    {
        $config = clone $config;

        $apiWrapper = new ApiWrapper(
            Algolia::getHttpClient(),
            $config,
            self::getClusterHosts($config)
        );

        return new static($apiWrapper, $config);
    }

    /**
     * Gets the cluster hosts depending on the config
     *
     * @param QuerySuggestionsConfig $config
     *
     * @return ClusterHosts
     */
    public static function getClusterHosts(QuerySuggestionsConfig $config)
    {
        if ($hosts = $config->getHosts()) {
            // If a list of hosts was passed, we ignore the cache
            $clusterHosts = ClusterHosts::create($hosts);
        } else {
            $url =
                $config->getRegion() !== null && $config->getRegion() !== ''
                    ? str_replace(
                        '{region}',
                        $config->getRegion(),
                        'query-suggestions.{region}.algolia.com'
                    )
                    : '';
            $clusterHosts = ClusterHosts::create($url);
        }

        return $clusterHosts;
    }

    /**
     * @return QuerySuggestionsConfig
     */
    public function getClientConfig()
    {
        return $this->config;
    }

    /**
     * Create a configuration.
     *
     * @param array $querySuggestionsIndexWithIndexParam querySuggestionsIndexWithIndexParam (required)
     *
     * @see \Algolia\AlgoliaSearch\Model\QuerySuggestions\QuerySuggestionsIndexWithIndexParam
     *
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\QuerySuggestions\SuccessResponse
     */
    public function createConfig(
        $querySuggestionsIndexWithIndexParam,
        $requestOptions = []
    ) {
        // verify the required parameter 'querySuggestionsIndexWithIndexParam' is set
        if (!isset($querySuggestionsIndexWithIndexParam)) {
            throw new \InvalidArgumentException(
                'Parameter `querySuggestionsIndexWithIndexParam` is required when calling `createConfig`.'
            );
        }

        $resourcePath = '/1/configs';
        $queryParameters = [];
        $headers = [];
        $httpBody = $querySuggestionsIndexWithIndexParam;

        return $this->sendRequest(
            'POST',
            $resourcePath,
            $headers,
            $queryParameters,
            $httpBody,
            $requestOptions
        );
    }

    /**
     * Send requests to the Algolia REST API.
     *
     * @param string $path The path of the API endpoint to target, anything after the /1 needs to be specified. (required)
     * @param array $parameters Query parameters to be applied to the current query. (optional)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|object
     */
    public function del($path, $parameters = null, $requestOptions = [])
    {
        // verify the required parameter 'path' is set
        if (!isset($path)) {
            throw new \InvalidArgumentException(
                'Parameter `path` is required when calling `del`.'
            );
        }

        $resourcePath = '/1{path}';
        $queryParameters = [];
        $headers = [];
        $httpBody = null;

        if ($parameters !== null) {
            $queryParameters = $parameters;
        }

        // path params
        if ($path !== null) {
            $resourcePath = str_replace('{path}', $path, $resourcePath);
        }

        return $this->sendRequest(
            'DELETE',
            $resourcePath,
            $headers,
            $queryParameters,
            $httpBody,
            $requestOptions
        );
    }

    /**
     * Delete a configuration.
     *
     * @param string $indexName The index in which to perform the request. (required)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\QuerySuggestions\SuccessResponse
     */
    public function deleteConfig($indexName, $requestOptions = [])
    {
        // verify the required parameter 'indexName' is set
        if (!isset($indexName)) {
            throw new \InvalidArgumentException(
                'Parameter `indexName` is required when calling `deleteConfig`.'
            );
        }

        $resourcePath = '/1/configs/{indexName}';
        $queryParameters = [];
        $headers = [];
        $httpBody = null;

        // path params
        if ($indexName !== null) {
            $resourcePath = str_replace(
                '{indexName}',
                ObjectSerializer::toPathValue($indexName),
                $resourcePath
            );
        }

        return $this->sendRequest(
            'DELETE',
            $resourcePath,
            $headers,
            $queryParameters,
            $httpBody,
            $requestOptions
        );
    }

    /**
     * Send requests to the Algolia REST API.
     *
     * @param string $path The path of the API endpoint to target, anything after the /1 needs to be specified. (required)
     * @param array $parameters Query parameters to be applied to the current query. (optional)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|object
     */
    public function get($path, $parameters = null, $requestOptions = [])
    {
        // verify the required parameter 'path' is set
        if (!isset($path)) {
            throw new \InvalidArgumentException(
                'Parameter `path` is required when calling `get`.'
            );
        }

        $resourcePath = '/1{path}';
        $queryParameters = [];
        $headers = [];
        $httpBody = null;

        if ($parameters !== null) {
            $queryParameters = $parameters;
        }

        // path params
        if ($path !== null) {
            $resourcePath = str_replace('{path}', $path, $resourcePath);
        }

        return $this->sendRequest(
            'GET',
            $resourcePath,
            $headers,
            $queryParameters,
            $httpBody,
            $requestOptions
        );
    }

    /**
     * List configurations.
     *
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\QuerySuggestions\QuerySuggestionsIndex[]
     */
    public function getAllConfigs($requestOptions = [])
    {
        $resourcePath = '/1/configs';
        $queryParameters = [];
        $headers = [];
        $httpBody = null;

        return $this->sendRequest(
            'GET',
            $resourcePath,
            $headers,
            $queryParameters,
            $httpBody,
            $requestOptions
        );
    }

    /**
     * Get a single configuration.
     *
     * @param string $indexName The index in which to perform the request. (required)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\QuerySuggestions\QuerySuggestionsIndex
     */
    public function getConfig($indexName, $requestOptions = [])
    {
        // verify the required parameter 'indexName' is set
        if (!isset($indexName)) {
            throw new \InvalidArgumentException(
                'Parameter `indexName` is required when calling `getConfig`.'
            );
        }

        $resourcePath = '/1/configs/{indexName}';
        $queryParameters = [];
        $headers = [];
        $httpBody = null;

        // path params
        if ($indexName !== null) {
            $resourcePath = str_replace(
                '{indexName}',
                ObjectSerializer::toPathValue($indexName),
                $resourcePath
            );
        }

        return $this->sendRequest(
            'GET',
            $resourcePath,
            $headers,
            $queryParameters,
            $httpBody,
            $requestOptions
        );
    }

    /**
     * Get configuration status.
     *
     * @param string $indexName The index in which to perform the request. (required)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\QuerySuggestions\Status
     */
    public function getConfigStatus($indexName, $requestOptions = [])
    {
        // verify the required parameter 'indexName' is set
        if (!isset($indexName)) {
            throw new \InvalidArgumentException(
                'Parameter `indexName` is required when calling `getConfigStatus`.'
            );
        }

        $resourcePath = '/1/configs/{indexName}/status';
        $queryParameters = [];
        $headers = [];
        $httpBody = null;

        // path params
        if ($indexName !== null) {
            $resourcePath = str_replace(
                '{indexName}',
                ObjectSerializer::toPathValue($indexName),
                $resourcePath
            );
        }

        return $this->sendRequest(
            'GET',
            $resourcePath,
            $headers,
            $queryParameters,
            $httpBody,
            $requestOptions
        );
    }

    /**
     * Get a log file.
     *
     * @param string $indexName The index in which to perform the request. (required)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\QuerySuggestions\LogFile[]
     */
    public function getLogFile($indexName, $requestOptions = [])
    {
        // verify the required parameter 'indexName' is set
        if (!isset($indexName)) {
            throw new \InvalidArgumentException(
                'Parameter `indexName` is required when calling `getLogFile`.'
            );
        }

        $resourcePath = '/1/logs/{indexName}';
        $queryParameters = [];
        $headers = [];
        $httpBody = null;

        // path params
        if ($indexName !== null) {
            $resourcePath = str_replace(
                '{indexName}',
                ObjectSerializer::toPathValue($indexName),
                $resourcePath
            );
        }

        return $this->sendRequest(
            'GET',
            $resourcePath,
            $headers,
            $queryParameters,
            $httpBody,
            $requestOptions
        );
    }

    /**
     * Send requests to the Algolia REST API.
     *
     * @param string $path The path of the API endpoint to target, anything after the /1 needs to be specified. (required)
     * @param array $parameters Query parameters to be applied to the current query. (optional)
     * @param array $body The parameters to send with the custom request. (optional)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|object
     */
    public function post(
        $path,
        $parameters = null,
        $body = null,
        $requestOptions = []
    ) {
        // verify the required parameter 'path' is set
        if (!isset($path)) {
            throw new \InvalidArgumentException(
                'Parameter `path` is required when calling `post`.'
            );
        }

        $resourcePath = '/1{path}';
        $queryParameters = [];
        $headers = [];
        $httpBody = isset($body) ? $body : [];

        if ($parameters !== null) {
            $queryParameters = $parameters;
        }

        // path params
        if ($path !== null) {
            $resourcePath = str_replace('{path}', $path, $resourcePath);
        }

        return $this->sendRequest(
            'POST',
            $resourcePath,
            $headers,
            $queryParameters,
            $httpBody,
            $requestOptions
        );
    }

    /**
     * Send requests to the Algolia REST API.
     *
     * @param string $path The path of the API endpoint to target, anything after the /1 needs to be specified. (required)
     * @param array $parameters Query parameters to be applied to the current query. (optional)
     * @param array $body The parameters to send with the custom request. (optional)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|object
     */
    public function put(
        $path,
        $parameters = null,
        $body = null,
        $requestOptions = []
    ) {
        // verify the required parameter 'path' is set
        if (!isset($path)) {
            throw new \InvalidArgumentException(
                'Parameter `path` is required when calling `put`.'
            );
        }

        $resourcePath = '/1{path}';
        $queryParameters = [];
        $headers = [];
        $httpBody = isset($body) ? $body : [];

        if ($parameters !== null) {
            $queryParameters = $parameters;
        }

        // path params
        if ($path !== null) {
            $resourcePath = str_replace('{path}', $path, $resourcePath);
        }

        return $this->sendRequest(
            'PUT',
            $resourcePath,
            $headers,
            $queryParameters,
            $httpBody,
            $requestOptions
        );
    }

    /**
     * Update a configuration.
     *
     * @param string $indexName The index in which to perform the request. (required)
     * @param array $querySuggestionsIndexParam querySuggestionsIndexParam (required)
     * - $querySuggestionsIndexParam['sourceIndices'] => (array) List of source indices used to generate a Query Suggestions index. (required)
     * - $querySuggestionsIndexParam['languages'] => (array) De-duplicate singular and plural suggestions. For example, let's say your index contains English content, and that two suggestions “shoe” and “shoes” end up in your Query Suggestions index. If the English language is configured, only the most popular of those two suggestions would remain.
     * - $querySuggestionsIndexParam['exclude'] => (array) List of words and patterns to exclude from the Query Suggestions index.
     *
     * @see \Algolia\AlgoliaSearch\Model\QuerySuggestions\QuerySuggestionsIndexParam
     *
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\QuerySuggestions\SuccessResponse
     */
    public function updateConfig(
        $indexName,
        $querySuggestionsIndexParam,
        $requestOptions = []
    ) {
        // verify the required parameter 'indexName' is set
        if (!isset($indexName)) {
            throw new \InvalidArgumentException(
                'Parameter `indexName` is required when calling `updateConfig`.'
            );
        }
        // verify the required parameter 'querySuggestionsIndexParam' is set
        if (!isset($querySuggestionsIndexParam)) {
            throw new \InvalidArgumentException(
                'Parameter `querySuggestionsIndexParam` is required when calling `updateConfig`.'
            );
        }

        $resourcePath = '/1/configs/{indexName}';
        $queryParameters = [];
        $headers = [];
        $httpBody = $querySuggestionsIndexParam;

        // path params
        if ($indexName !== null) {
            $resourcePath = str_replace(
                '{indexName}',
                ObjectSerializer::toPathValue($indexName),
                $resourcePath
            );
        }

        return $this->sendRequest(
            'PUT',
            $resourcePath,
            $headers,
            $queryParameters,
            $httpBody,
            $requestOptions
        );
    }

    private function sendRequest(
        $method,
        $resourcePath,
        $headers,
        $queryParameters,
        $httpBody,
        $requestOptions,
        $useReadTransporter = false
    ) {
        if (!isset($requestOptions['headers'])) {
            $requestOptions['headers'] = [];
        }
        if (!isset($requestOptions['queryParameters'])) {
            $requestOptions['queryParameters'] = [];
        }

        $requestOptions['headers'] = array_merge(
            $headers,
            $requestOptions['headers']
        );
        $requestOptions['queryParameters'] = array_merge(
            $queryParameters,
            $requestOptions['queryParameters']
        );
        $query = \GuzzleHttp\Psr7\Query::build(
            $requestOptions['queryParameters']
        );

        return $this->api->sendRequest(
            $method,
            $resourcePath . ($query ? "?{$query}" : ''),
            $httpBody,
            $requestOptions,
            $useReadTransporter
        );
    }
}
