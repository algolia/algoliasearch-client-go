package transport

import (
	"encoding/json"
	"math/rand"
	"net/url"
	"strconv"
	"time"

	"github.com/algolia/algoliasearch-client-go/algolia/debug"
)

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

func Shuffle(hosts []*StatefulHost) []*StatefulHost {
	if hosts == nil {
		return nil
	}
	shuffled := make([]*StatefulHost, len(hosts))
	for i, v := range rand.Perm(len(hosts)) {
		shuffled[i] = hosts[v]
	}
	return shuffled
}

func URLEncode(itf interface{}) string {
	jsonPayload, err := json.Marshal(itf)
	if err != nil {
		debug.Printf("cannot marshal payload to URL-encode: %v\n", err)
		return ""
	}
	var m map[string]interface{}
	err = json.Unmarshal(jsonPayload, &m)
	if err != nil {
		debug.Printf("cannot unmarshal payload to URL-encode: %v\n", err)
		return ""
	}

	var (
		values = make(url.Values)
		value  string
	)

	for key, itf := range m {
		switch v := itf.(type) {
		case string:
			value = v
		case float64:
			value = strconv.FormatFloat(v, 'f', -1, 64)
		case int:
			value = strconv.Itoa(v)
		default:
			jsonValue, _ := json.Marshal(v)
			value = string(jsonValue[:])
		}
		values.Add(key, value)
	}

	return values.Encode()
}
