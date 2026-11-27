package transport

import (
	"math/rand"
)

func Shuffle(hosts []StatefulHost) []StatefulHost {
	if hosts == nil {
		return nil
	}

	shuffled := make([]StatefulHost, len(hosts))
	for i, v := range rand.Perm(len(hosts)) { //nolint:gosec // host shuffling for load spreading, not security-sensitive
		shuffled[i] = hosts[v]
	}

	return shuffled
}
