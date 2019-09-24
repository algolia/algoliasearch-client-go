package search

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"

	iopt "github.com/algolia/algoliasearch-client-go/v3/algolia/internal/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/transport"
)

// GenerateSecuredAPIKey generates a public API key intended to restrict access
// to certain records. This new key is built upon the existing key named
// `apiKey` and the following options:
//
//   - opt.Filters (string of filters to apply automatically on search queries)
//   - opt.Referers (string slice of allowed referers)
//   - opt.RestrictIndices (comma-separated string list of the indices to restrict)
//   - opt.RestrictSources (string of the allowed IPv4 network)
//   - opt.UserToken (string identifier generally used to rate-limit users per IP)
//   - opt.ValidUntil (timestamp of the expiration date)
//
// More details here:
// https://www.algolia.com/doc/api-reference/api-methods/generate-secured-api-key/?language=python#parameters
func GenerateSecuredAPIKey(apiKey string, opts ...interface{}) (string, error) {
	h := hmac.New(sha256.New, []byte(apiKey))

	message := transport.URLEncode(newSecuredAPIKeyParams(opts...))
	_, err := h.Write([]byte(message))
	if err != nil {
		return "", err
	}

	checksum := hex.EncodeToString(h.Sum(nil))
	key := base64.StdEncoding.EncodeToString([]byte(checksum + message))

	return key, nil
}

type securedAPIKeyParams struct {
	// Filters is already available through the composition with QueryParams.
	Referers        *opt.ReferersOption        `json:"referers,omitempty"`
	RestrictIndices *opt.RestrictIndicesOption `json:"restrictIndices,omitempty"`
	RestrictSources *opt.RestrictSourcesOption `json:"restrictSources,omitempty"`
	UserToken       *opt.UserTokenOption       `json:"userToken,omitempty"`
	ValidUntil      int64                      `json:"validUntil,omitempty"`
	QueryParams
}

func newSecuredAPIKeyParams(opts ...interface{}) securedAPIKeyParams {
	var validUntilInt int64
	if validUntil := iopt.ExtractValidUntil(opts...); validUntil != nil {
		validUntilInt = validUntil.Get().Unix()
	}
	return securedAPIKeyParams{
		Referers:        iopt.ExtractReferers(opts...),
		RestrictIndices: iopt.ExtractRestrictIndices(opts...),
		RestrictSources: iopt.ExtractRestrictSources(opts...),
		UserToken:       iopt.ExtractUserToken(opts...),
		ValidUntil:      validUntilInt,
		QueryParams:     newQueryParams(opts...),
	}
}
