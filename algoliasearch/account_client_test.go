package algoliasearch

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAccountClient(t *testing.T) {
	client1, index1 := initClientAndIndex(t, "TestAccountClient")

	index2 := client1.InitIndex("go-TestAccountClient")
	{
		account := NewAccountClient()
		_, err := account.CopyIndex(index1, index2)
		require.Equal(t, SameAppIDErr, err)
	}

	client2 := initClient2(t)
	index2 = client2.InitIndex("go-TestAccountClient")

	{
		_, err := index2.Delete()
		require.NoError(t, err)
	}

	var taskIDs []int

	{
		res, err := index1.AddObject(Object{"objectID": "one"})
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	{
		res, err := index1.SaveRule(Rule{
			ObjectID:  "one",
			Condition: NewSimpleRuleCondition(Contains, "pattern"),
			Consequence: RuleConsequence{
				Params: Map{
					"query": Map{
						"edits": []Edit{DeleteEdit("pattern")},
					},
				},
			},
		}, false)
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	{
		res, err := index1.SaveSynonym(NewSynonym("one", []string{"one", "two"}), false)
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	{
		res, err := index1.SetSettings(Map{"searchableAttributes": []string{"objectID"}})
		require.NoError(t, err)
		taskIDs = append(taskIDs, res.TaskID)
	}

	waitTasksAsync(t, index1, taskIDs)
	taskIDs = []int{}

	{
		account := NewAccountClient()
		taskIDs, err := account.CopyIndex(index1, index2)
		require.NoError(t, err)
		waitTasksAsync(t, index2, taskIDs)
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
		account := NewAccountClient()
		_, err := account.CopyIndex(index1, index2)
		require.Equal(t, IndexAlreadyExistsErr, err)
	}
}
