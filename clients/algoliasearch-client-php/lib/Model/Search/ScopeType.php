<?php

// This file is generated, manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.

namespace Algolia\AlgoliaSearch\Model\Search;

/**
 * ScopeType Class Doc Comment
 *
 * @category Class
 * @package Algolia\AlgoliaSearch
 */
class ScopeType
{
    /**
     * Possible values of this enum
     */
    const SETTINGS = 'settings';

    const SYNONYMS = 'synonyms';

    const RULES = 'rules';

    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public static function getAllowableEnumValues()
    {
        return [self::SETTINGS, self::SYNONYMS, self::RULES];
    }
}
