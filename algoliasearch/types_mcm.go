package algoliasearch

type Cluster struct {
	ClusterName string `json:"clusterName"`
	NbRecords   int    `json:"nbRecords"`
	NbUserIDs   int    `json:"nbUserIDs"`
	DataSize    int    `json:"dataSize"`
}

type UserIDCore struct {
	ID        string `json:"userID"`
	NbRecords int    `json:"nbRecords"`
	DataSize  int    `json:"dataSize"`
}

type UserID struct {
	UserIDCore
	ClusterName string `json:"clusterName"`
}

type ListUserIDsRes struct {
	UserIDs     []UserID `json:"userIDs"`
	Page        int      `json:"page"`
	HitsPerPage int      `json:"hitsPerPage"`
}

type AssignUserIDRes struct {
	CreatedAt string `json:"createdAt"`
}

type RemoveUserIDRes struct {
	DeletedAt string `json:"deletedAt"`
}

type TopUserIDs struct {
	PerCluster map[string][]UserIDCore `json:"topUsers"`
}

type SearchUserIDRes struct {
	Hits        []UserIDHit `json:"hits"`
	NbHits      int         `json:"nbHits"`
	Page        int         `json:"page"`
	HitsPerPage int         `json:"hitsPerPage"`
	UpdatedAt   int         `json:"updatedAt"`
}

type UserIDHit struct {
	UserID
	ObjectID        string `json:"objectID"`
	HighlightResult Map    `json:"_highlightResult"`
}
