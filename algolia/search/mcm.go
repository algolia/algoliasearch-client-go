package search

import (
	"encoding/json"
	"fmt"

	iopt "github.com/algolia/algoliasearch-client-go/v3/algolia/internal/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
)

type ListClustersRes struct {
	Clusters []Cluster `json:"clusters"`
}

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

type HasPendingMappingsRes struct {
	Pending  bool                `json:"pending"`
	Clusters map[string][]string `json:"clusters"`
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
	ObjectID        string      `json:"objectID"`
	HighlightResult interface{} `json:"_highlightResult"`
}

func (h UserIDHit) UnmarshalHighlightResult(v interface{}) error {
	highlightResultPayload, err := json.Marshal(h.HighlightResult)
	if err != nil {
		return fmt.Errorf("cannot unmarshal HighlightResult from search userIDs response: %v", err)
	}
	return json.Unmarshal(highlightResultPayload, &v)
}

type searchUserIDsReq struct {
	Query       *opt.QueryOption       `json:"query,omitempty"`
	Cluster     *opt.ClusterOption     `json:"cluster,omitempty"`
	Page        *opt.PageOption        `json:"page,omitempty"`
	HitsPerPage *opt.HitsPerPageOption `json:"hitsPerPage,omitempty"`
}

func newSearchUserIDsReq(query string, opts ...interface{}) searchUserIDsReq {
	return searchUserIDsReq{
		Query:       opt.Query(query),
		Cluster:     iopt.ExtractCluster(opts...),
		Page:        iopt.ExtractPage(opts...),
		HitsPerPage: iopt.ExtractHitsPerPage(opts...),
	}
}
