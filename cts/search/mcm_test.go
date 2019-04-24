package search

import (
	"strings"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/cts"
	"github.com/stretchr/testify/require"
)

func TestMCM(t *testing.T) {
	t.Parallel()
	client := cts.InitSearchClientMCM(t)

	// Make sure we have at least 2 clusters and retrieve the first one
	res, err := client.ListClusters()
	require.NoError(t, err)
	require.True(t, len(res.Clusters) > 1)
	cluster := res.Clusters[0]

	userID := cts.GenerateCanonicalPrefixName()
	userID = strings.Replace(userID, ":", "-", -1)
	userID = strings.Replace(userID, "_", "-", -1)

	// Assign the userID and
	{
		_, err := client.AssignUserID(userID, cluster.ClusterName)
		require.NoError(t, err)
	}

	// Check that userID was properly assigned (using get/search/list)
	{
		cts.Retry(func() bool {
			_, err := client.GetUserID(userID)
			return err == nil
		})

		_, err := client.SearchUserIDs(userID)
		require.NoError(t, err)

		found := false
		page := 0
		hitsPerPage := 100

		for !found {
			res, err := client.ListUserIDs(
				opt.Page(page),
				opt.HitsPerPage(hitsPerPage),
			)
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
		cts.Retry(func() bool {
			_, err := client.RemoveUserID(userID)
			return err == nil
		})
	}

	// Check that userID was properly removed (using get)
	{
		cts.Retry(func() bool {
			_, err := client.GetUserID(userID)
			return err != nil
		})
	}

	// Remove old userIDs
	{
		var toRemove []string
		page := 0
		hitsPerPage := 100
		today := cts.TodayDate()

		for {
			res, err := client.ListUserIDs(opt.Page(page), opt.HitsPerPage(hitsPerPage))
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
