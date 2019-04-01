package search

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"

	iopt "github.com/algolia/algoliasearch-client-go/algolia/internal/opt"
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/algolia/transport"
)

func GenerateSecuredAPIKey(apiKey string, opts ...interface{}) (key string, err error) {
	message := transport.URLEncode(newSecuredAPIKeyParams(opts...))

	h := hmac.New(sha256.New, []byte(apiKey))
	h.Write([]byte(message))
	securedKey := hex.EncodeToString(h.Sum(nil))

	key = base64.StdEncoding.EncodeToString([]byte(securedKey + message))
	return
}

type securedAPIKeyParams struct {
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
