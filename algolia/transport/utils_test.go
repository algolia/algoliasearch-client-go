package transport

import (
	"fmt"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/call"
	"github.com/stretchr/testify/require"
)

func TestShuffle(t *testing.T) {
	var hosts []*StatefulHost
	for i := 0; i < 100; i++ {
		hosts = append(hosts, NewStatefulHost(fmt.Sprintf("%d", i), call.IsRead))
	}

	shuffled := Shuffle(hosts)
	require.ElementsMatch(t, shuffled, hosts)

	shuffledElementsCount := 0
	for i := 0; i < len(hosts); i++ {
		if hosts[i].host != shuffled[i].host {
			shuffledElementsCount++
		}
	}
	require.NotEqual(t, 0, shuffledElementsCount)
}
