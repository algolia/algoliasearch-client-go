package debug

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

var debug bool

// Enable enables debug logging. After this call, all Algolia HTTP requests and
// responses will be displayed. A minor overhead is expected when debug logging
// is enabled.
func Enable() {
	debug = true
}

// Disable disables debug logging. After this call, Algolia requests and
// responses will not be logged anymore. The minor overhead introduced when
// debug logging is enabled should not happen anymore.
func Disable() {
	debug = false
}

// Display displays the given parameter on the standard output in a custom way,
// depending on the given input type. This function is internally used by the
// Algolia API client to display, for instance, HTTP requests and responses when
// debug logging is enabled.
func Display(itf interface{}) {
	if !debug {
		return
	}
	switch v := itf.(type) {
	case *http.Request:
		printRequest(v)
	case *http.Response:
		printResponse(v)
	default:
		fmt.Printf("do not know how to debug-print %#v\n", v)
	}
}

func Println(a ...interface{}) {
	Printf("%s", fmt.Sprintln(a...))
}

func Printf(format string, a ...interface{}) {
	if !debug {
		return
	}
	msg := fmt.Sprintf(format, a...)
	fmt.Printf("> ALGOLIA DEBUG: %s", msg)
}

func copyReadCloser(r io.ReadCloser) (io.ReadCloser, string) {
	if r == nil {
		return nil, ""
	}
	data, err := ioutil.ReadAll(r)
	_ = r.Close()
	if err != nil {
		return nil, ""
	}
	return ioutil.NopCloser(bytes.NewReader(data)), string(data)
}

func prettyPrintJSON(input, prefix string) string {
	var b bytes.Buffer
	err := json.Indent(&b, []byte(input), prefix, "  ")
	if err != nil {
		return input
	}
	return strings.TrimSuffix(b.String(), "\n")
}

func printRequest(req *http.Request) {
	var body string
	req.Body, body = copyReadCloser(req.Body)
	body = prettyPrintJSON(body, "\t")
	fmt.Printf("> ALGOLIA DEBUG request:\n\tmethod=%q\n\turl=%q\n", req.Method, req.URL)
	for k, v := range req.Header {
		str := strings.Join(v, ",")
		if strings.Contains(strings.ToLower(k), "algolia") {
			str = strings.Repeat("*", len(str))
		}
		fmt.Printf("\theader=%s:%q\n", k, str)
	}
	fmt.Printf("\tbody=\n\t%s\n", body)
}

func printResponse(res *http.Response) {
	var body string
	if res != nil {
		res.Body, body = copyReadCloser(res.Body)
	}
	body = prettyPrintJSON(body, "\t")
	fmt.Printf("> ALGOLIA DEBUG response:\n\tbody=\n\t%s\n", body)
}
