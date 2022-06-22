/**
 * Removes stop (common) words from the query before executing it. RemoveStopWords is used in conjunction with the queryLanguages setting. List: language ISO codes for which ignoring plurals should be enabled. This list will override any values that you may have set in queryLanguages. True: enables the stop word functionality, ensuring that stop words are removed from consideration in a search. The languages supported here are either every language, or those set by queryLanguages. False: disables stop word functionality, allowing stop words to be taken into account in a search.
 */
export type RemoveStopWords = string[] | boolean;
