<?php

namespace Algolia\AlgoliaSearch\Model\Search;

/**
 * IndexSettings Class Doc Comment
 *
 * @category Class
 * @description The Algolia index settings.
 *
 * @package Algolia\AlgoliaSearch
 */
class IndexSettings extends \Algolia\AlgoliaSearch\Model\AbstractModel implements
        ModelInterface,
        \ArrayAccess,
        \JsonSerializable
{
    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @var string[]
     */
    protected static $modelTypes = [
        'replicas' => 'string[]',
        'paginationLimitedTo' => 'int',
        'disableTypoToleranceOnWords' => 'string[]',
        'attributesToTransliterate' => 'string[]',
        'camelCaseAttributes' => 'string[]',
        'decompoundedAttributes' => 'object',
        'indexLanguages' => 'string[]',
        'filterPromotes' => 'bool',
        'disablePrefixOnAttributes' => 'string[]',
        'allowCompressionOfIntegerArray' => 'bool',
        'numericAttributesForFiltering' => 'string[]',
        'userData' => 'object',
        'searchableAttributes' => 'string[]',
        'attributesForFaceting' => 'string[]',
        'unretrievableAttributes' => 'string[]',
        'attributesToRetrieve' => 'string[]',
        'restrictSearchableAttributes' => 'string[]',
        'ranking' => 'string[]',
        'customRanking' => 'string[]',
        'relevancyStrictness' => 'int',
        'attributesToHighlight' => 'string[]',
        'attributesToSnippet' => 'string[]',
        'highlightPreTag' => 'string',
        'highlightPostTag' => 'string',
        'snippetEllipsisText' => 'string',
        'restrictHighlightAndSnippetArrays' => 'bool',
        'hitsPerPage' => 'int',
        'minWordSizefor1Typo' => 'int',
        'minWordSizefor2Typos' => 'int',
        'typoTolerance' => '\Algolia\AlgoliaSearch\Model\Search\TypoTolerance',
        'allowTyposOnNumericTokens' => 'bool',
        'disableTypoToleranceOnAttributes' => 'string[]',
        'separatorsToIndex' => 'string',
        'ignorePlurals' => 'string',
        'removeStopWords' => 'string',
        'keepDiacriticsOnCharacters' => 'string',
        'queryLanguages' => 'string[]',
        'decompoundQuery' => 'bool',
        'enableRules' => 'bool',
        'enablePersonalization' => 'bool',
        'queryType' => '\Algolia\AlgoliaSearch\Model\Search\QueryType',
        'removeWordsIfNoResults' => '\Algolia\AlgoliaSearch\Model\Search\RemoveWordsIfNoResults',
        'advancedSyntax' => 'bool',
        'optionalWords' => 'string[]',
        'disableExactOnAttributes' => 'string[]',
        'exactOnSingleWordQuery' => '\Algolia\AlgoliaSearch\Model\Search\ExactOnSingleWordQuery',
        'alternativesAsExact' => '\Algolia\AlgoliaSearch\Model\Search\AlternativesAsExact[]',
        'advancedSyntaxFeatures' => '\Algolia\AlgoliaSearch\Model\Search\AdvancedSyntaxFeatures[]',
        'distinct' => 'int',
        'synonyms' => 'bool',
        'replaceSynonymsInHighlight' => 'bool',
        'minProximity' => 'int',
        'responseFields' => 'string[]',
        'maxFacetHits' => 'int',
        'attributeCriteriaComputedByMinProximity' => 'bool',
        'renderingContent' => 'object',
    ];

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @var string[]
     */
    protected static $modelFormats = [
        'replicas' => null,
        'paginationLimitedTo' => null,
        'disableTypoToleranceOnWords' => null,
        'attributesToTransliterate' => null,
        'camelCaseAttributes' => null,
        'decompoundedAttributes' => null,
        'indexLanguages' => null,
        'filterPromotes' => null,
        'disablePrefixOnAttributes' => null,
        'allowCompressionOfIntegerArray' => null,
        'numericAttributesForFiltering' => null,
        'userData' => null,
        'searchableAttributes' => null,
        'attributesForFaceting' => null,
        'unretrievableAttributes' => null,
        'attributesToRetrieve' => null,
        'restrictSearchableAttributes' => null,
        'ranking' => null,
        'customRanking' => null,
        'relevancyStrictness' => null,
        'attributesToHighlight' => null,
        'attributesToSnippet' => null,
        'highlightPreTag' => null,
        'highlightPostTag' => null,
        'snippetEllipsisText' => null,
        'restrictHighlightAndSnippetArrays' => null,
        'hitsPerPage' => null,
        'minWordSizefor1Typo' => null,
        'minWordSizefor2Typos' => null,
        'typoTolerance' => null,
        'allowTyposOnNumericTokens' => null,
        'disableTypoToleranceOnAttributes' => null,
        'separatorsToIndex' => null,
        'ignorePlurals' => null,
        'removeStopWords' => null,
        'keepDiacriticsOnCharacters' => null,
        'queryLanguages' => null,
        'decompoundQuery' => null,
        'enableRules' => null,
        'enablePersonalization' => null,
        'queryType' => null,
        'removeWordsIfNoResults' => null,
        'advancedSyntax' => null,
        'optionalWords' => null,
        'disableExactOnAttributes' => null,
        'exactOnSingleWordQuery' => null,
        'alternativesAsExact' => null,
        'advancedSyntaxFeatures' => null,
        'distinct' => null,
        'synonyms' => null,
        'replaceSynonymsInHighlight' => null,
        'minProximity' => null,
        'responseFields' => null,
        'maxFacetHits' => null,
        'attributeCriteriaComputedByMinProximity' => null,
        'renderingContent' => null,
    ];

    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function modelTypes()
    {
        return self::$modelTypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function modelFormats()
    {
        return self::$modelFormats;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'replicas' => 'setReplicas',
        'paginationLimitedTo' => 'setPaginationLimitedTo',
        'disableTypoToleranceOnWords' => 'setDisableTypoToleranceOnWords',
        'attributesToTransliterate' => 'setAttributesToTransliterate',
        'camelCaseAttributes' => 'setCamelCaseAttributes',
        'decompoundedAttributes' => 'setDecompoundedAttributes',
        'indexLanguages' => 'setIndexLanguages',
        'filterPromotes' => 'setFilterPromotes',
        'disablePrefixOnAttributes' => 'setDisablePrefixOnAttributes',
        'allowCompressionOfIntegerArray' => 'setAllowCompressionOfIntegerArray',
        'numericAttributesForFiltering' => 'setNumericAttributesForFiltering',
        'userData' => 'setUserData',
        'searchableAttributes' => 'setSearchableAttributes',
        'attributesForFaceting' => 'setAttributesForFaceting',
        'unretrievableAttributes' => 'setUnretrievableAttributes',
        'attributesToRetrieve' => 'setAttributesToRetrieve',
        'restrictSearchableAttributes' => 'setRestrictSearchableAttributes',
        'ranking' => 'setRanking',
        'customRanking' => 'setCustomRanking',
        'relevancyStrictness' => 'setRelevancyStrictness',
        'attributesToHighlight' => 'setAttributesToHighlight',
        'attributesToSnippet' => 'setAttributesToSnippet',
        'highlightPreTag' => 'setHighlightPreTag',
        'highlightPostTag' => 'setHighlightPostTag',
        'snippetEllipsisText' => 'setSnippetEllipsisText',
        'restrictHighlightAndSnippetArrays' => 'setRestrictHighlightAndSnippetArrays',
        'hitsPerPage' => 'setHitsPerPage',
        'minWordSizefor1Typo' => 'setMinWordSizefor1Typo',
        'minWordSizefor2Typos' => 'setMinWordSizefor2Typos',
        'typoTolerance' => 'setTypoTolerance',
        'allowTyposOnNumericTokens' => 'setAllowTyposOnNumericTokens',
        'disableTypoToleranceOnAttributes' => 'setDisableTypoToleranceOnAttributes',
        'separatorsToIndex' => 'setSeparatorsToIndex',
        'ignorePlurals' => 'setIgnorePlurals',
        'removeStopWords' => 'setRemoveStopWords',
        'keepDiacriticsOnCharacters' => 'setKeepDiacriticsOnCharacters',
        'queryLanguages' => 'setQueryLanguages',
        'decompoundQuery' => 'setDecompoundQuery',
        'enableRules' => 'setEnableRules',
        'enablePersonalization' => 'setEnablePersonalization',
        'queryType' => 'setQueryType',
        'removeWordsIfNoResults' => 'setRemoveWordsIfNoResults',
        'advancedSyntax' => 'setAdvancedSyntax',
        'optionalWords' => 'setOptionalWords',
        'disableExactOnAttributes' => 'setDisableExactOnAttributes',
        'exactOnSingleWordQuery' => 'setExactOnSingleWordQuery',
        'alternativesAsExact' => 'setAlternativesAsExact',
        'advancedSyntaxFeatures' => 'setAdvancedSyntaxFeatures',
        'distinct' => 'setDistinct',
        'synonyms' => 'setSynonyms',
        'replaceSynonymsInHighlight' => 'setReplaceSynonymsInHighlight',
        'minProximity' => 'setMinProximity',
        'responseFields' => 'setResponseFields',
        'maxFacetHits' => 'setMaxFacetHits',
        'attributeCriteriaComputedByMinProximity' => 'setAttributeCriteriaComputedByMinProximity',
        'renderingContent' => 'setRenderingContent',
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'replicas' => 'getReplicas',
        'paginationLimitedTo' => 'getPaginationLimitedTo',
        'disableTypoToleranceOnWords' => 'getDisableTypoToleranceOnWords',
        'attributesToTransliterate' => 'getAttributesToTransliterate',
        'camelCaseAttributes' => 'getCamelCaseAttributes',
        'decompoundedAttributes' => 'getDecompoundedAttributes',
        'indexLanguages' => 'getIndexLanguages',
        'filterPromotes' => 'getFilterPromotes',
        'disablePrefixOnAttributes' => 'getDisablePrefixOnAttributes',
        'allowCompressionOfIntegerArray' => 'getAllowCompressionOfIntegerArray',
        'numericAttributesForFiltering' => 'getNumericAttributesForFiltering',
        'userData' => 'getUserData',
        'searchableAttributes' => 'getSearchableAttributes',
        'attributesForFaceting' => 'getAttributesForFaceting',
        'unretrievableAttributes' => 'getUnretrievableAttributes',
        'attributesToRetrieve' => 'getAttributesToRetrieve',
        'restrictSearchableAttributes' => 'getRestrictSearchableAttributes',
        'ranking' => 'getRanking',
        'customRanking' => 'getCustomRanking',
        'relevancyStrictness' => 'getRelevancyStrictness',
        'attributesToHighlight' => 'getAttributesToHighlight',
        'attributesToSnippet' => 'getAttributesToSnippet',
        'highlightPreTag' => 'getHighlightPreTag',
        'highlightPostTag' => 'getHighlightPostTag',
        'snippetEllipsisText' => 'getSnippetEllipsisText',
        'restrictHighlightAndSnippetArrays' => 'getRestrictHighlightAndSnippetArrays',
        'hitsPerPage' => 'getHitsPerPage',
        'minWordSizefor1Typo' => 'getMinWordSizefor1Typo',
        'minWordSizefor2Typos' => 'getMinWordSizefor2Typos',
        'typoTolerance' => 'getTypoTolerance',
        'allowTyposOnNumericTokens' => 'getAllowTyposOnNumericTokens',
        'disableTypoToleranceOnAttributes' => 'getDisableTypoToleranceOnAttributes',
        'separatorsToIndex' => 'getSeparatorsToIndex',
        'ignorePlurals' => 'getIgnorePlurals',
        'removeStopWords' => 'getRemoveStopWords',
        'keepDiacriticsOnCharacters' => 'getKeepDiacriticsOnCharacters',
        'queryLanguages' => 'getQueryLanguages',
        'decompoundQuery' => 'getDecompoundQuery',
        'enableRules' => 'getEnableRules',
        'enablePersonalization' => 'getEnablePersonalization',
        'queryType' => 'getQueryType',
        'removeWordsIfNoResults' => 'getRemoveWordsIfNoResults',
        'advancedSyntax' => 'getAdvancedSyntax',
        'optionalWords' => 'getOptionalWords',
        'disableExactOnAttributes' => 'getDisableExactOnAttributes',
        'exactOnSingleWordQuery' => 'getExactOnSingleWordQuery',
        'alternativesAsExact' => 'getAlternativesAsExact',
        'advancedSyntaxFeatures' => 'getAdvancedSyntaxFeatures',
        'distinct' => 'getDistinct',
        'synonyms' => 'getSynonyms',
        'replaceSynonymsInHighlight' => 'getReplaceSynonymsInHighlight',
        'minProximity' => 'getMinProximity',
        'responseFields' => 'getResponseFields',
        'maxFacetHits' => 'getMaxFacetHits',
        'attributeCriteriaComputedByMinProximity' => 'getAttributeCriteriaComputedByMinProximity',
        'renderingContent' => 'getRenderingContent',
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * Associative array for storing property values
     *
     * @var mixed[]
     */
    protected $container = [];

    /**
     * Constructor
     *
     * @param mixed[] $data Associated array of property values
     */
    public function __construct(array $data = null)
    {
        if (isset($data['replicas'])) {
            $this->container['replicas'] = $data['replicas'];
        }
        if (isset($data['paginationLimitedTo'])) {
            $this->container['paginationLimitedTo'] =
                $data['paginationLimitedTo'];
        }
        if (isset($data['disableTypoToleranceOnWords'])) {
            $this->container['disableTypoToleranceOnWords'] =
                $data['disableTypoToleranceOnWords'];
        }
        if (isset($data['attributesToTransliterate'])) {
            $this->container['attributesToTransliterate'] =
                $data['attributesToTransliterate'];
        }
        if (isset($data['camelCaseAttributes'])) {
            $this->container['camelCaseAttributes'] =
                $data['camelCaseAttributes'];
        }
        if (isset($data['decompoundedAttributes'])) {
            $this->container['decompoundedAttributes'] =
                $data['decompoundedAttributes'];
        }
        if (isset($data['indexLanguages'])) {
            $this->container['indexLanguages'] = $data['indexLanguages'];
        }
        if (isset($data['filterPromotes'])) {
            $this->container['filterPromotes'] = $data['filterPromotes'];
        }
        if (isset($data['disablePrefixOnAttributes'])) {
            $this->container['disablePrefixOnAttributes'] =
                $data['disablePrefixOnAttributes'];
        }
        if (isset($data['allowCompressionOfIntegerArray'])) {
            $this->container['allowCompressionOfIntegerArray'] =
                $data['allowCompressionOfIntegerArray'];
        }
        if (isset($data['numericAttributesForFiltering'])) {
            $this->container['numericAttributesForFiltering'] =
                $data['numericAttributesForFiltering'];
        }
        if (isset($data['userData'])) {
            $this->container['userData'] = $data['userData'];
        }
        if (isset($data['searchableAttributes'])) {
            $this->container['searchableAttributes'] =
                $data['searchableAttributes'];
        }
        if (isset($data['attributesForFaceting'])) {
            $this->container['attributesForFaceting'] =
                $data['attributesForFaceting'];
        }
        if (isset($data['unretrievableAttributes'])) {
            $this->container['unretrievableAttributes'] =
                $data['unretrievableAttributes'];
        }
        if (isset($data['attributesToRetrieve'])) {
            $this->container['attributesToRetrieve'] =
                $data['attributesToRetrieve'];
        }
        if (isset($data['restrictSearchableAttributes'])) {
            $this->container['restrictSearchableAttributes'] =
                $data['restrictSearchableAttributes'];
        }
        if (isset($data['ranking'])) {
            $this->container['ranking'] = $data['ranking'];
        }
        if (isset($data['customRanking'])) {
            $this->container['customRanking'] = $data['customRanking'];
        }
        if (isset($data['relevancyStrictness'])) {
            $this->container['relevancyStrictness'] =
                $data['relevancyStrictness'];
        }
        if (isset($data['attributesToHighlight'])) {
            $this->container['attributesToHighlight'] =
                $data['attributesToHighlight'];
        }
        if (isset($data['attributesToSnippet'])) {
            $this->container['attributesToSnippet'] =
                $data['attributesToSnippet'];
        }
        if (isset($data['highlightPreTag'])) {
            $this->container['highlightPreTag'] = $data['highlightPreTag'];
        }
        if (isset($data['highlightPostTag'])) {
            $this->container['highlightPostTag'] = $data['highlightPostTag'];
        }
        if (isset($data['snippetEllipsisText'])) {
            $this->container['snippetEllipsisText'] =
                $data['snippetEllipsisText'];
        }
        if (isset($data['restrictHighlightAndSnippetArrays'])) {
            $this->container['restrictHighlightAndSnippetArrays'] =
                $data['restrictHighlightAndSnippetArrays'];
        }
        if (isset($data['hitsPerPage'])) {
            $this->container['hitsPerPage'] = $data['hitsPerPage'];
        }
        if (isset($data['minWordSizefor1Typo'])) {
            $this->container['minWordSizefor1Typo'] =
                $data['minWordSizefor1Typo'];
        }
        if (isset($data['minWordSizefor2Typos'])) {
            $this->container['minWordSizefor2Typos'] =
                $data['minWordSizefor2Typos'];
        }
        if (isset($data['typoTolerance'])) {
            $this->container['typoTolerance'] = $data['typoTolerance'];
        }
        if (isset($data['allowTyposOnNumericTokens'])) {
            $this->container['allowTyposOnNumericTokens'] =
                $data['allowTyposOnNumericTokens'];
        }
        if (isset($data['disableTypoToleranceOnAttributes'])) {
            $this->container['disableTypoToleranceOnAttributes'] =
                $data['disableTypoToleranceOnAttributes'];
        }
        if (isset($data['separatorsToIndex'])) {
            $this->container['separatorsToIndex'] = $data['separatorsToIndex'];
        }
        if (isset($data['ignorePlurals'])) {
            $this->container['ignorePlurals'] = $data['ignorePlurals'];
        }
        if (isset($data['removeStopWords'])) {
            $this->container['removeStopWords'] = $data['removeStopWords'];
        }
        if (isset($data['keepDiacriticsOnCharacters'])) {
            $this->container['keepDiacriticsOnCharacters'] =
                $data['keepDiacriticsOnCharacters'];
        }
        if (isset($data['queryLanguages'])) {
            $this->container['queryLanguages'] = $data['queryLanguages'];
        }
        if (isset($data['decompoundQuery'])) {
            $this->container['decompoundQuery'] = $data['decompoundQuery'];
        }
        if (isset($data['enableRules'])) {
            $this->container['enableRules'] = $data['enableRules'];
        }
        if (isset($data['enablePersonalization'])) {
            $this->container['enablePersonalization'] =
                $data['enablePersonalization'];
        }
        if (isset($data['queryType'])) {
            $this->container['queryType'] = $data['queryType'];
        }
        if (isset($data['removeWordsIfNoResults'])) {
            $this->container['removeWordsIfNoResults'] =
                $data['removeWordsIfNoResults'];
        }
        if (isset($data['advancedSyntax'])) {
            $this->container['advancedSyntax'] = $data['advancedSyntax'];
        }
        if (isset($data['optionalWords'])) {
            $this->container['optionalWords'] = $data['optionalWords'];
        }
        if (isset($data['disableExactOnAttributes'])) {
            $this->container['disableExactOnAttributes'] =
                $data['disableExactOnAttributes'];
        }
        if (isset($data['exactOnSingleWordQuery'])) {
            $this->container['exactOnSingleWordQuery'] =
                $data['exactOnSingleWordQuery'];
        }
        if (isset($data['alternativesAsExact'])) {
            $this->container['alternativesAsExact'] =
                $data['alternativesAsExact'];
        }
        if (isset($data['advancedSyntaxFeatures'])) {
            $this->container['advancedSyntaxFeatures'] =
                $data['advancedSyntaxFeatures'];
        }
        if (isset($data['distinct'])) {
            $this->container['distinct'] = $data['distinct'];
        }
        if (isset($data['synonyms'])) {
            $this->container['synonyms'] = $data['synonyms'];
        }
        if (isset($data['replaceSynonymsInHighlight'])) {
            $this->container['replaceSynonymsInHighlight'] =
                $data['replaceSynonymsInHighlight'];
        }
        if (isset($data['minProximity'])) {
            $this->container['minProximity'] = $data['minProximity'];
        }
        if (isset($data['responseFields'])) {
            $this->container['responseFields'] = $data['responseFields'];
        }
        if (isset($data['maxFacetHits'])) {
            $this->container['maxFacetHits'] = $data['maxFacetHits'];
        }
        if (isset($data['attributeCriteriaComputedByMinProximity'])) {
            $this->container['attributeCriteriaComputedByMinProximity'] =
                $data['attributeCriteriaComputedByMinProximity'];
        }
        if (isset($data['renderingContent'])) {
            $this->container['renderingContent'] = $data['renderingContent'];
        }
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        if (
            isset($this->container['distinct']) &&
            $this->container['distinct'] > 4
        ) {
            $invalidProperties[] =
                "invalid value for 'distinct', must be smaller than or equal to 4.";
        }

        if (
            isset($this->container['distinct']) &&
            $this->container['distinct'] < 0
        ) {
            $invalidProperties[] =
                "invalid value for 'distinct', must be bigger than or equal to 0.";
        }

        if (
            isset($this->container['minProximity']) &&
            $this->container['minProximity'] > 7
        ) {
            $invalidProperties[] =
                "invalid value for 'minProximity', must be smaller than or equal to 7.";
        }

        if (
            isset($this->container['minProximity']) &&
            $this->container['minProximity'] < 1
        ) {
            $invalidProperties[] =
                "invalid value for 'minProximity', must be bigger than or equal to 1.";
        }

        if (
            isset($this->container['maxFacetHits']) &&
            $this->container['maxFacetHits'] > 100
        ) {
            $invalidProperties[] =
                "invalid value for 'maxFacetHits', must be smaller than or equal to 100.";
        }

        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return count($this->listInvalidProperties()) === 0;
    }

    /**
     * Gets replicas
     *
     * @return string[]|null
     */
    public function getReplicas()
    {
        return $this->container['replicas'] ?? null;
    }

    /**
     * Sets replicas
     *
     * @param string[]|null $replicas creates replicas, exact copies of an index
     *
     * @return self
     */
    public function setReplicas($replicas)
    {
        $this->container['replicas'] = $replicas;

        return $this;
    }

    /**
     * Gets paginationLimitedTo
     *
     * @return int|null
     */
    public function getPaginationLimitedTo()
    {
        return $this->container['paginationLimitedTo'] ?? null;
    }

    /**
     * Sets paginationLimitedTo
     *
     * @param int|null $paginationLimitedTo set the maximum number of hits accessible via pagination
     *
     * @return self
     */
    public function setPaginationLimitedTo($paginationLimitedTo)
    {
        $this->container['paginationLimitedTo'] = $paginationLimitedTo;

        return $this;
    }

    /**
     * Gets disableTypoToleranceOnWords
     *
     * @return string[]|null
     */
    public function getDisableTypoToleranceOnWords()
    {
        return $this->container['disableTypoToleranceOnWords'] ?? null;
    }

    /**
     * Sets disableTypoToleranceOnWords
     *
     * @param string[]|null $disableTypoToleranceOnWords a list of words for which you want to turn off typo tolerance
     *
     * @return self
     */
    public function setDisableTypoToleranceOnWords($disableTypoToleranceOnWords)
    {
        $this->container[
            'disableTypoToleranceOnWords'
        ] = $disableTypoToleranceOnWords;

        return $this;
    }

    /**
     * Gets attributesToTransliterate
     *
     * @return string[]|null
     */
    public function getAttributesToTransliterate()
    {
        return $this->container['attributesToTransliterate'] ?? null;
    }

    /**
     * Sets attributesToTransliterate
     *
     * @param string[]|null $attributesToTransliterate specify on which attributes to apply transliteration
     *
     * @return self
     */
    public function setAttributesToTransliterate($attributesToTransliterate)
    {
        $this->container[
            'attributesToTransliterate'
        ] = $attributesToTransliterate;

        return $this;
    }

    /**
     * Gets camelCaseAttributes
     *
     * @return string[]|null
     */
    public function getCamelCaseAttributes()
    {
        return $this->container['camelCaseAttributes'] ?? null;
    }

    /**
     * Sets camelCaseAttributes
     *
     * @param string[]|null $camelCaseAttributes list of attributes on which to do a decomposition of camel case words
     *
     * @return self
     */
    public function setCamelCaseAttributes($camelCaseAttributes)
    {
        $this->container['camelCaseAttributes'] = $camelCaseAttributes;

        return $this;
    }

    /**
     * Gets decompoundedAttributes
     *
     * @return object|null
     */
    public function getDecompoundedAttributes()
    {
        return $this->container['decompoundedAttributes'] ?? null;
    }

    /**
     * Sets decompoundedAttributes
     *
     * @param object|null $decompoundedAttributes specify on which attributes in your index Algolia should apply word segmentation, also known as decompounding
     *
     * @return self
     */
    public function setDecompoundedAttributes($decompoundedAttributes)
    {
        $this->container['decompoundedAttributes'] = $decompoundedAttributes;

        return $this;
    }

    /**
     * Gets indexLanguages
     *
     * @return string[]|null
     */
    public function getIndexLanguages()
    {
        return $this->container['indexLanguages'] ?? null;
    }

    /**
     * Sets indexLanguages
     *
     * @param string[]|null $indexLanguages sets the languages at the index level for language-specific processing such as tokenization and normalization
     *
     * @return self
     */
    public function setIndexLanguages($indexLanguages)
    {
        $this->container['indexLanguages'] = $indexLanguages;

        return $this;
    }

    /**
     * Gets filterPromotes
     *
     * @return bool|null
     */
    public function getFilterPromotes()
    {
        return $this->container['filterPromotes'] ?? null;
    }

    /**
     * Sets filterPromotes
     *
     * @param bool|null $filterPromotes whether promoted results should match the filters of the current search, except for geographic filters
     *
     * @return self
     */
    public function setFilterPromotes($filterPromotes)
    {
        $this->container['filterPromotes'] = $filterPromotes;

        return $this;
    }

    /**
     * Gets disablePrefixOnAttributes
     *
     * @return string[]|null
     */
    public function getDisablePrefixOnAttributes()
    {
        return $this->container['disablePrefixOnAttributes'] ?? null;
    }

    /**
     * Sets disablePrefixOnAttributes
     *
     * @param string[]|null $disablePrefixOnAttributes list of attributes on which you want to disable prefix matching
     *
     * @return self
     */
    public function setDisablePrefixOnAttributes($disablePrefixOnAttributes)
    {
        $this->container[
            'disablePrefixOnAttributes'
        ] = $disablePrefixOnAttributes;

        return $this;
    }

    /**
     * Gets allowCompressionOfIntegerArray
     *
     * @return bool|null
     */
    public function getAllowCompressionOfIntegerArray()
    {
        return $this->container['allowCompressionOfIntegerArray'] ?? null;
    }

    /**
     * Sets allowCompressionOfIntegerArray
     *
     * @param bool|null $allowCompressionOfIntegerArray enables compression of large integer arrays
     *
     * @return self
     */
    public function setAllowCompressionOfIntegerArray(
        $allowCompressionOfIntegerArray
    ) {
        $this->container[
            'allowCompressionOfIntegerArray'
        ] = $allowCompressionOfIntegerArray;

        return $this;
    }

    /**
     * Gets numericAttributesForFiltering
     *
     * @return string[]|null
     */
    public function getNumericAttributesForFiltering()
    {
        return $this->container['numericAttributesForFiltering'] ?? null;
    }

    /**
     * Sets numericAttributesForFiltering
     *
     * @param string[]|null $numericAttributesForFiltering list of numeric attributes that can be used as numerical filters
     *
     * @return self
     */
    public function setNumericAttributesForFiltering(
        $numericAttributesForFiltering
    ) {
        $this->container[
            'numericAttributesForFiltering'
        ] = $numericAttributesForFiltering;

        return $this;
    }

    /**
     * Gets userData
     *
     * @return object|null
     */
    public function getUserData()
    {
        return $this->container['userData'] ?? null;
    }

    /**
     * Sets userData
     *
     * @param object|null $userData lets you store custom data in your indices
     *
     * @return self
     */
    public function setUserData($userData)
    {
        $this->container['userData'] = $userData;

        return $this;
    }

    /**
     * Gets searchableAttributes
     *
     * @return string[]|null
     */
    public function getSearchableAttributes()
    {
        return $this->container['searchableAttributes'] ?? null;
    }

    /**
     * Sets searchableAttributes
     *
     * @param string[]|null $searchableAttributes the complete list of attributes used for searching
     *
     * @return self
     */
    public function setSearchableAttributes($searchableAttributes)
    {
        $this->container['searchableAttributes'] = $searchableAttributes;

        return $this;
    }

    /**
     * Gets attributesForFaceting
     *
     * @return string[]|null
     */
    public function getAttributesForFaceting()
    {
        return $this->container['attributesForFaceting'] ?? null;
    }

    /**
     * Sets attributesForFaceting
     *
     * @param string[]|null $attributesForFaceting the complete list of attributes that will be used for faceting
     *
     * @return self
     */
    public function setAttributesForFaceting($attributesForFaceting)
    {
        $this->container['attributesForFaceting'] = $attributesForFaceting;

        return $this;
    }

    /**
     * Gets unretrievableAttributes
     *
     * @return string[]|null
     */
    public function getUnretrievableAttributes()
    {
        return $this->container['unretrievableAttributes'] ?? null;
    }

    /**
     * Sets unretrievableAttributes
     *
     * @param string[]|null $unretrievableAttributes list of attributes that can't be retrieved at query time
     *
     * @return self
     */
    public function setUnretrievableAttributes($unretrievableAttributes)
    {
        $this->container['unretrievableAttributes'] = $unretrievableAttributes;

        return $this;
    }

    /**
     * Gets attributesToRetrieve
     *
     * @return string[]|null
     */
    public function getAttributesToRetrieve()
    {
        return $this->container['attributesToRetrieve'] ?? null;
    }

    /**
     * Sets attributesToRetrieve
     *
     * @param string[]|null $attributesToRetrieve this parameter controls which attributes to retrieve and which not to retrieve
     *
     * @return self
     */
    public function setAttributesToRetrieve($attributesToRetrieve)
    {
        $this->container['attributesToRetrieve'] = $attributesToRetrieve;

        return $this;
    }

    /**
     * Gets restrictSearchableAttributes
     *
     * @return string[]|null
     */
    public function getRestrictSearchableAttributes()
    {
        return $this->container['restrictSearchableAttributes'] ?? null;
    }

    /**
     * Sets restrictSearchableAttributes
     *
     * @param string[]|null $restrictSearchableAttributes restricts a given query to look in only a subset of your searchable attributes
     *
     * @return self
     */
    public function setRestrictSearchableAttributes(
        $restrictSearchableAttributes
    ) {
        $this->container[
            'restrictSearchableAttributes'
        ] = $restrictSearchableAttributes;

        return $this;
    }

    /**
     * Gets ranking
     *
     * @return string[]|null
     */
    public function getRanking()
    {
        return $this->container['ranking'] ?? null;
    }

    /**
     * Sets ranking
     *
     * @param string[]|null $ranking controls how Algolia should sort your results
     *
     * @return self
     */
    public function setRanking($ranking)
    {
        $this->container['ranking'] = $ranking;

        return $this;
    }

    /**
     * Gets customRanking
     *
     * @return string[]|null
     */
    public function getCustomRanking()
    {
        return $this->container['customRanking'] ?? null;
    }

    /**
     * Sets customRanking
     *
     * @param string[]|null $customRanking specifies the custom ranking criterion
     *
     * @return self
     */
    public function setCustomRanking($customRanking)
    {
        $this->container['customRanking'] = $customRanking;

        return $this;
    }

    /**
     * Gets relevancyStrictness
     *
     * @return int|null
     */
    public function getRelevancyStrictness()
    {
        return $this->container['relevancyStrictness'] ?? null;
    }

    /**
     * Sets relevancyStrictness
     *
     * @param int|null $relevancyStrictness controls the relevancy threshold below which less relevant results aren't included in the results
     *
     * @return self
     */
    public function setRelevancyStrictness($relevancyStrictness)
    {
        $this->container['relevancyStrictness'] = $relevancyStrictness;

        return $this;
    }

    /**
     * Gets attributesToHighlight
     *
     * @return string[]|null
     */
    public function getAttributesToHighlight()
    {
        return $this->container['attributesToHighlight'] ?? null;
    }

    /**
     * Sets attributesToHighlight
     *
     * @param string[]|null $attributesToHighlight list of attributes to highlight
     *
     * @return self
     */
    public function setAttributesToHighlight($attributesToHighlight)
    {
        $this->container['attributesToHighlight'] = $attributesToHighlight;

        return $this;
    }

    /**
     * Gets attributesToSnippet
     *
     * @return string[]|null
     */
    public function getAttributesToSnippet()
    {
        return $this->container['attributesToSnippet'] ?? null;
    }

    /**
     * Sets attributesToSnippet
     *
     * @param string[]|null $attributesToSnippet list of attributes to snippet, with an optional maximum number of words to snippet
     *
     * @return self
     */
    public function setAttributesToSnippet($attributesToSnippet)
    {
        $this->container['attributesToSnippet'] = $attributesToSnippet;

        return $this;
    }

    /**
     * Gets highlightPreTag
     *
     * @return string|null
     */
    public function getHighlightPreTag()
    {
        return $this->container['highlightPreTag'] ?? null;
    }

    /**
     * Sets highlightPreTag
     *
     * @param string|null $highlightPreTag the HTML string to insert before the highlighted parts in all highlight and snippet results
     *
     * @return self
     */
    public function setHighlightPreTag($highlightPreTag)
    {
        $this->container['highlightPreTag'] = $highlightPreTag;

        return $this;
    }

    /**
     * Gets highlightPostTag
     *
     * @return string|null
     */
    public function getHighlightPostTag()
    {
        return $this->container['highlightPostTag'] ?? null;
    }

    /**
     * Sets highlightPostTag
     *
     * @param string|null $highlightPostTag the HTML string to insert after the highlighted parts in all highlight and snippet results
     *
     * @return self
     */
    public function setHighlightPostTag($highlightPostTag)
    {
        $this->container['highlightPostTag'] = $highlightPostTag;

        return $this;
    }

    /**
     * Gets snippetEllipsisText
     *
     * @return string|null
     */
    public function getSnippetEllipsisText()
    {
        return $this->container['snippetEllipsisText'] ?? null;
    }

    /**
     * Sets snippetEllipsisText
     *
     * @param string|null $snippetEllipsisText string used as an ellipsis indicator when a snippet is truncated
     *
     * @return self
     */
    public function setSnippetEllipsisText($snippetEllipsisText)
    {
        $this->container['snippetEllipsisText'] = $snippetEllipsisText;

        return $this;
    }

    /**
     * Gets restrictHighlightAndSnippetArrays
     *
     * @return bool|null
     */
    public function getRestrictHighlightAndSnippetArrays()
    {
        return $this->container['restrictHighlightAndSnippetArrays'] ?? null;
    }

    /**
     * Sets restrictHighlightAndSnippetArrays
     *
     * @param bool|null $restrictHighlightAndSnippetArrays restrict highlighting and snippeting to items that matched the query
     *
     * @return self
     */
    public function setRestrictHighlightAndSnippetArrays(
        $restrictHighlightAndSnippetArrays
    ) {
        $this->container[
            'restrictHighlightAndSnippetArrays'
        ] = $restrictHighlightAndSnippetArrays;

        return $this;
    }

    /**
     * Gets hitsPerPage
     *
     * @return int|null
     */
    public function getHitsPerPage()
    {
        return $this->container['hitsPerPage'] ?? null;
    }

    /**
     * Sets hitsPerPage
     *
     * @param int|null $hitsPerPage set the number of hits per page
     *
     * @return self
     */
    public function setHitsPerPage($hitsPerPage)
    {
        $this->container['hitsPerPage'] = $hitsPerPage;

        return $this;
    }

    /**
     * Gets minWordSizefor1Typo
     *
     * @return int|null
     */
    public function getMinWordSizefor1Typo()
    {
        return $this->container['minWordSizefor1Typo'] ?? null;
    }

    /**
     * Sets minWordSizefor1Typo
     *
     * @param int|null $minWordSizefor1Typo minimum number of characters a word in the query string must contain to accept matches with 1 typo
     *
     * @return self
     */
    public function setMinWordSizefor1Typo($minWordSizefor1Typo)
    {
        $this->container['minWordSizefor1Typo'] = $minWordSizefor1Typo;

        return $this;
    }

    /**
     * Gets minWordSizefor2Typos
     *
     * @return int|null
     */
    public function getMinWordSizefor2Typos()
    {
        return $this->container['minWordSizefor2Typos'] ?? null;
    }

    /**
     * Sets minWordSizefor2Typos
     *
     * @param int|null $minWordSizefor2Typos minimum number of characters a word in the query string must contain to accept matches with 2 typos
     *
     * @return self
     */
    public function setMinWordSizefor2Typos($minWordSizefor2Typos)
    {
        $this->container['minWordSizefor2Typos'] = $minWordSizefor2Typos;

        return $this;
    }

    /**
     * Gets typoTolerance
     *
     * @return \Algolia\AlgoliaSearch\Model\Search\TypoTolerance|null
     */
    public function getTypoTolerance()
    {
        return $this->container['typoTolerance'] ?? null;
    }

    /**
     * Sets typoTolerance
     *
     * @param \Algolia\AlgoliaSearch\Model\Search\TypoTolerance|null $typoTolerance typoTolerance
     *
     * @return self
     */
    public function setTypoTolerance($typoTolerance)
    {
        $this->container['typoTolerance'] = $typoTolerance;

        return $this;
    }

    /**
     * Gets allowTyposOnNumericTokens
     *
     * @return bool|null
     */
    public function getAllowTyposOnNumericTokens()
    {
        return $this->container['allowTyposOnNumericTokens'] ?? null;
    }

    /**
     * Sets allowTyposOnNumericTokens
     *
     * @param bool|null $allowTyposOnNumericTokens whether to allow typos on numbers (\"numeric tokens\") in the query string
     *
     * @return self
     */
    public function setAllowTyposOnNumericTokens($allowTyposOnNumericTokens)
    {
        $this->container[
            'allowTyposOnNumericTokens'
        ] = $allowTyposOnNumericTokens;

        return $this;
    }

    /**
     * Gets disableTypoToleranceOnAttributes
     *
     * @return string[]|null
     */
    public function getDisableTypoToleranceOnAttributes()
    {
        return $this->container['disableTypoToleranceOnAttributes'] ?? null;
    }

    /**
     * Sets disableTypoToleranceOnAttributes
     *
     * @param string[]|null $disableTypoToleranceOnAttributes list of attributes on which you want to disable typo tolerance
     *
     * @return self
     */
    public function setDisableTypoToleranceOnAttributes(
        $disableTypoToleranceOnAttributes
    ) {
        $this->container[
            'disableTypoToleranceOnAttributes'
        ] = $disableTypoToleranceOnAttributes;

        return $this;
    }

    /**
     * Gets separatorsToIndex
     *
     * @return string|null
     */
    public function getSeparatorsToIndex()
    {
        return $this->container['separatorsToIndex'] ?? null;
    }

    /**
     * Sets separatorsToIndex
     *
     * @param string|null $separatorsToIndex control which separators are indexed
     *
     * @return self
     */
    public function setSeparatorsToIndex($separatorsToIndex)
    {
        $this->container['separatorsToIndex'] = $separatorsToIndex;

        return $this;
    }

    /**
     * Gets ignorePlurals
     *
     * @return string|null
     */
    public function getIgnorePlurals()
    {
        return $this->container['ignorePlurals'] ?? null;
    }

    /**
     * Sets ignorePlurals
     *
     * @param string|null $ignorePlurals treats singular, plurals, and other forms of declensions as matching terms
     *
     * @return self
     */
    public function setIgnorePlurals($ignorePlurals)
    {
        $this->container['ignorePlurals'] = $ignorePlurals;

        return $this;
    }

    /**
     * Gets removeStopWords
     *
     * @return string|null
     */
    public function getRemoveStopWords()
    {
        return $this->container['removeStopWords'] ?? null;
    }

    /**
     * Sets removeStopWords
     *
     * @param string|null $removeStopWords removes stop (common) words from the query before executing it
     *
     * @return self
     */
    public function setRemoveStopWords($removeStopWords)
    {
        $this->container['removeStopWords'] = $removeStopWords;

        return $this;
    }

    /**
     * Gets keepDiacriticsOnCharacters
     *
     * @return string|null
     */
    public function getKeepDiacriticsOnCharacters()
    {
        return $this->container['keepDiacriticsOnCharacters'] ?? null;
    }

    /**
     * Sets keepDiacriticsOnCharacters
     *
     * @param string|null $keepDiacriticsOnCharacters list of characters that the engine shouldn't automatically normalize
     *
     * @return self
     */
    public function setKeepDiacriticsOnCharacters($keepDiacriticsOnCharacters)
    {
        $this->container[
            'keepDiacriticsOnCharacters'
        ] = $keepDiacriticsOnCharacters;

        return $this;
    }

    /**
     * Gets queryLanguages
     *
     * @return string[]|null
     */
    public function getQueryLanguages()
    {
        return $this->container['queryLanguages'] ?? null;
    }

    /**
     * Sets queryLanguages
     *
     * @param string[]|null $queryLanguages sets the languages to be used by language-specific settings and functionalities such as ignorePlurals, removeStopWords, and CJK word-detection
     *
     * @return self
     */
    public function setQueryLanguages($queryLanguages)
    {
        $this->container['queryLanguages'] = $queryLanguages;

        return $this;
    }

    /**
     * Gets decompoundQuery
     *
     * @return bool|null
     */
    public function getDecompoundQuery()
    {
        return $this->container['decompoundQuery'] ?? null;
    }

    /**
     * Sets decompoundQuery
     *
     * @param bool|null $decompoundQuery splits compound words into their composing atoms in the query
     *
     * @return self
     */
    public function setDecompoundQuery($decompoundQuery)
    {
        $this->container['decompoundQuery'] = $decompoundQuery;

        return $this;
    }

    /**
     * Gets enableRules
     *
     * @return bool|null
     */
    public function getEnableRules()
    {
        return $this->container['enableRules'] ?? null;
    }

    /**
     * Sets enableRules
     *
     * @param bool|null $enableRules whether Rules should be globally enabled
     *
     * @return self
     */
    public function setEnableRules($enableRules)
    {
        $this->container['enableRules'] = $enableRules;

        return $this;
    }

    /**
     * Gets enablePersonalization
     *
     * @return bool|null
     */
    public function getEnablePersonalization()
    {
        return $this->container['enablePersonalization'] ?? null;
    }

    /**
     * Sets enablePersonalization
     *
     * @param bool|null $enablePersonalization enable the Personalization feature
     *
     * @return self
     */
    public function setEnablePersonalization($enablePersonalization)
    {
        $this->container['enablePersonalization'] = $enablePersonalization;

        return $this;
    }

    /**
     * Gets queryType
     *
     * @return \Algolia\AlgoliaSearch\Model\Search\QueryType|null
     */
    public function getQueryType()
    {
        return $this->container['queryType'] ?? null;
    }

    /**
     * Sets queryType
     *
     * @param \Algolia\AlgoliaSearch\Model\Search\QueryType|null $queryType queryType
     *
     * @return self
     */
    public function setQueryType($queryType)
    {
        $this->container['queryType'] = $queryType;

        return $this;
    }

    /**
     * Gets removeWordsIfNoResults
     *
     * @return \Algolia\AlgoliaSearch\Model\Search\RemoveWordsIfNoResults|null
     */
    public function getRemoveWordsIfNoResults()
    {
        return $this->container['removeWordsIfNoResults'] ?? null;
    }

    /**
     * Sets removeWordsIfNoResults
     *
     * @param \Algolia\AlgoliaSearch\Model\Search\RemoveWordsIfNoResults|null $removeWordsIfNoResults removeWordsIfNoResults
     *
     * @return self
     */
    public function setRemoveWordsIfNoResults($removeWordsIfNoResults)
    {
        $this->container['removeWordsIfNoResults'] = $removeWordsIfNoResults;

        return $this;
    }

    /**
     * Gets advancedSyntax
     *
     * @return bool|null
     */
    public function getAdvancedSyntax()
    {
        return $this->container['advancedSyntax'] ?? null;
    }

    /**
     * Sets advancedSyntax
     *
     * @param bool|null $advancedSyntax enables the advanced query syntax
     *
     * @return self
     */
    public function setAdvancedSyntax($advancedSyntax)
    {
        $this->container['advancedSyntax'] = $advancedSyntax;

        return $this;
    }

    /**
     * Gets optionalWords
     *
     * @return string[]|null
     */
    public function getOptionalWords()
    {
        return $this->container['optionalWords'] ?? null;
    }

    /**
     * Sets optionalWords
     *
     * @param string[]|null $optionalWords a list of words that should be considered as optional when found in the query
     *
     * @return self
     */
    public function setOptionalWords($optionalWords)
    {
        $this->container['optionalWords'] = $optionalWords;

        return $this;
    }

    /**
     * Gets disableExactOnAttributes
     *
     * @return string[]|null
     */
    public function getDisableExactOnAttributes()
    {
        return $this->container['disableExactOnAttributes'] ?? null;
    }

    /**
     * Sets disableExactOnAttributes
     *
     * @param string[]|null $disableExactOnAttributes list of attributes on which you want to disable the exact ranking criterion
     *
     * @return self
     */
    public function setDisableExactOnAttributes($disableExactOnAttributes)
    {
        $this->container[
            'disableExactOnAttributes'
        ] = $disableExactOnAttributes;

        return $this;
    }

    /**
     * Gets exactOnSingleWordQuery
     *
     * @return \Algolia\AlgoliaSearch\Model\Search\ExactOnSingleWordQuery|null
     */
    public function getExactOnSingleWordQuery()
    {
        return $this->container['exactOnSingleWordQuery'] ?? null;
    }

    /**
     * Sets exactOnSingleWordQuery
     *
     * @param \Algolia\AlgoliaSearch\Model\Search\ExactOnSingleWordQuery|null $exactOnSingleWordQuery exactOnSingleWordQuery
     *
     * @return self
     */
    public function setExactOnSingleWordQuery($exactOnSingleWordQuery)
    {
        $this->container['exactOnSingleWordQuery'] = $exactOnSingleWordQuery;

        return $this;
    }

    /**
     * Gets alternativesAsExact
     *
     * @return \Algolia\AlgoliaSearch\Model\Search\AlternativesAsExact[]|null
     */
    public function getAlternativesAsExact()
    {
        return $this->container['alternativesAsExact'] ?? null;
    }

    /**
     * Sets alternativesAsExact
     *
     * @param \Algolia\AlgoliaSearch\Model\Search\AlternativesAsExact[]|null $alternativesAsExact list of alternatives that should be considered an exact match by the exact ranking criterion
     *
     * @return self
     */
    public function setAlternativesAsExact($alternativesAsExact)
    {
        $this->container['alternativesAsExact'] = $alternativesAsExact;

        return $this;
    }

    /**
     * Gets advancedSyntaxFeatures
     *
     * @return \Algolia\AlgoliaSearch\Model\Search\AdvancedSyntaxFeatures[]|null
     */
    public function getAdvancedSyntaxFeatures()
    {
        return $this->container['advancedSyntaxFeatures'] ?? null;
    }

    /**
     * Sets advancedSyntaxFeatures
     *
     * @param \Algolia\AlgoliaSearch\Model\Search\AdvancedSyntaxFeatures[]|null $advancedSyntaxFeatures allows you to specify which advanced syntax features are active when ‘advancedSyntax' is enabled
     *
     * @return self
     */
    public function setAdvancedSyntaxFeatures($advancedSyntaxFeatures)
    {
        $this->container['advancedSyntaxFeatures'] = $advancedSyntaxFeatures;

        return $this;
    }

    /**
     * Gets distinct
     *
     * @return int|null
     */
    public function getDistinct()
    {
        return $this->container['distinct'] ?? null;
    }

    /**
     * Sets distinct
     *
     * @param int|null $distinct enables de-duplication or grouping of results
     *
     * @return self
     */
    public function setDistinct($distinct)
    {
        if (!is_null($distinct) && $distinct > 4) {
            throw new \InvalidArgumentException(
                'invalid value for $distinct when calling IndexSettings., must be smaller than or equal to 4.'
            );
        }
        if (!is_null($distinct) && $distinct < 0) {
            throw new \InvalidArgumentException(
                'invalid value for $distinct when calling IndexSettings., must be bigger than or equal to 0.'
            );
        }

        $this->container['distinct'] = $distinct;

        return $this;
    }

    /**
     * Gets synonyms
     *
     * @return bool|null
     */
    public function getSynonyms()
    {
        return $this->container['synonyms'] ?? null;
    }

    /**
     * Sets synonyms
     *
     * @param bool|null $synonyms whether to take into account an index's synonyms for a particular search
     *
     * @return self
     */
    public function setSynonyms($synonyms)
    {
        $this->container['synonyms'] = $synonyms;

        return $this;
    }

    /**
     * Gets replaceSynonymsInHighlight
     *
     * @return bool|null
     */
    public function getReplaceSynonymsInHighlight()
    {
        return $this->container['replaceSynonymsInHighlight'] ?? null;
    }

    /**
     * Sets replaceSynonymsInHighlight
     *
     * @param bool|null $replaceSynonymsInHighlight whether to highlight and snippet the original word that matches the synonym or the synonym itself
     *
     * @return self
     */
    public function setReplaceSynonymsInHighlight($replaceSynonymsInHighlight)
    {
        $this->container[
            'replaceSynonymsInHighlight'
        ] = $replaceSynonymsInHighlight;

        return $this;
    }

    /**
     * Gets minProximity
     *
     * @return int|null
     */
    public function getMinProximity()
    {
        return $this->container['minProximity'] ?? null;
    }

    /**
     * Sets minProximity
     *
     * @param int|null $minProximity precision of the proximity ranking criterion
     *
     * @return self
     */
    public function setMinProximity($minProximity)
    {
        if (!is_null($minProximity) && $minProximity > 7) {
            throw new \InvalidArgumentException(
                'invalid value for $minProximity when calling IndexSettings., must be smaller than or equal to 7.'
            );
        }
        if (!is_null($minProximity) && $minProximity < 1) {
            throw new \InvalidArgumentException(
                'invalid value for $minProximity when calling IndexSettings., must be bigger than or equal to 1.'
            );
        }

        $this->container['minProximity'] = $minProximity;

        return $this;
    }

    /**
     * Gets responseFields
     *
     * @return string[]|null
     */
    public function getResponseFields()
    {
        return $this->container['responseFields'] ?? null;
    }

    /**
     * Sets responseFields
     *
     * @param string[]|null $responseFields Choose which fields to return in the API response. This parameters applies to search and browse queries.
     *
     * @return self
     */
    public function setResponseFields($responseFields)
    {
        $this->container['responseFields'] = $responseFields;

        return $this;
    }

    /**
     * Gets maxFacetHits
     *
     * @return int|null
     */
    public function getMaxFacetHits()
    {
        return $this->container['maxFacetHits'] ?? null;
    }

    /**
     * Sets maxFacetHits
     *
     * @param int|null $maxFacetHits Maximum number of facet hits to return during a search for facet values. For performance reasons, the maximum allowed number of returned values is 100.
     *
     * @return self
     */
    public function setMaxFacetHits($maxFacetHits)
    {
        if (!is_null($maxFacetHits) && $maxFacetHits > 100) {
            throw new \InvalidArgumentException(
                'invalid value for $maxFacetHits when calling IndexSettings., must be smaller than or equal to 100.'
            );
        }

        $this->container['maxFacetHits'] = $maxFacetHits;

        return $this;
    }

    /**
     * Gets attributeCriteriaComputedByMinProximity
     *
     * @return bool|null
     */
    public function getAttributeCriteriaComputedByMinProximity()
    {
        return $this->container['attributeCriteriaComputedByMinProximity'] ??
            null;
    }

    /**
     * Sets attributeCriteriaComputedByMinProximity
     *
     * @param bool|null $attributeCriteriaComputedByMinProximity when attribute is ranked above proximity in your ranking formula, proximity is used to select which searchable attribute is matched in the attribute ranking stage
     *
     * @return self
     */
    public function setAttributeCriteriaComputedByMinProximity(
        $attributeCriteriaComputedByMinProximity
    ) {
        $this->container[
            'attributeCriteriaComputedByMinProximity'
        ] = $attributeCriteriaComputedByMinProximity;

        return $this;
    }

    /**
     * Gets renderingContent
     *
     * @return object|null
     */
    public function getRenderingContent()
    {
        return $this->container['renderingContent'] ?? null;
    }

    /**
     * Sets renderingContent
     *
     * @param object|null $renderingContent Content defining how the search interface should be rendered. Can be set via the settings for a default value and can be overridden via rules.
     *
     * @return self
     */
    public function setRenderingContent($renderingContent)
    {
        $this->container['renderingContent'] = $renderingContent;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param int $offset Offset
     *
     * @return bool
     */
    public function offsetExists($offset)
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param int $offset Offset
     *
     * @return mixed|null
     */
    public function offsetGet($offset)
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param int|null $offset Offset
     * @param mixed    $value  Value to be set
     *
     * @return void
     */
    public function offsetSet($offset, $value)
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param int $offset Offset
     *
     * @return void
     */
    public function offsetUnset($offset)
    {
        unset($this->container[$offset]);
    }
}
