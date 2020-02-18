package recommendation

import (
	"fmt"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/call"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/region"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/transport"
)

func defaultHosts(r region.Region) (hosts []*transport.StatefulHost) {
	switch r {
	case region.EU, region.US:
		hosts = append(hosts, transport.NewStatefulHost(fmt.Sprintf("recommendation.%s.algolia.com", r), call.IsReadWrite))
	default:
		hosts = append(hosts, transport.NewStatefulHost("recommendation.us.algolia.com", call.IsReadWrite))
	}
	return
}
