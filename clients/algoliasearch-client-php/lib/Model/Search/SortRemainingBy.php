<?php

// This file is generated, manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.

namespace Algolia\AlgoliaSearch\Model\Search;

/**
 * SortRemainingBy Class Doc Comment
 *
 * @category Class
 *
 * @description How to display the remaining items.   - &#x60;count&#x60;: facet count (descending).   - &#x60;alpha&#x60;: alphabetical (ascending).   - &#x60;hidden&#x60;: show only pinned values.
 *
 * @package Algolia\AlgoliaSearch
 */
class SortRemainingBy
{
    /**
     * Possible values of this enum
     */
    const COUNT = 'count';

    const ALPHA = 'alpha';

    const HIDDEN = 'hidden';

    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public static function getAllowableEnumValues()
    {
        return [self::COUNT, self::ALPHA, self::HIDDEN];
    }
}