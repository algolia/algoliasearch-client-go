package search

import (
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/cts"
)

func TestMCM(t *testing.T) {
	t.Parallel()
	client := cts.InitSearchClientMCM(t)

	// Make sure we have at least 2 clusters and retrieve the first one
	res, err := client.ListClusters()
	require.NoError(t, err)
	require.True(t, len(res.Clusters) > 1)
	cluster := res.Clusters[0]

	userIDPrefix := cts.GenerateCanonicalPrefixName()
	userIDPrefix = strings.Replace(userIDPrefix, ":", "-", -1)
	userIDPrefix = strings.Replace(userIDPrefix, "_", "-", -1)

	userIDs := []string{
		userIDPrefix + "-0",
		userIDPrefix + "-1",
		userIDPrefix + "-2",
	}

	// Assign a single userID
	{
		_, err := client.AssignUserID(userIDs[0], cluster.ClusterName)
		require.NoError(t, err)
	}

	// Assign multiple userIDs
	{
		_, err := client.AssignUserIDs(
			[]string{userIDs[1], userIDs[2]},
			cluster.ClusterName,
		)
		require.NoError(t, err)
	}

	// Check that userIDs were properly assigned (using get/search/list)
	{

		// Check (asynchronously) that userIDs were properly assigned (using get/search)
		{
			var wg sync.WaitGroup

			for _, userID := range userIDs {
				wg.Add(1)
				go func(userID string) {
					cts.Retry(func() bool {
						_, err := client.GetUserID(userID)
						return err == nil
					})
					wg.Done()
				}(userID)

				wg.Add(1)
				go func(userID string) {

					cts.Retry(func() bool {
						_, err := client.SearchUserIDs(userID)
						return err == nil
					})
					wg.Done()
				}(userID)
			}

			wg.Wait()
		}

		// Check that userIDs were properly assigned (using list)
		{
			found := make(map[string]int)
			page := 0
			hitsPerPage := 100

			for {
				res, err := client.ListUserIDs(
					opt.Page(page),
					opt.HitsPerPage(hitsPerPage),
				)
				require.NoError(t, err)

				for _, u := range res.UserIDs {
					for _, userID := range userIDs {
						if u.ID == userID {
							found[u.ID]++
						}
					}
				}
				if len(found) == len(userIDs) {
					break
				}
				if len(res.UserIDs) < hitsPerPage {
					break
				}
				page++
			}

			require.Len(t, found, len(userIDs))
		}
	}

	// Retrieve the Top10 userIDs
	{
		res, err := client.GetTopUserIDs()
		require.NoError(t, err)
		require.True(t, len(res.PerCluster) > 0)
	}

	// Remove (asynchronously) the previously assigned userIDs
	{
		var wg sync.WaitGroup

		for _, userID := range userIDs {
			wg.Add(1)
			go func(userID string) {
				cts.Retry(func() bool {
					_, err := client.RemoveUserID(userID)
					return err == nil
				})
				wg.Done()
			}(userID)
		}

		wg.Wait()
	}

	// Check (asynchronously) that userID was properly removed (using get)
	{
		var wg sync.WaitGroup

		for _, userID := range userIDs {
			wg.Add(1)
			go func(userID string) {
				cts.Retry(func() bool {
					_, err := client.GetUserID(userID)
					return err != nil
				})
				wg.Done()
			}(userID)
		}

		wg.Wait()
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
