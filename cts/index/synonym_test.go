package index

import (
	"io"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/wait"
	"github.com/algolia/algoliasearch-client-go/v3/cts"
	"github.com/stretchr/testify/require"
)

func TestSynonym(t *testing.T) {
	t.Parallel()
	_, index, _ := cts.InitSearchClient1AndIndex(t)

	{
		res, err := index.SaveObjects([]map[string]string{
			{"console": "Sony PlayStation <PLAYSTATIONVERSION>"},
			{"console": "Nintendo Switch"},
			{"console": "Nintendo Wii U"},
			{"console": "Nintendo Game Boy Advance"},
			{"console": "Microsoft Xbox"},
			{"console": "Microsoft Xbox 360"},
			{"console": "Microsoft Xbox One"},
		}, opt.AutoGenerateObjectIDIfNotExist(true))
		require.NoError(t, err)
		require.NoError(t, res.Wait())
	}

	synonyms := []search.Synonym{
		search.NewRegularSynonym("gba", "gba", "gameboy advance", "game boy advance"),
		search.NewOneWaySynonym("wii_to_wii_u", "wii", "wii U"),
		search.NewAltCorrection2("psone", "psone", "playstation1"),
		search.NewAltCorrection1("ps4", "ps4", "playstation4"),
		search.NewPlaceholder("playstation_version_placeholder", "<PLAYSTATIONVERSION>", "1", "One", "2", "3", "4", "4 Pro"),
	}

	g := wait.NewGroup()

	{
		res, err := index.SaveSynonym(synonyms[0])
		require.NoError(t, err)
		g.Collect(res)
	}

	{
		res, err := index.SaveSynonyms(synonyms[1:])
		require.NoError(t, err)
		g.Collect(res)
	}

	require.NoError(t, g.Wait())

	for _, synonym := range synonyms {
		found, err := index.GetSynonym(synonym.ObjectID())
		require.NoError(t, err, "should find synonym whose objectID is %s", synonym.ObjectID())
		require.Equal(t, synonym, found)
	}

	{
		res, err := index.SearchSynonyms("", opt.Page(0), opt.HitsPerPage(10))
		require.NoError(t, err)

		found, err := res.Synonyms()
		require.NoError(t, err)
		require.ElementsMatch(t, synonyms, found)
	}

	{
		var found []search.Synonym
		it, err := index.BrowseSynonyms()
		require.NoError(t, err)

		for {
			syn, err := it.Next()
			if err != nil {
				require.Equal(t, io.EOF, err)
				break
			}
			found = append(found, syn)
		}
		require.ElementsMatch(t, found, synonyms)
	}

	{
		res, err := index.DeleteSynonym(synonyms[0].ObjectID(), false)
		require.NoError(t, err)
		require.NoError(t, res.Wait())

		_, err = index.GetSynonym(synonyms[0].ObjectID())
		require.Error(t, err)
	}

	{
		res, err := index.ClearSynonyms(false)
		require.NoError(t, err)
		require.NoError(t, res.Wait())
	}

	{
		res, err := index.SearchSynonyms("", opt.Page(0), opt.HitsPerPage(10))
		require.NoError(t, err)

		found, err := res.Synonyms()
		require.NoError(t, err)
		require.Equal(t, 0, len(found))
	}
}
