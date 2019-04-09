package insights

import (
	"fmt"

	"github.com/algolia/algoliasearch-client-go/algolia/call"
	"github.com/algolia/algoliasearch-client-go/algolia/region"
	"github.com/algolia/algoliasearch-client-go/algolia/transport"
)

func defaultHosts(r region.Region) (hosts []*transport.StatefulHost) {
	switch r {
	case region.EU, region.US:
		hosts = append(hosts, transport.NewStatefulHost(fmt.Sprintf("insights.%s.algolia.io", r), call.IsReadWrite))
	default:
		hosts = append(hosts, transport.NewStatefulHost("insights.algolia.io", call.IsReadWrite))
	}
	return
}
