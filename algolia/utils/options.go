package utils

import (
	"context"
	"net/url"
	"time"
)

type Options struct {
	// -- Request options for API calls
	Context      context.Context
	QueryParams  url.Values
	HeaderParams map[string]string

	// -- ChunkedBatch options
	WaitForTasks bool
	BatchSize    int

	// -- Iterable options
	MaxRetries    int
	Timeout       func(int) time.Duration
	Aggregator    func(any, error)
	IterableError *IterableError
}

// --------- Request options for API calls ---------

type RequestOption interface {
	Apply(*Options)
}

type requestOption func(*Options)

func (r requestOption) Apply(o *Options) {
	r(o)
}

func WithContext(ctx context.Context) requestOption {
	return requestOption(func(o *Options) {
		o.Context = ctx
	})
}

func WithHeaderParam(key string, value any) requestOption {
	return requestOption(func(o *Options) {
		o.HeaderParams[key] = ParameterToString(value)
	})
}

func WithQueryParam(key string, value any) requestOption {
	return requestOption(func(o *Options) {
		o.QueryParams.Set(QueryParameterToString(key), QueryParameterToString(value))
	})
}

// --------- ChunkedBatch options ---------

type ChunkedBatchOption interface {
	RequestOption
	chunkedBatch()
}

type chunkedBatchOption func(*Options)

var (
	_ ChunkedBatchOption = (*chunkedBatchOption)(nil)
	_ ChunkedBatchOption = (*requestOption)(nil)
)

func (c chunkedBatchOption) Apply(o *Options) {
	c(o)
}

func (c chunkedBatchOption) chunkedBatch() {}

func (r requestOption) chunkedBatch() {}

func WithWaitForTasks(waitForTasks bool) chunkedBatchOption {
	return chunkedBatchOption(func(o *Options) {
		o.WaitForTasks = waitForTasks
	})
}

func WithBatchSize(batchSize int) chunkedBatchOption {
	return chunkedBatchOption(func(o *Options) {
		o.BatchSize = batchSize
	})
}

// --------- Iterable options ---------.
type IterableOption interface {
	RequestOption
	iterable()
}

type iterableOption func(*Options)

var (
	_ IterableOption = (*iterableOption)(nil)
	_ IterableOption = (*requestOption)(nil)
)

func (i iterableOption) Apply(o *Options) {
	i(o)
}

func (r requestOption) iterable() {}

func (i iterableOption) iterable() {}

func WithMaxRetries(maxRetries int) iterableOption {
	return iterableOption(func(o *Options) {
		o.MaxRetries = maxRetries
	})
}

func WithTimeout(timeout func(int) time.Duration) iterableOption {
	return iterableOption(func(o *Options) {
		o.Timeout = timeout
	})
}

func WithAggregator(aggregator func(any, error)) iterableOption {
	return iterableOption(func(o *Options) {
		o.Aggregator = aggregator
	})
}

func WithIterableError(iterableError *IterableError) iterableOption {
	return iterableOption(func(o *Options) {
		o.IterableError = iterableError
	})
}

// --------- Helper to convert options ---------

func ToRequestOptions[T RequestOption](opts []T) []RequestOption {
	requestOpts := make([]RequestOption, 0, len(opts))

	for _, opt := range opts {
		requestOpts = append(requestOpts, opt)
	}

	return requestOpts
}

func ToIterableOptions(opts []ChunkedBatchOption) []IterableOption {
	iterableOpts := make([]IterableOption, 0, len(opts))

	for _, opt := range opts {
		if opt, ok := opt.(IterableOption); ok {
			iterableOpts = append(iterableOpts, opt)
		}
	}

	return iterableOpts
}
