package transport

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
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
			value = string(jsonValue)
		}
		values.Add(key, value)
	}

	return values.Encode()
}

func URLDecode(data []byte, itf interface{}) error {
	dataStr := string(data)

	values, err := url.ParseQuery(dataStr)
	if err != nil {
		return fmt.Errorf("cannot parse query %q: %v", dataStr, err)
	}

	m := make(map[string]interface{})
	for k, v := range values {
		if len(v) == 1 {
			m[k] = convertValue(v[0])
		} else {
			var s []interface{}
			for _, value := range v {
				s = append(s, convertValue(value))
			}
			m[k] = s
		}
	}

	data, err = json.Marshal(m)
	if err != nil {
		return fmt.Errorf("cannot encode temporary map %#v: %v", m, err)
	}

	err = json.Unmarshal(data, &itf)
	return err
}

func convertValue(v string) interface{} {
	if len(v) >= 2 && v[0] == '[' && v[len(v)-1] == ']' {
		var s []interface{}
		for _, v := range strings.Split(v[1:len(v)-1], ",") {
			s = append(s, convertValue(v))
		}
		return s
	}

	trimmed := strings.Trim(v, `"`)
	if b, err := strconv.ParseBool(trimmed); err == nil {
		return b
	}
	if i, err := strconv.ParseInt(trimmed, 10, 64); err == nil {
		return i
	}
	if f, err := strconv.ParseFloat(trimmed, 64); err == nil {
		return f
	}
	return trimmed
}
