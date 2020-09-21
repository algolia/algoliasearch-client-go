package extra

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/errs"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

func TestIntermediateErrors(t *testing.T) {
	algoliaErr := errs.AlgoliaErr{
		Status:  500,
		Message: "There was an internal server error",
	}

	client := search.NewClientWithConfig(search.Configuration{
		Requester: &httpNetworkErrorRequester{
			responseErr: algoliaErr,
		},
	})

	_, err := client.ListIndices(opt.ExposeIntermediateNetworkErrors(true))
	require.Error(t, err)

	noMoreHostToTryErr, ok := err.(*errs.NoMoreHostToTryErr)
	require.True(t, ok)
	errors := noMoreHostToTryErr.IntermediateNetworkErrors()
	require.Len(t, errors, 4)

	for _, e := range errors {
		require.Equal(t, algoliaErr.Error(), e.Error())
	}
}

type httpNetworkErrorRequester struct {
	responseErr errs.AlgoliaErr
}

func (r *httpNetworkErrorRequester) Request(req *http.Request) (*http.Response, error) {
	data, _ := json.Marshal(r.responseErr)
	return &http.Response{
		Status:        strconv.Itoa(r.responseErr.Status),
		StatusCode:    r.responseErr.Status,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewBuffer(data)),
		ContentLength: int64(len(data)),
		Request:       req,
		Header:        make(http.Header, 0),
	}, nil
}
