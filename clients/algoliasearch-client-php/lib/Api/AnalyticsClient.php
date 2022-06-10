<?php

namespace Algolia\AlgoliaSearch\Api;

use Algolia\AlgoliaSearch\Algolia;
use Algolia\AlgoliaSearch\Configuration\AnalyticsConfig;
use Algolia\AlgoliaSearch\ObjectSerializer;
use Algolia\AlgoliaSearch\RetryStrategy\ApiWrapper;
use Algolia\AlgoliaSearch\RetryStrategy\ApiWrapperInterface;
use Algolia\AlgoliaSearch\RetryStrategy\ClusterHosts;

/**
 * AnalyticsClient Class Doc Comment
 *
 * @category Class
 * @package  Algolia\AlgoliaSearch
 */
class AnalyticsClient
{
    /**
     * @var ApiWrapperInterface
     */
    protected $api;

    /**
     * @var AnalyticsConfig
     */
    protected $config;

    /**
     * @param AnalyticsConfig $config
     * @param ApiWrapperInterface $apiWrapper
     */
    public function __construct(
        ApiWrapperInterface $apiWrapper,
        AnalyticsConfig $config
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
        $allowedRegions = self::getAllowedRegions();
        $config = AnalyticsConfig::create(
            $appId,
            $apiKey,
            $region,
            $allowedRegions
        );

        return static::createWithConfig($config);
    }

    /**
     * Returns the allowed regions for the config
     */
    public static function getAllowedRegions()
    {
        return ['de', 'us'];
    }

