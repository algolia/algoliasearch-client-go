package search

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStopword_MarshalJSON(t *testing.T) {
	stopWord := NewStopword("myObject", "en", "bad-word", "enabled")
	serialized, err := json.Marshal(stopWord)
	require.NoError(t, err)
	require.Equal(t, string(serialized), "{\"objectID\":\"myObject\",\"language\":\"en\",\"word\":\"bad-word\",\"state\":\"enabled\"}")

	stopWord = NewStopword("myObject", "en", "bad-word", "disabled")
	serialized, err = json.Marshal(stopWord)
	require.NoError(t, err)
	require.Equal(t, string(serialized), "{\"objectID\":\"myObject\",\"language\":\"en\",\"word\":\"bad-word\",\"state\":\"disabled\"}")
}

func TestStopword_UnmarshalJSON(t *testing.T) {

	// Implicitly enabled
	payload := `{
		"objectID": "myObject",
		"language": "en",
		"word": "bad-word"
	}`
	var s Stopword
	err := json.Unmarshal([]byte(payload), &s)
	require.NoError(t, err)
	require.Equal(t, s.Language(), "en")
	require.Equal(t, s.ObjectID(), "myObject")
	require.Equal(t, s.Word, "bad-word")
	require.Equal(t, s.State, "enabled")

	// Explicitly enabled
	payload = `{
		"objectID": "myObject",
		"language": "en",
		"word": "bad-word",
		"state": "enabled"
	}`

	err = json.Unmarshal([]byte(payload), &s)
	require.NoError(t, err)
	require.Equal(t, s.Language(), "en")
	require.Equal(t, s.ObjectID(), "myObject")
	require.Equal(t, s.Word, "bad-word")
	require.Equal(t, s.State, "enabled")

	// Explicitly disabled
	payload = `{
		"objectID": "myObject",
		"language": "en",
		"word": "bad-word",
		"state": "disabled"
	}`

	err = json.Unmarshal([]byte(payload), &s)
	require.NoError(t, err)
	require.Equal(t, s.Language(), "en")
	require.Equal(t, s.ObjectID(), "myObject")
	require.Equal(t, s.Word, "bad-word")
	require.Equal(t, s.State, "disabled")

}
