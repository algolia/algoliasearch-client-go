package algoliasearch

import (
	"fmt"
	"math/rand"
	"time"
)

// randDuration generates a pseudo-random `time.Duration` between 1 and `max`.
func randDuration(max time.Duration) time.Duration {
	rand.Seed(time.Now().Unix())
	nbNanoseconds := 1 + int64(rand.Int63n(max.Nanoseconds()))
	return time.Duration(nbNanoseconds) * time.Nanosecond
}

func wrongType(expected, key string) error {
	return fmt.Errorf("Expected type `%s` for `%s`", expected, key)
}

func unknownField(field string) error {
	return fmt.Errorf("Unexpected field `%s`", field)
}
