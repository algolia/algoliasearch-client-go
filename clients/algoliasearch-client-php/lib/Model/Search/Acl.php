<?php

// This file is generated, manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.

namespace Algolia\AlgoliaSearch\Model\Search;

/**
 * Acl Class Doc Comment
 *
 * @category Class
 * @description List of rights for the API key. The following rights can be used:  addObject: allows to add/update an object in the index (copy/move index are also allowed with this right). analytics: allows to retrieve the analytics through the Analytics API. browse: allows to retrieve all index content via the browse API. deleteIndex: allows to delete or clear index content. deleteObject: allows to delete objects from the index. editSettings: allows to change index settings. listIndexes: allows to list all accessible indices. logs: allows to get the logs. recommendation: Allows usage of the Personalization dashboard and the Recommendation API. search: allows to search the index. seeUnretrievableAttributes: disable unretrievableAttributes feature for all operations returning records. settings: allows to get index settings.
 *
 * @package Algolia\AlgoliaSearch
 */
class Acl
{
    /**
     * Possible values of this enum
     */
    const ADD_OBJECT = 'addObject';

    const ANALYTICS = 'analytics';

    const BROWSE = 'browse';

    const DELETE_OBJECT = 'deleteObject';

    const DELETE_INDEX = 'deleteIndex';

    const EDIT_SETTINGS = 'editSettings';

    const LIST_INDEXES = 'listIndexes';

    const LOGS = 'logs';

    const PERSONALIZATION = 'personalization';

    const RECOMMENDATION = 'recommendation';

    const SEARCH = 'search';

    const SEE_UNRETRIEVABLE_ATTRIBUTES = 'seeUnretrievableAttributes';

    const SETTINGS = 'settings';

    const USAGE = 'usage';

    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public static function getAllowableEnumValues()
    {
        return [
            self::ADD_OBJECT,
            self::ANALYTICS,
            self::BROWSE,
            self::DELETE_OBJECT,
            self::DELETE_INDEX,
            self::EDIT_SETTINGS,
            self::LIST_INDEXES,
            self::LOGS,
            self::PERSONALIZATION,
            self::RECOMMENDATION,
            self::SEARCH,
            self::SEE_UNRETRIEVABLE_ATTRIBUTES,
            self::SETTINGS,
            self::USAGE,
        ];
    }
}
