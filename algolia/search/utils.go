package search

import (
	"encoding/json"

	"github.com/algolia/algoliasearch-client-go/algolia/call"
	"github.com/algolia/algoliasearch-client-go/algolia/transport"
)

func defaultHosts(appID string) (hosts []*transport.StatefulHost) {
	hosts = append(hosts, transport.NewStatefulHost(appID+"-dsn.algolia.net", call.IsRead))
	hosts = append(hosts, transport.NewStatefulHost(appID+".algolia.net", call.IsWrite))
	hosts = append(hosts, transport.Shuffle(
		[]*transport.StatefulHost{
			transport.NewStatefulHost(appID+"-1.algolianet.com", call.IsReadWrite),
			transport.NewStatefulHost(appID+"-2.algolianet.com", call.IsReadWrite),
			transport.NewStatefulHost(appID+"-3.algolianet.com", call.IsReadWrite),
		},
	)...)
	return
}

func hasObjectIDField(object interface{}) bool {
	data, err := json.Marshal(object)
	if err != nil {
		return false
	}
	var m map[string]interface{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return false
	}
	objectID, ok := m["objectID"]
	if !ok {
		return false
	}

	switch objectID.(type) {
	case string, float64:
		return true
	default:
		return false
	}
}
