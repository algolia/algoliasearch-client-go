<?php

namespace Algolia\AlgoliaSearch\Model\Search;

/**
 * SearchTypeDefault Class Doc Comment
 *
 * @category Class
 * @description Perform a search query with &#x60;default&#x60;, will search for facet values if &#x60;facet&#x60; is given.
 *
 * @package Algolia\AlgoliaSearch
 */
class SearchTypeDefault
{
    /**
     * Possible values of this enum
     */
    const _DEFAULT = 'default';

    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public static function getAllowableEnumValues()
    {
        return [self::_DEFAULT];
    }
}
