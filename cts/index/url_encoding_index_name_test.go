package index

import (
	"fmt"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/wait"
	"github.com/algolia/algoliasearch-client-go/v3/cts"
)

func TestEnableURLEncodingIndexName(t *testing.T) {
	t.Parallel()

	var (
		client              = cts.InitSearchClient1(t)
		indexNames          []string
		baseIndexName       = cts.GenerateIndexName(t)
		forbiddenCharacters = "$&*,/;\\`|~"
	)

	for c := 0; c < 128; c++ {
		if !unicode.IsPrint(rune(c)) ||
			unicode.IsNumber(rune(c)) ||
			unicode.IsLetter(rune(c)) ||
			strings.Contains(forbiddenCharacters, fmt.Sprint(c)) {
			continue
		}
		indexNames = append(indexNames, baseIndexName+fmt.Sprint(c))
	}

	g := wait.NewGroup()

	for _, indexName := range indexNames {
		res, err := client.InitIndex(indexName).SaveObject(map[string]string{
			"objectID": indexName,
		})
		require.NoError(t, err, "should save object in index %q", indexName)
		g.Collect(res)
	}

	require.NoError(t, g.Wait())

	res, err := client.ListIndices()
	require.NoError(t, err)

	var listedIndexNames []string
	for _, index := range res.Items {
		listedIndexNames = append(listedIndexNames, index.Name)
	}

	for _, indexName := range indexNames {
		found := assert.Contains(t, listedIndexNames, indexName)
		if !found {
			fmt.Printf("%s not found among listed indices (%d):\n", indexName, len(listedIndexNames))
			for _, listedIndexName := range listedIndexNames {
				if strings.Contains(listedIndexName, baseIndexName) {
					fmt.Printf("> %s\n", listedIndexName)
				}
			}
		}
	}
}
