package search

import "github.com/algolia/algoliasearch-client-go/v3/algolia/wait"

type IndexInterface interface {
	// Misc
	WaitTask(taskID int64, opts ...interface{}) error
	WaitRecommendTask(taskID int64, opts ...interface{}) error
	GetStatus(taskID int64, opts ...interface{}) (res TaskStatusRes, err error)
	GetRecommendStatus(taskID int64, opts ...interface{}) (res TaskStatusRes, err error)
	GetAppID() string
	GetName() string
	ClearObjects(opts ...interface{}) (res UpdateTaskRes, err error)
	Delete(opts ...interface{}) (res DeleteTaskRes, err error)
	Exists() (exists bool, err error)

	// Indexing
	GetObject(objectID string, object interface{}, opts ...interface{}) error
	GetObjects(objectIDs []string, objects interface{}, opts ...interface{}) error
	SaveObject(object interface{}, opts ...interface{}) (res SaveObjectRes, err error)
	SaveObjects(objects interface{}, opts ...interface{}) (res GroupBatchRes, err error)
	PartialUpdateObject(object interface{}, opts ...interface{}) (res UpdateTaskRes, err error)
	PartialUpdateObjects(objects interface{}, opts ...interface{}) (res GroupBatchRes, err error)
	DeleteObject(objectID string, opts ...interface{}) (res DeleteTaskRes, err error)
	DeleteObjects(objectIDs []string, opts ...interface{}) (res BatchRes, err error)
	DeleteBy(opts ...interface{}) (res UpdateTaskRes, err error)
	Batch(operations []BatchOperation, opts ...interface{}) (res BatchRes, err error)

	// Query rules
	GetRule(objectID string, opts ...interface{}) (rule Rule, err error)
	SaveRule(rule Rule, opts ...interface{}) (res UpdateTaskRes, err error)
	SaveRules(rules []Rule, opts ...interface{}) (res UpdateTaskRes, err error)
	ClearRules(opts ...interface{}) (res UpdateTaskRes, err error)
	DeleteRule(objectID string, opts ...interface{}) (res UpdateTaskRes, err error)

	// Synonyms
	GetSynonym(objectID string, opts ...interface{}) (synonym Synonym, err error)
	SaveSynonym(synonym Synonym, opts ...interface{}) (res UpdateTaskRes, err error)
	SaveSynonyms(synonyms []Synonym, opts ...interface{}) (res UpdateTaskRes, err error)
	ClearSynonyms(opts ...interface{}) (res UpdateTaskRes, err error)
	DeleteSynonym(objectID string, opts ...interface{}) (res DeleteTaskRes, err error)

	// Browsing
	BrowseObjects(opts ...interface{}) (*ObjectIterator, error)
	BrowseRules(opts ...interface{}) (*RuleIterator, error)
	BrowseSynonyms(opts ...interface{}) (*SynonymIterator, error)

	// Replacing
	ReplaceAllObjects(objects interface{}, opts ...interface{}) (*wait.Group, error)
	ReplaceAllRules(rules []Rule, opts ...interface{}) (UpdateTaskRes, error)
	ReplaceAllSynonyms(synonyms []Synonym, opts ...interface{}) (UpdateTaskRes, error)

	// Searching
	Search(query string, opts ...interface{}) (res QueryRes, err error)
	SearchForFacetValues(facet, query string, opts ...interface{}) (res SearchForFacetValuesRes, err error)
	SearchRules(query string, opts ...interface{}) (res SearchRulesRes, err error)
	SearchSynonyms(query string, opts ...interface{}) (res SearchSynonymsRes, err error)

	// Settings
	GetSettings(opts ...interface{}) (settings Settings, err error)
	SetSettings(settings Settings, opts ...interface{}) (res UpdateTaskRes, err error)
}

var _ IndexInterface = &Index{}
