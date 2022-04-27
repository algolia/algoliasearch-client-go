<?php

namespace Algolia\AlgoliaSearch\Model\Recommend;

use Algolia\AlgoliaSearch\ObjectSerializer;

/**
 * ExactOnSingleWordQuery Class Doc Comment
 *
 * @category Class
 * @description Controls how the exact ranking criterion is computed when the query contains only one word.
 * @package Algolia\AlgoliaSearch
 */
class ExactOnSingleWordQuery
{
    /**
     * Possible values of this enum
     */
    public const ATTRIBUTE = 'attribute';

    public const NONE = 'none';

    public const WORD = 'word';

    /**
     * Gets allowable values of the enum
     * @return string[]
     */
    public static function getAllowableEnumValues()
    {
        return [self::ATTRIBUTE, self::NONE, self::WORD];
    }
}
