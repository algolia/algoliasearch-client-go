package index

import (
	"io"
	"sync"
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

	{
		var wg sync.WaitGroup

		for _, synonym := range synonyms {
			wg.Add(1)
			expected := synonym
			go func(wg *sync.WaitGroup, expected search.Synonym) {
				defer wg.Done()
				found, err := index.GetSynonym(expected.ObjectID())
				require.NoError(t, err, "should find synonym whose objectID is %s", expected.ObjectID())
				require.Equal(t, expected, found)
			}(&wg, expected)
		}

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			res, err := index.SearchSynonyms("", opt.Page(0), opt.HitsPerPage(10))
			require.NoError(t, err)

			found, err := res.Synonyms()
			require.NoError(t, err)
			require.ElementsMatch(t, synonyms, found)
		}(&wg)

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
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
		}(&wg)

		wg.Wait()
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
