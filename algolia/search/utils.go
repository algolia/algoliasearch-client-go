package search

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/algolia/algoliasearch-client-go/algolia/call"
	"github.com/algolia/algoliasearch-client-go/algolia/rand"
	"github.com/algolia/algoliasearch-client-go/algolia/transport"
)

const (
	jsonNull = "null"
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

func noWait(_ int) error {
	return nil
}

func waitWithRetry(f func() (bool, error)) error {
	var maxDuration = time.Second

	for {
		done, err := f()
		if done {
			return err
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

func getObjectID(object interface{}) (string, bool) {
	data, err := json.Marshal(object)
	if err != nil {
		return "", false
	}
	var m map[string]interface{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return "", false
	}
	objectID, ok := m["objectID"]
	if !ok {
		return "", false
	}

	switch v := objectID.(type) {
	case string:
		return v, v != ""
	case float64:
		return fmt.Sprintf("%d", int(v)), true
	default:
		return "", false
	}
}

func hasObjectID(object interface{}) bool {
	_, ok := getObjectID(object)
	return ok
}

type ObjectIDProvider interface {
	ObjectID() string
}

func getObjectIDWithProvider(object interface{}) (string, bool) {
	if itf, ok := object.(ObjectIDProvider); ok {
		return itf.ObjectID(), true
	}
	return "", false
}

func hasObjectIDWithProvider(object interface{}) bool {
	_, ok := getObjectIDWithProvider(object)
	return ok
}

func getObjectIDWithReflect(object interface{}) (string, bool) {
	ve := reflect.ValueOf(object)
	t := reflect.TypeOf(object)
	for i := 0; i < ve.NumField(); i++ {
		jsonTagValue, ok := t.Field(i).Tag.Lookup("json")
		if ok && strings.Contains(jsonTagValue, "objectID") {
			return fmt.Sprintf("%v", ve.Field(i).Interface()), true
		}
	}
	return "", false
}

func hasObjectIDWithReflect(object interface{}) bool {
	_, ok := getObjectIDWithReflect(object)
	return ok
}
