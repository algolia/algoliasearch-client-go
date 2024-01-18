package transport

import (
	"time"

	"github.com/algolia/algoliasearch-client-go/v4/algolia/compression"
)

type Configuration struct {
	AppID  string
	ApiKey string

	Hosts          []string
	DefaultHeader  map[string]string
	UserAgent      string
	Debug          bool
	Requester      Requester
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	ConnectTimeout time.Duration
	Compression    compression.Compression
}
