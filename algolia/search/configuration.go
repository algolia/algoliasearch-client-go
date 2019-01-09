package search

import (
	"time"

	"github.com/algolia/algoliasearch-client-go/algolia/transport"
)

type Configuration struct {
	AppID        string
	ApiKey       string
	Hosts        []string
	MaxBatchSize int
	Requester    transport.Requester
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
