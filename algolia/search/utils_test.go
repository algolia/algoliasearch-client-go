package search

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type CompanyWithCorrectObjectID struct {
	ObjectID string `json:"objectID"`
	Company  string `json:"company"`
}

type CompanyWithWrongObjectID struct {
	ObjectID string `json:"objectid"`
	Company  string `json:"company"`
}

type CompanyWithoutObjectID struct {
	Company string `json:"company"`
}

type CompanyWithCustomerSerializer struct {
	Company string
}

func (c CompanyWithCustomerSerializer) MarshalJSON() ([]byte, error) {
	data := fmt.Sprintf(`{"objectID":"%s"}`, c.Company)
	return []byte(data), nil
}

func TestHasObjectIDField(t *testing.T) {
	m := map[string]interface{}{"company": "algolia"}
	require.False(t, hasObjectID(m))

	m["objectID"] = ""
	require.False(t, hasObjectID(m))

	m["objectID"] = 42
	require.True(t, hasObjectID(m))

	m["objectID"] = 42.3
	require.True(t, hasObjectID(m))

	m["objectID"] = "one"
	require.True(t, hasObjectID(m))

	require.False(t, hasObjectID(CompanyWithoutObjectID{"algolia"}))
	require.False(t, hasObjectID(CompanyWithWrongObjectID{"one", "algolia"}))
	require.True(t, hasObjectID(CompanyWithCorrectObjectID{"one", "algolia"}))
	require.True(t, hasObjectID(CompanyWithCustomerSerializer{"algolia"}))

	require.False(t, hasObjectID(nil))
}
