package search

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/transport"
)

// Key represents an Algolia API key used by the API to identify and/or restrict
// calls.
type Key struct {
	ACL                    []string       `json:"acl,omitempty"`
	CreatedAt              time.Time      `json:"-"`
	Description            string         `json:"description,omitempty"`
	Indexes                []string       `json:"indexes,omitempty"`
	MaxQueriesPerIPPerHour int            `json:"maxQueriesPerIPPerHour,omitempty"`
	MaxHitsPerQuery        int            `json:"maxHitsPerQuery,omitempty"`
	Referers               []string       `json:"referers,omitempty"`
	QueryParameters        KeyQueryParams `json:"-"`
	Validity               time.Duration  `json:"-"`
	Value                  string         `json:"-"`
}

// SetQueryParameters properly encodes any given query parameters into the
// QueryParameters field of the Key.
func (k *Key) SetQueryParameters(opts ...interface{}) *Key {
	k.QueryParameters = newKeyQueryParams(opts...)
	return k
}

func (k Key) MarshalJSON() ([]byte, error) {
	type key Key
	return json.Marshal(struct {
		key
		QueryParameters string `json:"queryParameters,omitempty"`
		Validity        int64  `json:"validity,omitempty"`
	}{
		key:             key(k),
		QueryParameters: transport.URLEncode(k.QueryParameters),
		Validity:        int64(k.Validity.Seconds()),
	})
}

func (k *Key) UnmarshalJSON(data []byte) error {
	if string(data) == jsonNull {
		return nil
	}

	type key Key
	var tmp struct {
		key
		CreatedAt       int64  `json:"createdAt"`
		QueryParameters string `json:"queryParameters"`
		Validity        int64  `json:"validity"`
		Value           string `json:"value"`
	}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	*k = Key(tmp.key)
	k.CreatedAt = time.Unix(tmp.CreatedAt, 0)
	k.Validity = time.Duration(tmp.Validity) * time.Second
	k.Value = tmp.Value

	err = transport.URLDecode(
		[]byte(tmp.QueryParameters),
		&k.QueryParameters,
	)
	if err != nil {
		return fmt.Errorf("cannot decode QueryParameters %q: %v", tmp.QueryParameters, err)
	}

	return nil
}

// Equal returns true if the Keys are equal. It returns false otherwise.
func (k Key) Equal(k2 Key) bool {
	k.CreatedAt, k2.CreatedAt = time.Time{}, time.Time{}
	k.Validity, k2.Validity = 0, 0
	return reflect.DeepEqual(k, k2)
}
