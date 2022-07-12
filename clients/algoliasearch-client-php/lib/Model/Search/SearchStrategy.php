<?php

// This file is generated, manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.

namespace Algolia\AlgoliaSearch\Model\Search;

/**
 * SearchStrategy Class Doc Comment
 *
 * @category Class
 * @package Algolia\AlgoliaSearch
 */
class SearchStrategy
{
    /**
     * Possible values of this enum
     */
    const NONE = 'none';

    const STOP_IF_ENOUGH_MATCHES = 'stopIfEnoughMatches';

    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public static function getAllowableEnumValues()
    {
        return [self::NONE, self::STOP_IF_ENOUGH_MATCHES];
    }
}
