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

type contact struct {
	Age                int
	Name               string `json:"name"`
	ObjID              string `json:"objectID,omitempty"`
	PublicField        []string
	PublicFieldWithTag []string `json:"publicFieldWithTag"`
	privateField       []string
}

func (c contact) ObjectID() string {
	return c.ObjID
}

func BenchmarkHasObjectID(b *testing.B) {
	c := contact{
		Age:                42,
		Name:               "Elon Musk",
		ObjID:              "test",
		PublicField:        []string{"field1", "field2"},
		PublicFieldWithTag: []string{"field3"},
		privateField:       []string{"field4", "field5"},
	}

	b.Run("hasObjectID", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			hasObjectID(c)
		}
	})

	b.Run("hasObjectIDWithProvider", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			hasObjectIDWithProvider(c)
		}
	})

	b.Run("hasObjectIDWithReflect", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			hasObjectIDWithReflect(c)
		}
	})

}

func TestHasObjectID_checkDifferentImplementations(t *testing.T) {
	c := contact{
		Age:                42,
		Name:               "Elon Musk",
		ObjID:              "test",
		PublicField:        []string{"field1", "field2"},
		PublicFieldWithTag: []string{"field3"},
		privateField:       []string{"field4", "field5"},
	}

	{
		okHas := hasObjectID(c)
		objectID, okGet := getObjectID(c)
		require.True(t, okHas)
		require.True(t, okGet)
		require.Equal(t, "test", objectID)
	}

	{
		okHas := hasObjectIDWithReflect(c)
		objectID, okGet := getObjectIDWithReflect(c)
		require.True(t, okHas)
		require.True(t, okGet)
		require.Equal(t, "test", objectID)
	}

	{
		okHas := hasObjectIDWithProvider(c)
		objectID, okGet := getObjectIDWithProvider(c)
		require.True(t, okHas)
		require.True(t, okGet)
		require.Equal(t, "test", objectID)
	}

}
