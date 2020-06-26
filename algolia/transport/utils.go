package transport

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/debug"
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
			if v == nil {
				value = ""
			} else {
				jsonValue, _ := json.Marshal(v)
				value = string(jsonValue)
			}
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
		m[k] = convertValue(v[0])
	}

	data, err = json.Marshal(m)
	if err != nil {
		return fmt.Errorf("cannot encode temporary map %#v: %v", m, err)
	}

	err = json.Unmarshal(data, &itf)
	return err
}

func convertValue(v string) interface{} {
	if len(v) == 0 {
		return nil
	}
	if len(v) == 1 {
		return v
	}
	if isArray(v) {
		var arr []interface{}
		err := json.Unmarshal([]byte(v), &arr)
		if err != nil {
			debug.Printf("cannot decode array value %s: %v\n", v, err)
		}
		return arr
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

func isArray(v string) bool {
	firstChar := v[0]
	lastChar := v[len(v)-1]
	return firstChar == '[' && lastChar == ']'
}
