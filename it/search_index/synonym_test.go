package search_index

import (
	"sync"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/algolia/algoliasearch-client-go/it"
	"github.com/stretchr/testify/require"
)

func TestSynonym(t *testing.T) {
	t.Parallel()
	_, index, _ := it.InitSearchClient1AndIndex(t)

	{
		res, err := index.AddObjects([]algoliasearch.Object{
			{"console": "Sony PlayStation <PLAYSTATIONVERSION>"},
			{"console": "Nintendo Switch"},
			{"console": "Nintendo Wii U"},
			{"console": "Nintendo Game Boy Advance"},
			{"console": "Microsoft Xbox"},
			{"console": "Microsoft Xbox 360"},
			{"console": "Microsoft Xbox One"},
		})
		require.NoError(t, err)
		it.WaitTasks(t, index, res.TaskID)
	}

	synonyms := []algoliasearch.Synonym{
		algoliasearch.NewSynonym("gba", []string{"gba", "gameboy advance", "game boy advance"}),
		algoliasearch.NewOneWaySynonym("wii_to_wii_u", "wii", []string{"wii U"}),
		algoliasearch.NewPlaceholderSynonym("playstation_version_placeholder", "<PLAYSTATIONVERSION>", []string{"1", "One", "2", "3", "4", "4 Pro"}),
		algoliasearch.NewAltCorrectionSynonym("ps4", []string{"playstation4"}, "ps4", algoliasearch.AltCorrection1),
		algoliasearch.NewAltCorrectionSynonym("psone", []string{"playstation1"}, "psone", algoliasearch.AltCorrection2),
	}

	var taskIDs []int

	{
		res, err := index.SaveSynonym(synonyms[0], false)
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	{
		res, err := index.BatchSynonyms(synonyms[1:], false, false)
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	it.WaitTasks(t, index, taskIDs...)
	taskIDs = []int{}

	{
		var wg sync.WaitGroup

		for _, synonym := range synonyms {
			wg.Add(1)
			expected := synonym
			objectID := synonym.ObjectID
			go func(wg *sync.WaitGroup, objectID string, expected algoliasearch.Synonym) {
				defer wg.Done()
				found, err := index.GetSynonym(objectID)
				require.NoError(t, err, "should find synonym whose objectID is %s", objectID)
				require.Equal(t, expected, found)
			}(&wg, objectID, expected)
		}

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			found, err := index.SearchSynonyms("", nil, 0, 10)
			require.NoError(t, err)
			for i := range found {
				found[i].HighlightResult = nil
			}
			require.ElementsMatch(t, synonyms, found)
		}(&wg)

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			var found []algoliasearch.Synonym
			it := algoliasearch.NewSynonymIterator(index)

			for {
				syn, err := it.Next()
				if err != nil {
					require.Equal(t, algoliasearch.NoMoreSynonymsErr, err)
					break
				}
				found = append(found, *syn)
			}
			require.ElementsMatch(t, found, synonyms)
		}(&wg)

		wg.Wait()
	}

	{
		res, err := index.DeleteSynonym(synonyms[0].ObjectID, false)
		require.NoError(t, err)
		it.WaitTasks(t, index, res.TaskID)

		_, err = index.GetSynonym(synonyms[0].ObjectID)
		require.Error(t, err)
	}

	{
		res, err := index.ClearSynonyms(false)
		require.NoError(t, err)
		it.WaitTasks(t, index, res.TaskID)
	}

	{
		found, err := index.SearchSynonyms("", nil, 0, 10)
		require.NoError(t, err)
		require.Equal(t, 0, len(found))
	}
}
