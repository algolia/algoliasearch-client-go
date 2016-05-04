package algoliasearch

type ListKeysRes struct {
	Keys []KeyRes `json:"keys"`
}

type KeyRes struct {
	ACL      []string `json:"acl"`
	Validity int64    `json:"validity"`
	Value    string   `json:"value"`
}

type AddKeyRes struct {
	CreatedAt string `json:"createdAt"`
	Key       string `json:"key"`
}

type UpdateKeyRes struct {
	UpdatedAt string `json:"updatedAt"`
	Key       string `json:"key"`
}

type GetKeyRes struct {
	ACL      []string `json:"acl"`
	Validity int64    `json:"validity"`
	Value    string   `json:"value"`
}

type DeleteKeyRes struct {
	DeletedAt string `json:"deletedAt"`
}
