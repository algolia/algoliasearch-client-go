package search

import (
	"encoding/json"
	"reflect"
	"time"

	"github.com/algolia/algoliasearch-client-go/algolia/transport"
)

// Key represents an Algolia API key used by the API to identify and/or restrict
// calls.
type Key struct {
	ACL                    []string      `json:"acl,omitempty"`
	CreatedAt              time.Time     `json:"-"`
	Description            string        `json:"description,omitempty"`
	Indexes                []string      `json:"indexes,omitempty"`
	MaxQueriesPerIPPerHour int           `json:"maxQueriesPerIPPerHour,omitempty"`
	MaxHitsPerQuery        int           `json:"maxHitsPerQuery,omitempty"`
	Referers               []string      `json:"referers,omitempty"`
	RestrictSources        string        `json:"restrictSources,omitempty"`
	QueryParameters        string        `json:"queryParameters,omitempty"`
	Validity               time.Duration `json:"validity,omitempty"`
	Value                  string        `json:"-"`
}

// SetQueryParameters properly encodes any given query parameters into the
// QueryParameters field of the Key.
func (k *Key) SetQueryParameters(opts ...interface{}) {
	k.QueryParameters = transport.URLEncode(newQueryParams(opts...))
}

func (k *Key) UnmarshalJSON(data []byte) error {
	if string(data) == jsonNull {
		return nil
	}

	type key Key
	var tmpKey struct {
		key
		CreatedAt int64  `json:"createdAt"`
		Value     string `json:"value"`
	}
	err := json.Unmarshal(data, &tmpKey)
	if err != nil {
		return err
	}

	k.ACL = tmpKey.ACL
	k.CreatedAt = time.Unix(tmpKey.CreatedAt, 0)
	k.Description = tmpKey.Description
	k.Indexes = tmpKey.Indexes
	k.MaxQueriesPerIPPerHour = tmpKey.MaxQueriesPerIPPerHour
	k.MaxHitsPerQuery = tmpKey.MaxHitsPerQuery
	k.Referers = tmpKey.Referers
	k.RestrictSources = tmpKey.RestrictSources
	k.QueryParameters = tmpKey.QueryParameters
	k.Validity = tmpKey.Validity
	k.Value = tmpKey.Value

	return nil
}

// Equal returns true if the Keys are equal. It returns false otherwise.
func (k Key) Equal(k2 Key) bool {
	k.CreatedAt, k2.CreatedAt = time.Time{}, time.Time{}
	k.Validity, k2.Validity = 0, 0
	return reflect.DeepEqual(k, k2)
}
