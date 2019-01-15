package search

import (
	"fmt"
	"time"

	"github.com/algolia/algoliasearch-client-go/algolia/rand"
	"github.com/algolia/algoliasearch-client-go/algolia/transport"
)

type Index struct {
	appID        string
	name         string
	maxBatchSize int
	transport    *transport.Transport
}

func newIndex(appID, name string, maxBatchSize int, transport *transport.Transport) *Index {
	return &Index{
		appID:        appID,
		name:         name,
		maxBatchSize: maxBatchSize,
		transport:    transport,
	}
}

func (i *Index) path(format string, a ...interface{}) string {
	prefix := fmt.Sprintf("/1/indexes/%s", i.name)
	suffix := fmt.Sprintf(format, a...)
	return prefix + suffix
}

func (i *Index) waitTask(taskID int) error {
	var maxDuration = time.Second

	for {
		res, err := i.GetStatus(taskID)
		if err != nil {
			return err
		}

		if res.Status == "published" {
			return nil
		}

		sleepDuration := rand.Duration(maxDuration)
		time.Sleep(sleepDuration)

		// Increase the upper boundary used to generate the sleep duration
		if maxDuration < 10*time.Minute {
			maxDuration *= 2
			if maxDuration > 10*time.Minute {
				maxDuration = 10 * time.Minute
			}
		}
	}
}
