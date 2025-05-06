package debug

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

var debug bool
var logger *slog.Logger

// Enable enables debug logging. After this call, all Algolia HTTP requests and
// responses will be displayed. A minor overhead is expected when debug logging
// is enabled.
func Enable() {
	debug = true
}

// EnableWithSlog enables debug logging with a custom slog.Logger.
func EnableWithSlog(l *slog.Logger) {
	debug = true
	logger = l
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
	start := time.Now()
	var msg string
	switch v := itf.(type) {
	case *http.Request:
		msg = debugRequest(v)
	case *http.Response:
		msg = debugResponse(v)
	default:
		msg = fmt.Sprintf("do not know how to display %#v", v)
	}
	Println(msg)
	Print(fmt.Sprintf("took %s\n", time.Since(start)))
}

func Println(a ...interface{}) {
	Printf("%s", fmt.Sprintln(a...))
}

func Printf(format string, a ...interface{}) {
	if !debug {
		return
	}
	msg := fmt.Sprintf(format, a...)
	Print(fmt.Sprintf("> ALGOLIA DEBUG: %s", msg))
}

func Print(msg string) {
	if !debug {
		return
	}
	if logger != nil {
		logger.Debug(msg)
	} else {
		fmt.Printf(msg)
	}
}
