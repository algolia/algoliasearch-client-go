// This file is generated, manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.

/**
 * Treats singular, plurals, and other forms of declensions as matching terms. IgnorePlurals is used in conjunction with the queryLanguages setting. List: language ISO codes for which ignoring plurals should be enabled. This list will override any values that you may have set in queryLanguages. True: enables the ignore plurals functionality, where singulars and plurals are considered equivalent (foot = feet). The languages supported here are either every language (this is the default, see list of languages below), or those set by queryLanguages. False: disables ignore plurals, where singulars and plurals are not considered the same for matching purposes (foot will not find feet).
 */
export type IgnorePlurals = string[] | boolean;