    /**
     * Instantiate the client with configuration
     *
     * @param AnalyticsConfig $config Configuration
     */
    public static function createWithConfig(AnalyticsConfig $config)
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
     * @param AnalyticsConfig $config
     *
     * @return ClusterHosts
     */
    public static function getClusterHosts(AnalyticsConfig $config)
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
                        'analytics.{region}.algolia.com'
                    )
                    : 'analytics.algolia.com';
            $clusterHosts = ClusterHosts::create($url);
        }

        return $clusterHosts;
    }

    /**
     * @return AnalyticsConfig
     */
    public function getClientConfig()
    {
        return $this->config;
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
        if ($path === null) {
            throw new \InvalidArgumentException(
                'Parameter `path` is required when calling `del`.'
            );
        }

        $resourcePath = '/1{path}';
        $queryParameters = [];
        $headers = [];
        $httpBody = [];

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
        if ($path === null) {
            throw new \InvalidArgumentException(
                'Parameter `path` is required when calling `get`.'
            );
        }

        $resourcePath = '/1{path}';
        $queryParameters = [];
        $headers = [];
        $httpBody = [];

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
     * Get average click position.
     *
     * @param string $index The index name to target. (required)
     * @param string $startDate The lower bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param string $endDate The upper bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param string $tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded. (optional)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\Analytics\GetAverageClickPositionResponse
     */
    public function getAverageClickPosition(
        $index,
        $startDate = null,
        $endDate = null,
        $tags = null,
        $requestOptions = []
    ) {
        // verify the required parameter 'index' is set
        if ($index === null) {
            throw new \InvalidArgumentException(
                'Parameter `index` is required when calling `getAverageClickPosition`.'
            );
        }
        if (
            $startDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $startDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "startDate" when calling AnalyticsClient.getAverageClickPosition, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        if (
            $endDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $endDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "endDate" when calling AnalyticsClient.getAverageClickPosition, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        $resourcePath = '/2/clicks/averageClickPosition';
        $queryParameters = [];
        $headers = [];
        $httpBody = [];

        if ($index !== null) {
            $queryParameters['index'] = $index;
        }

        if ($startDate !== null) {
            $queryParameters['startDate'] = $startDate;
        }

        if ($endDate !== null) {
            $queryParameters['endDate'] = $endDate;
        }

        if ($tags !== null) {
            $queryParameters['tags'] = $tags;
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
     * Get clicks per positions.
     *
     * @param string $index The index name to target. (required)
     * @param string $startDate The lower bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param string $endDate The upper bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param string $tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded. (optional)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\Analytics\GetClickPositionsResponse
     */
    public function getClickPositions(
        $index,
        $startDate = null,
        $endDate = null,
        $tags = null,
        $requestOptions = []
    ) {
        // verify the required parameter 'index' is set
        if ($index === null) {
            throw new \InvalidArgumentException(
                'Parameter `index` is required when calling `getClickPositions`.'
            );
        }
        if (
            $startDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $startDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "startDate" when calling AnalyticsClient.getClickPositions, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        if (
            $endDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $endDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "endDate" when calling AnalyticsClient.getClickPositions, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        $resourcePath = '/2/clicks/positions';
        $queryParameters = [];
        $headers = [];
        $httpBody = [];

        if ($index !== null) {
            $queryParameters['index'] = $index;
        }

        if ($startDate !== null) {
            $queryParameters['startDate'] = $startDate;
        }

        if ($endDate !== null) {
            $queryParameters['endDate'] = $endDate;
        }

        if ($tags !== null) {
            $queryParameters['tags'] = $tags;
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
     * Get click-through rate (CTR).
     *
     * @param string $index The index name to target. (required)
     * @param string $startDate The lower bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param string $endDate The upper bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param string $tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded. (optional)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\Analytics\GetClickThroughRateResponse
     */
    public function getClickThroughRate(
        $index,
        $startDate = null,
        $endDate = null,
        $tags = null,
        $requestOptions = []
    ) {
        // verify the required parameter 'index' is set
        if ($index === null) {
            throw new \InvalidArgumentException(
                'Parameter `index` is required when calling `getClickThroughRate`.'
            );
        }
        if (
            $startDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $startDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "startDate" when calling AnalyticsClient.getClickThroughRate, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        if (
            $endDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $endDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "endDate" when calling AnalyticsClient.getClickThroughRate, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        $resourcePath = '/2/clicks/clickThroughRate';
        $queryParameters = [];
        $headers = [];
        $httpBody = [];

        if ($index !== null) {
            $queryParameters['index'] = $index;
        }

        if ($startDate !== null) {
            $queryParameters['startDate'] = $startDate;
        }

        if ($endDate !== null) {
            $queryParameters['endDate'] = $endDate;
        }

        if ($tags !== null) {
            $queryParameters['tags'] = $tags;
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
     * Get conversion rate (CR).
     *
     * @param string $index The index name to target. (required)
     * @param string $startDate The lower bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param string $endDate The upper bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param string $tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded. (optional)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\Analytics\GetConversationRateResponse
     */
    public function getConversationRate(
        $index,
        $startDate = null,
        $endDate = null,
        $tags = null,
        $requestOptions = []
    ) {
        // verify the required parameter 'index' is set
        if ($index === null) {
            throw new \InvalidArgumentException(
                'Parameter `index` is required when calling `getConversationRate`.'
            );
        }
        if (
            $startDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $startDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "startDate" when calling AnalyticsClient.getConversationRate, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        if (
            $endDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $endDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "endDate" when calling AnalyticsClient.getConversationRate, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        $resourcePath = '/2/conversions/conversionRate';
        $queryParameters = [];
        $headers = [];
        $httpBody = [];

        if ($index !== null) {
            $queryParameters['index'] = $index;
        }

        if ($startDate !== null) {
            $queryParameters['startDate'] = $startDate;
        }

        if ($endDate !== null) {
            $queryParameters['endDate'] = $endDate;
        }

        if ($tags !== null) {
            $queryParameters['tags'] = $tags;
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
     * Get no click rate.
     *
     * @param string $index The index name to target. (required)
     * @param string $startDate The lower bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param string $endDate The upper bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param string $tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded. (optional)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\Analytics\GetNoClickRateResponse
     */
    public function getNoClickRate(
        $index,
        $startDate = null,
        $endDate = null,
        $tags = null,
        $requestOptions = []
    ) {
        // verify the required parameter 'index' is set
        if ($index === null) {
            throw new \InvalidArgumentException(
                'Parameter `index` is required when calling `getNoClickRate`.'
            );
        }
        if (
            $startDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $startDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "startDate" when calling AnalyticsClient.getNoClickRate, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        if (
            $endDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $endDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "endDate" when calling AnalyticsClient.getNoClickRate, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        $resourcePath = '/2/searches/noClickRate';
        $queryParameters = [];
        $headers = [];
        $httpBody = [];

        if ($index !== null) {
            $queryParameters['index'] = $index;
        }

        if ($startDate !== null) {
            $queryParameters['startDate'] = $startDate;
        }

        if ($endDate !== null) {
            $queryParameters['endDate'] = $endDate;
        }

        if ($tags !== null) {
            $queryParameters['tags'] = $tags;
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
     * Get no results rate.
     *
     * @param string $index The index name to target. (required)
     * @param string $startDate The lower bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param string $endDate The upper bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param string $tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded. (optional)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\Analytics\GetNoResultsRateResponse
     */
    public function getNoResultsRate(
        $index,
        $startDate = null,
        $endDate = null,
        $tags = null,
        $requestOptions = []
    ) {
        // verify the required parameter 'index' is set
        if ($index === null) {
            throw new \InvalidArgumentException(
                'Parameter `index` is required when calling `getNoResultsRate`.'
            );
        }
        if (
            $startDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $startDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "startDate" when calling AnalyticsClient.getNoResultsRate, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        if (
            $endDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $endDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "endDate" when calling AnalyticsClient.getNoResultsRate, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        $resourcePath = '/2/searches/noResultRate';
        $queryParameters = [];
        $headers = [];
        $httpBody = [];

        if ($index !== null) {
            $queryParameters['index'] = $index;
        }

        if ($startDate !== null) {
            $queryParameters['startDate'] = $startDate;
        }

        if ($endDate !== null) {
            $queryParameters['endDate'] = $endDate;
        }

        if ($tags !== null) {
            $queryParameters['tags'] = $tags;
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
     * Get searches count.
     *
     * @param string $index The index name to target. (required)
     * @param string $startDate The lower bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param string $endDate The upper bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param string $tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded. (optional)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\Analytics\GetSearchesCountResponse
     */
    public function getSearchesCount(
        $index,
        $startDate = null,
        $endDate = null,
        $tags = null,
        $requestOptions = []
    ) {
        // verify the required parameter 'index' is set
        if ($index === null) {
            throw new \InvalidArgumentException(
                'Parameter `index` is required when calling `getSearchesCount`.'
            );
        }
        if (
            $startDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $startDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "startDate" when calling AnalyticsClient.getSearchesCount, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        if (
            $endDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $endDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "endDate" when calling AnalyticsClient.getSearchesCount, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        $resourcePath = '/2/searches/count';
        $queryParameters = [];
        $headers = [];
        $httpBody = [];

        if ($index !== null) {
            $queryParameters['index'] = $index;
        }

        if ($startDate !== null) {
            $queryParameters['startDate'] = $startDate;
        }

        if ($endDate !== null) {
            $queryParameters['endDate'] = $endDate;
        }

        if ($tags !== null) {
            $queryParameters['tags'] = $tags;
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
     * Get top searches with no clicks.
     *
     * @param string $index The index name to target. (required)
     * @param string $startDate The lower bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param string $endDate The upper bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param int $limit Number of records to return. Limit is the size of the page. (optional, default to 10)
     * @param int $offset Position of the starting record. Used for paging. 0 is the first record. (optional, default to 0)
     * @param string $tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded. (optional)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\Analytics\GetSearchesNoClicksResponse
     */
    public function getSearchesNoClicks(
        $index,
        $startDate = null,
        $endDate = null,
        $limit = null,
        $offset = null,
        $tags = null,
        $requestOptions = []
    ) {
        // verify the required parameter 'index' is set
        if ($index === null) {
            throw new \InvalidArgumentException(
                'Parameter `index` is required when calling `getSearchesNoClicks`.'
            );
        }
        if (
            $startDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $startDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "startDate" when calling AnalyticsClient.getSearchesNoClicks, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        if (
            $endDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $endDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "endDate" when calling AnalyticsClient.getSearchesNoClicks, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        $resourcePath = '/2/searches/noClicks';
        $queryParameters = [];
        $headers = [];
        $httpBody = [];

        if ($index !== null) {
            $queryParameters['index'] = $index;
        }

        if ($startDate !== null) {
            $queryParameters['startDate'] = $startDate;
        }

        if ($endDate !== null) {
            $queryParameters['endDate'] = $endDate;
        }

        if ($limit !== null) {
            $queryParameters['limit'] = $limit;
        }

        if ($offset !== null) {
            $queryParameters['offset'] = $offset;
        }

        if ($tags !== null) {
            $queryParameters['tags'] = $tags;
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
     * Get top searches with no results.
     *
     * @param string $index The index name to target. (required)
     * @param string $startDate The lower bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param string $endDate The upper bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param int $limit Number of records to return. Limit is the size of the page. (optional, default to 10)
     * @param int $offset Position of the starting record. Used for paging. 0 is the first record. (optional, default to 0)
     * @param string $tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded. (optional)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\Analytics\GetSearchesNoResultsResponse
     */
    public function getSearchesNoResults(
        $index,
        $startDate = null,
        $endDate = null,
        $limit = null,
        $offset = null,
        $tags = null,
        $requestOptions = []
    ) {
        // verify the required parameter 'index' is set
        if ($index === null) {
            throw new \InvalidArgumentException(
                'Parameter `index` is required when calling `getSearchesNoResults`.'
            );
        }
        if (
            $startDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $startDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "startDate" when calling AnalyticsClient.getSearchesNoResults, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        if (
            $endDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $endDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "endDate" when calling AnalyticsClient.getSearchesNoResults, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        $resourcePath = '/2/searches/noResults';
        $queryParameters = [];
        $headers = [];
        $httpBody = [];

        if ($index !== null) {
            $queryParameters['index'] = $index;
        }

        if ($startDate !== null) {
            $queryParameters['startDate'] = $startDate;
        }

        if ($endDate !== null) {
            $queryParameters['endDate'] = $endDate;
        }

        if ($limit !== null) {
            $queryParameters['limit'] = $limit;
        }

        if ($offset !== null) {
            $queryParameters['offset'] = $offset;
        }

        if ($tags !== null) {
            $queryParameters['tags'] = $tags;
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
     * Get Analytics API status.
     *
     * @param string $index The index name to target. (required)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\Analytics\GetStatusResponse
     */
    public function getStatus($index, $requestOptions = [])
    {
        // verify the required parameter 'index' is set
        if ($index === null) {
            throw new \InvalidArgumentException(
                'Parameter `index` is required when calling `getStatus`.'
            );
        }

        $resourcePath = '/2/status';
        $queryParameters = [];
        $headers = [];
        $httpBody = [];

        if ($index !== null) {
            $queryParameters['index'] = $index;
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
     * Get top countries.
     *
     * @param string $index The index name to target. (required)
     * @param string $startDate The lower bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param string $endDate The upper bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param int $limit Number of records to return. Limit is the size of the page. (optional, default to 10)
     * @param int $offset Position of the starting record. Used for paging. 0 is the first record. (optional, default to 0)
     * @param string $tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded. (optional)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\Analytics\GetTopCountriesResponse
     */
    public function getTopCountries(
        $index,
        $startDate = null,
        $endDate = null,
        $limit = null,
        $offset = null,
        $tags = null,
        $requestOptions = []
    ) {
        // verify the required parameter 'index' is set
        if ($index === null) {
            throw new \InvalidArgumentException(
                'Parameter `index` is required when calling `getTopCountries`.'
            );
        }
        if (
            $startDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $startDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "startDate" when calling AnalyticsClient.getTopCountries, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        if (
            $endDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $endDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "endDate" when calling AnalyticsClient.getTopCountries, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        $resourcePath = '/2/countries';
        $queryParameters = [];
        $headers = [];
        $httpBody = [];

        if ($index !== null) {
            $queryParameters['index'] = $index;
        }

        if ($startDate !== null) {
            $queryParameters['startDate'] = $startDate;
        }

        if ($endDate !== null) {
            $queryParameters['endDate'] = $endDate;
        }

        if ($limit !== null) {
            $queryParameters['limit'] = $limit;
        }

        if ($offset !== null) {
            $queryParameters['offset'] = $offset;
        }

        if ($tags !== null) {
            $queryParameters['tags'] = $tags;
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
     * Get top filter attributes.
     *
     * @param string $index The index name to target. (required)
     * @param string $search The query term to search for. Must match the exact user input. (optional)
     * @param string $startDate The lower bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param string $endDate The upper bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param int $limit Number of records to return. Limit is the size of the page. (optional, default to 10)
     * @param int $offset Position of the starting record. Used for paging. 0 is the first record. (optional, default to 0)
     * @param string $tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded. (optional)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\Analytics\GetTopFilterAttributesResponse
     */
    public function getTopFilterAttributes(
        $index,
        $search = null,
        $startDate = null,
        $endDate = null,
        $limit = null,
        $offset = null,
        $tags = null,
        $requestOptions = []
    ) {
        // verify the required parameter 'index' is set
        if ($index === null) {
            throw new \InvalidArgumentException(
                'Parameter `index` is required when calling `getTopFilterAttributes`.'
            );
        }
        if (
            $startDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $startDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "startDate" when calling AnalyticsClient.getTopFilterAttributes, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        if (
            $endDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $endDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "endDate" when calling AnalyticsClient.getTopFilterAttributes, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        $resourcePath = '/2/filters';
        $queryParameters = [];
        $headers = [];
        $httpBody = [];

        if ($index !== null) {
            $queryParameters['index'] = $index;
        }

        if ($search !== null) {
            $queryParameters['search'] = $search;
        }

        if ($startDate !== null) {
            $queryParameters['startDate'] = $startDate;
        }

        if ($endDate !== null) {
            $queryParameters['endDate'] = $endDate;
        }

        if ($limit !== null) {
            $queryParameters['limit'] = $limit;
        }

        if ($offset !== null) {
            $queryParameters['offset'] = $offset;
        }

        if ($tags !== null) {
            $queryParameters['tags'] = $tags;
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
     * Get top filters for the an attribute.
     *
     * @param string $attribute The exact name of the attribute. (required)
     * @param string $index The index name to target. (required)
     * @param string $search The query term to search for. Must match the exact user input. (optional)
     * @param string $startDate The lower bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param string $endDate The upper bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param int $limit Number of records to return. Limit is the size of the page. (optional, default to 10)
     * @param int $offset Position of the starting record. Used for paging. 0 is the first record. (optional, default to 0)
     * @param string $tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded. (optional)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\Analytics\GetTopFilterForAttributeResponse
     */
    public function getTopFilterForAttribute(
        $attribute,
        $index,
        $search = null,
        $startDate = null,
        $endDate = null,
        $limit = null,
        $offset = null,
        $tags = null,
        $requestOptions = []
    ) {
        // verify the required parameter 'attribute' is set
        if ($attribute === null) {
            throw new \InvalidArgumentException(
                'Parameter `attribute` is required when calling `getTopFilterForAttribute`.'
            );
        }
        // verify the required parameter 'index' is set
        if ($index === null) {
            throw new \InvalidArgumentException(
                'Parameter `index` is required when calling `getTopFilterForAttribute`.'
            );
        }
        if (
            $startDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $startDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "startDate" when calling AnalyticsClient.getTopFilterForAttribute, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        if (
            $endDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $endDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "endDate" when calling AnalyticsClient.getTopFilterForAttribute, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        $resourcePath = '/2/filters/{attribute}';
        $queryParameters = [];
        $headers = [];
        $httpBody = [];

        if ($index !== null) {
            $queryParameters['index'] = $index;
        }

        if ($search !== null) {
            $queryParameters['search'] = $search;
        }

        if ($startDate !== null) {
            $queryParameters['startDate'] = $startDate;
        }

        if ($endDate !== null) {
            $queryParameters['endDate'] = $endDate;
        }

        if ($limit !== null) {
            $queryParameters['limit'] = $limit;
        }

        if ($offset !== null) {
            $queryParameters['offset'] = $offset;
        }

        if ($tags !== null) {
            $queryParameters['tags'] = $tags;
        }

        // path params
        if ($attribute !== null) {
            $resourcePath = str_replace(
                '{attribute}',
                ObjectSerializer::toPathValue($attribute),
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
     * Get top filters for a no result search.
     *
     * @param string $index The index name to target. (required)
     * @param string $search The query term to search for. Must match the exact user input. (optional)
     * @param string $startDate The lower bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param string $endDate The upper bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param int $limit Number of records to return. Limit is the size of the page. (optional, default to 10)
     * @param int $offset Position of the starting record. Used for paging. 0 is the first record. (optional, default to 0)
     * @param string $tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded. (optional)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\Analytics\GetTopFiltersNoResultsResponse
     */
    public function getTopFiltersNoResults(
        $index,
        $search = null,
        $startDate = null,
        $endDate = null,
        $limit = null,
        $offset = null,
        $tags = null,
        $requestOptions = []
    ) {
        // verify the required parameter 'index' is set
        if ($index === null) {
            throw new \InvalidArgumentException(
                'Parameter `index` is required when calling `getTopFiltersNoResults`.'
            );
        }
        if (
            $startDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $startDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "startDate" when calling AnalyticsClient.getTopFiltersNoResults, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        if (
            $endDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $endDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "endDate" when calling AnalyticsClient.getTopFiltersNoResults, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        $resourcePath = '/2/filters/noResults';
        $queryParameters = [];
        $headers = [];
        $httpBody = [];

        if ($index !== null) {
            $queryParameters['index'] = $index;
        }

        if ($search !== null) {
            $queryParameters['search'] = $search;
        }

        if ($startDate !== null) {
            $queryParameters['startDate'] = $startDate;
        }

        if ($endDate !== null) {
            $queryParameters['endDate'] = $endDate;
        }

        if ($limit !== null) {
            $queryParameters['limit'] = $limit;
        }

        if ($offset !== null) {
            $queryParameters['offset'] = $offset;
        }

        if ($tags !== null) {
            $queryParameters['tags'] = $tags;
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
     * Get top hits.
     *
     * @param string $index The index name to target. (required)
     * @param string $search The query term to search for. Must match the exact user input. (optional)
     * @param bool $clickAnalytics Whether to include the click-through and conversion rates for a search. (optional, default to false)
     * @param string $startDate The lower bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param string $endDate The upper bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param int $limit Number of records to return. Limit is the size of the page. (optional, default to 10)
     * @param int $offset Position of the starting record. Used for paging. 0 is the first record. (optional, default to 0)
     * @param string $tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded. (optional)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\Analytics\GetTopHitsResponse
     */
    public function getTopHits(
        $index,
        $search = null,
        $clickAnalytics = null,
        $startDate = null,
        $endDate = null,
        $limit = null,
        $offset = null,
        $tags = null,
        $requestOptions = []
    ) {
        // verify the required parameter 'index' is set
        if ($index === null) {
            throw new \InvalidArgumentException(
                'Parameter `index` is required when calling `getTopHits`.'
            );
        }
        if (
            $startDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $startDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "startDate" when calling AnalyticsClient.getTopHits, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        if (
            $endDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $endDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "endDate" when calling AnalyticsClient.getTopHits, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        $resourcePath = '/2/hits';
        $queryParameters = [];
        $headers = [];
        $httpBody = [];

        if ($index !== null) {
            $queryParameters['index'] = $index;
        }

        if ($search !== null) {
            $queryParameters['search'] = $search;
        }

        if ($clickAnalytics !== null) {
            $queryParameters['clickAnalytics'] = $clickAnalytics;
        }

        if ($startDate !== null) {
            $queryParameters['startDate'] = $startDate;
        }

        if ($endDate !== null) {
            $queryParameters['endDate'] = $endDate;
        }

        if ($limit !== null) {
            $queryParameters['limit'] = $limit;
        }

        if ($offset !== null) {
            $queryParameters['offset'] = $offset;
        }

        if ($tags !== null) {
            $queryParameters['tags'] = $tags;
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
     * Get top searches.
     *
     * @param string $index The index name to target. (required)
     * @param bool $clickAnalytics Whether to include the click-through and conversion rates for a search. (optional, default to false)
     * @param string $startDate The lower bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param string $endDate The upper bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param array $orderBy Reorder the results. (optional)
     * @param array $direction The sorting of the result. (optional)
     * @param int $limit Number of records to return. Limit is the size of the page. (optional, default to 10)
     * @param int $offset Position of the starting record. Used for paging. 0 is the first record. (optional, default to 0)
     * @param string $tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded. (optional)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\Analytics\GetTopSearchesResponse
     */
    public function getTopSearches(
        $index,
        $clickAnalytics = null,
        $startDate = null,
        $endDate = null,
        $orderBy = null,
        $direction = null,
        $limit = null,
        $offset = null,
        $tags = null,
        $requestOptions = []
    ) {
        // verify the required parameter 'index' is set
        if ($index === null) {
            throw new \InvalidArgumentException(
                'Parameter `index` is required when calling `getTopSearches`.'
            );
        }
        if (
            $startDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $startDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "startDate" when calling AnalyticsClient.getTopSearches, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        if (
            $endDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $endDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "endDate" when calling AnalyticsClient.getTopSearches, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        $resourcePath = '/2/searches';
        $queryParameters = [];
        $headers = [];
        $httpBody = [];

        if ($index !== null) {
            $queryParameters['index'] = $index;
        }

        if ($clickAnalytics !== null) {
            $queryParameters['clickAnalytics'] = $clickAnalytics;
        }

        if ($startDate !== null) {
            $queryParameters['startDate'] = $startDate;
        }

        if ($endDate !== null) {
            $queryParameters['endDate'] = $endDate;
        }

        if ($orderBy !== null) {
            $queryParameters['orderBy'] = $orderBy;
        }

        if ($direction !== null) {
            $queryParameters['direction'] = $direction;
        }

        if ($limit !== null) {
            $queryParameters['limit'] = $limit;
        }

        if ($offset !== null) {
            $queryParameters['offset'] = $offset;
        }

        if ($tags !== null) {
            $queryParameters['tags'] = $tags;
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
     * Get users count.
     *
     * @param string $index The index name to target. (required)
     * @param string $startDate The lower bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param string $endDate The upper bound timestamp (a date, a string like \&quot;2006-01-02\&quot;) of the period to analyze. (optional)
     * @param string $tags Filter metrics on the provided tags. Each tag must correspond to an analyticsTags set at search time. Multiple tags can be combined with the operators OR and AND. If a tag contains characters like spaces or parentheses, it should be URL encoded. (optional)
     * @param array $requestOptions the requestOptions to send along with the query, they will be merged with the transporter requestOptions
     *
     * @return array<string, mixed>|\Algolia\AlgoliaSearch\Model\Analytics\GetUsersCountResponse
     */
    public function getUsersCount(
        $index,
        $startDate = null,
        $endDate = null,
        $tags = null,
        $requestOptions = []
    ) {
        // verify the required parameter 'index' is set
        if ($index === null) {
            throw new \InvalidArgumentException(
                'Parameter `index` is required when calling `getUsersCount`.'
            );
        }
        if (
            $startDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $startDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "startDate" when calling AnalyticsClient.getUsersCount, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        if (
            $endDate !== null &&
            !preg_match(
                '/^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/',
                $endDate
            )
        ) {
            throw new \InvalidArgumentException(
                'invalid value for "endDate" when calling AnalyticsClient.getUsersCount, must conform to the pattern /^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/.'
            );
        }

        $resourcePath = '/2/users/count';
        $queryParameters = [];
        $headers = [];
        $httpBody = [];

        if ($index !== null) {
            $queryParameters['index'] = $index;
        }

        if ($startDate !== null) {
            $queryParameters['startDate'] = $startDate;
        }

        if ($endDate !== null) {
            $queryParameters['endDate'] = $endDate;
        }

        if ($tags !== null) {
            $queryParameters['tags'] = $tags;
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
        if ($path === null) {
            throw new \InvalidArgumentException(
                'Parameter `path` is required when calling `post`.'
            );
        }

        $resourcePath = '/1{path}';
        $queryParameters = [];
        $headers = [];
        $httpBody = [];

        if ($parameters !== null) {
            $queryParameters = $parameters;
        }

        // path params
        if ($path !== null) {
            $resourcePath = str_replace('{path}', $path, $resourcePath);
        }

        if (isset($body)) {
            $httpBody = $body;
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
        if ($path === null) {
            throw new \InvalidArgumentException(
                'Parameter `path` is required when calling `put`.'
            );
        }

        $resourcePath = '/1{path}';
        $queryParameters = [];
        $headers = [];
        $httpBody = [];

        if ($parameters !== null) {
            $queryParameters = $parameters;
        }

        // path params
        if ($path !== null) {
            $resourcePath = str_replace('{path}', $path, $resourcePath);
        }

        if (isset($body)) {
            $httpBody = $body;
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
