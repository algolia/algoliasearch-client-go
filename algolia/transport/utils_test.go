package transport

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/call"
)

func TestShuffle(t *testing.T) {
	var hosts []*StatefulHost
	for i := 0; i < 100; i++ {
		hosts = append(hosts, NewStatefulHost(fmt.Sprintf("%d", i), call.IsRead))
	}

	shuffled := Shuffle(hosts)
	require.ElementsMatch(t, shuffled, hosts)

	shuffledElementsCount := 0
	for i := 0; i < len(hosts); i++ {
		if hosts[i].host != shuffled[i].host {
			shuffledElementsCount++
		}
	}
	require.NotEqual(t, 0, shuffledElementsCount)
}

func TestURLEncodeDecode(t *testing.T) {
	for _, c := range []struct {
		input   interface{}
		encoded string
	}{
		{
			struct {
				Book     string `json:"book"`
				ObjectID string `json:"objectID"`
			}{"harry potter", "one"},
			"book=harry+potter&objectID=one",
		},
		{
			struct {
				Book     string `json:"book"`
				ObjectID string `json:"objectID"`
			}{"harry potter", ""},
			"book=harry+potter&objectID=",
		},
		{
			struct {
				Book     string `json:"book"`
				ObjectID string `json:"objectID,omitempty"`
			}{"harry potter", ""},
			"book=harry+potter",
		},
		{
			struct {
				Book     *string `json:"book,omitempty"`
				ObjectID string  `json:"objectID,omitempty"`
			}{nil, ""},
			"",
		},
		{
			struct {
				Book     *string `json:"book,omitempty"`
				ObjectID string  `json:"objectID,omitempty"`
			}{new(string), ""},
			"book=",
		},
		{
			struct {
				Book     string   `json:"book"`
				Keywords []string `json:"keywords"`
				Price    float64  `json:"price"`
			}{"harry potter", []string{"magic", "fiction"}, 19.99},
			"book=harry+potter&keywords=%5B%22magic%22%2C%22fiction%22%5D&price=19.99",
		},
	} {
		encoded1 := URLEncode(c.input)
		require.Equal(t, c.encoded, encoded1, "wrong URL-encoded string for input %#v", c.input)

		decoded := c.input
		err := URLDecode([]byte(encoded1), &decoded)
		require.NoError(t, err, "cannot decode back the input %#v once encoded (%s)", c.input, encoded1)

		encoded2 := URLEncode(decoded)
		require.Equal(t, c.encoded, encoded2, "wrong URL-encoded string for input %#v", c.input)
	}
}
