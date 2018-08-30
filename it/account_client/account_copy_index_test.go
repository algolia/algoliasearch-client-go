package account_client

import (
	"testing"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/algolia/algoliasearch-client-go/it"
	"github.com/stretchr/testify/require"
)

func TestAccountCopyIndex(t *testing.T) {
	t.Parallel()
	client1, index1, indexName1 := it.InitSearchClient1AndIndex(t)
	indexName2 := indexName1 + "-copy"

	index2 := client1.InitIndex(indexName2)
	{
		account := algoliasearch.NewAccountClient()
		_, err := account.CopyIndex(index1, index2)
		require.Equal(t, algoliasearch.SameAppIDErr, err)
	}

	client2 := it.InitSearchClient2(t)
	index2 = client2.InitIndex(indexName2)

	{
		_, err := index2.Delete()
		require.NoError(t, err)
	}

	var taskIDs []int

	{
		res, err := index1.AddObject(algoliasearch.Object{"objectID": "one"})
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	{
		res, err := index1.SaveRule(algoliasearch.Rule{
			ObjectID:  "one",
			Condition: algoliasearch.NewSimpleRuleCondition(algoliasearch.Contains, "pattern"),
			Consequence: algoliasearch.RuleConsequence{
				Params: algoliasearch.Map{
					"query": algoliasearch.Map{
						"edits": []algoliasearch.Edit{algoliasearch.DeleteEdit("pattern")},
					},
				},
			},
		}, false)
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	{
		res, err := index1.SaveSynonym(algoliasearch.NewSynonym("one", []string{"one", "two"}), false)
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	{
		res, err := index1.SetSettings(algoliasearch.Map{"searchableAttributes": []string{"objectID"}})
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	it.WaitTasks(t, index1, taskIDs...)
	taskIDs = []int{}

	{
		account := algoliasearch.NewAccountClient()
		taskIDs, err := account.CopyIndex(index1, index2)
		require.NoError(t, err)
		it.WaitTasks(t, index2, taskIDs...)
	}

	{
		_, err := index2.GetObject("one", nil)
		require.NoError(t, err)

		_, err = index2.GetRule("one")
		require.NoError(t, err)

		_, err = index2.GetSynonym("one")
		require.NoError(t, err)

		settings, err := index2.GetSettings()
		require.NoError(t, err)
		require.Equal(t, []string{"objectID"}, settings.SearchableAttributes)
	}

	{
		account := algoliasearch.NewAccountClient()
		_, err := account.CopyIndex(index1, index2)
		require.Equal(t, algoliasearch.IndexAlreadyExistsErr, err)
	}
}
