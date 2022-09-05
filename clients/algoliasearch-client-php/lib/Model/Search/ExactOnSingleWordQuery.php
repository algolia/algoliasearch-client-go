<?php

// This file is generated, manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.

namespace Algolia\AlgoliaSearch\Model\Search;

/**
 * ExactOnSingleWordQuery Class Doc Comment
 *
 * @category Class
 *
 * @description Controls how the exact ranking criterion is computed when the query contains only one word.
 *
 * @package Algolia\AlgoliaSearch
 */
class ExactOnSingleWordQuery
{
    /**
     * Possible values of this enum
     */
    const ATTRIBUTE = 'attribute';

    const NONE = 'none';

    const WORD = 'word';

    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public static function getAllowableEnumValues()
    {
        return [self::ATTRIBUTE, self::NONE, self::WORD];
    }
}
