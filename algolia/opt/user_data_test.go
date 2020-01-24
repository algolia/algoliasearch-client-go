package opt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserData_DecodeIntger(t *testing.T) {
	var i int
	userData := UserData(42)
	require.Equal(t, map[string]interface{}{}, userData.Get())
	require.NoError(t, userData.Decode(&i))
	require.Equal(t, 42, i)
}

func TestUserData_DecodeString(t *testing.T) {
	var s string
	userData := UserData("algolia")
	require.Equal(t, map[string]interface{}{}, userData.Get())
	require.NoError(t, userData.Decode(&s))
	require.Equal(t, "algolia", s)
}

func TestUserData_DecodeStringSlice(t *testing.T) {
	var s []string
	userData := UserData([]string{"algolia", "Algolia"})
	require.Equal(t, map[string]interface{}{}, userData.Get())
	require.NoError(t, userData.Decode(&s))
	require.Equal(t, []string{"algolia", "Algolia"}, s)
}

func TestUserData_DecodeMap(t *testing.T) {
	var m map[string]interface{}
	userData := UserData(map[string]interface{}{"k1": "v1", "k2": 2})
	require.Equal(t, map[string]interface{}{"k1": "v1", "k2": 2}, userData.Get())
	require.NoError(t, userData.Decode(&m))
	require.Equal(t, map[string]interface{}{"k1": "v1", "k2": float64(2)}, m)
}
