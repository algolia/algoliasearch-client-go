package search_index

import (
	"strings"
	"testing"

	"github.com/algolia/algoliasearch-client-go/it"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"

	"github.com/stretchr/testify/require"
)

func TestMCM(t *testing.T) {
	t.Parallel()
	client := it.InitSearchClientMCM(t)

	// Make sure we have at least 2 clusters and retrieve the first one
	clusters, err := client.ListClusters()
	require.NoError(t, err)
	require.True(t, len(clusters) > 1)
	cluster := clusters[0]

	userID := it.GenerateCanonicalPrefixName()
	userID = strings.Replace(userID, ":", "-", -1)
	userID = strings.Replace(userID, "_", "-", -1)

	// Assign the userID and
	{
		_, err := client.AssignUserID(userID, cluster.ClusterName)
		require.NoError(t, err)
	}

	// Check that userID was properly assigned (using get/search/list)
	{
		it.Retry(func() bool {
			_, err := client.GetUserID(userID)
			return err == nil
		})

		_, err := client.SearchUserIDs(userID, algoliasearch.Map{})
		require.NoError(t, err)

		page := 0
		hitsPerPage := 100
		found := false

		for !found {
			res, err := client.ListUserIDs(page, hitsPerPage)
			require.NoError(t, err)

			for _, u := range res.UserIDs {
				if u.ID == userID {
					found = true
					break
				}
			}

			if len(res.UserIDs) < hitsPerPage {
				break
			}
			page++
		}

		require.True(t, found)
	}

	// Retrieve the Top10 userIDs
	{
		res, err := client.GetTopUserIDs()
		require.NoError(t, err)
		require.True(t, len(res.PerCluster) > 0)
	}

	// Remove the previously assigned userID
	{
		it.Retry(func() bool {
			_, err := client.RemoveUserID(userID)
			return err == nil
		})
	}

	// Check that userID was properly removed (using get)
	{
		it.Retry(func() bool {
			_, err := client.GetUserID(userID)
			return err != nil
		})
	}

	// Remove old userIDs
	{
		var toRemove []string
		page := 0
		hitsPerPage := 100
		today := it.TodayDate()

		for {
			res, err := client.ListUserIDs(page, hitsPerPage)
			require.NoError(t, err)
			for _, u := range res.UserIDs {
				if strings.HasPrefix(u.ID, "go-") &&
					!strings.HasPrefix(u.ID, "go-"+today) {
					toRemove = append(toRemove, u.ID)
				}
			}

			if len(res.UserIDs) < hitsPerPage {
				break
			}
			page++
		}

		for _, u := range toRemove {
			_, _ = client.RemoveUserID(u)
		}
	}
}
