package algoliasearch

import (
	"fmt"
	"os"
	"strconv"
)

var debugLevel int

func init() {
	env := os.Getenv("ALGOLIA_DEBUG")
	if env == "" {
		return

	}
	if level, err := strconv.Atoi(env); err == nil {
		debugLevel = level
	}
}

// debug is an internal helper function used to print debugging messages to the
// standard output if the ALGOLIA_DEBUG environment variable is set. Accepted
// values for ALGOLIA_DEBUG are 1 (length-truncated output) and 2 (full
// output).
func debug(format string, a ...interface{}) {
	if debugLevel == 0 {
		return
	}
	msg := fmt.Sprintf(format, a...)
	if debugLevel == 1 && len(msg) > 120 {
		fmt.Printf("%.120s...\n", msg)
	} else {
		fmt.Println(msg)
	}

}
