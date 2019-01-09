package rand

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

func Duration(max time.Duration) time.Duration {
	nbNanoseconds := 1 + int(rand.Int63n(max.Nanoseconds()))
	return time.Duration(nbNanoseconds) * time.Nanosecond
}
