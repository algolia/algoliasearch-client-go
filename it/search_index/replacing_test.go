package search_index

import (
	"testing"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/algolia/algoliasearch-client-go/it"
	"github.com/stretchr/testify/require"
)

func TestReplacing(t *testing.T) {
	t.Parallel()
	_, index, _ := it.InitSearchClient1AndIndex(t)

	var taskIDs []int

	{
		res, err := index.AddObject(algoliasearch.Object{"objectID": "one"})
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	{
		res, err := index.SaveRule(algoliasearch.Rule{
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
		res, err := index.SaveSynonym(algoliasearch.NewSynonym("one", []string{"one", "two"}), false)
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	it.WaitTasks(t, index, taskIDs...)
	taskIDs = []int{}

	{
		err := index.ReplaceAllObjects([]algoliasearch.Object{{"objectID": "two"}})
		require.NoError(t, err)
	}

	{
		res, err := index.ReplaceAllRules([]algoliasearch.Rule{
			{
				ObjectID:  "two",
				Condition: algoliasearch.NewSimpleRuleCondition(algoliasearch.Contains, "pattern"),
				Consequence: algoliasearch.RuleConsequence{
					Params: algoliasearch.Map{
						"query": algoliasearch.Map{
							"edits": []algoliasearch.Edit{algoliasearch.DeleteEdit("pattern")},
						},
					},
				},
			},
		})
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	{
		res, err := index.ReplaceAllSynonyms([]algoliasearch.Synonym{
			algoliasearch.NewSynonym("two", []string{"one", "two"}),
		})
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	it.WaitTasks(t, index, taskIDs...)

	{
		_, err := index.GetObject("one", nil)
		require.Error(t, err)

		_, err = index.GetObject("two", nil)
		require.NoError(t, err)

		_, err = index.GetRule("one")
		require.Error(t, err)

		_, err = index.GetRule("two")
		require.NoError(t, err)

		_, err = index.GetSynonym("one")
		require.Error(t, err)

		_, err = index.GetSynonym("two")
		require.NoError(t, err)
	}
}
