package search

import (
	"encoding/base64"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/call"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/errs"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
)

// GetAPIKey retrieves the API key identified by the given keyID.
func (c *Client) GetAPIKey(keyID string, opts ...interface{}) (key Key, err error) {
	path := c.path("/keys/%s", keyID)
	err = c.transport.Request(&key, http.MethodGet, path, nil, call.Read, opts...)
	key.Value = keyID
	return
}

// AddAPIKey creates a new API key. Once created, the key can be referenced by
// other methods via the Key field of the response which represents its keyID.
func (c *Client) AddAPIKey(key Key, opts ...interface{}) (res CreateKeyRes, err error) {
	path := c.path("/keys")
	err = c.transport.Request(&res, http.MethodPost, path, key, call.Write, opts...)
	res.wait = c.waitKeyIsAvailable(res.Key)
	return
}

// UpdateAPIKey updates the API key identified by its Value field and updates
// all its non-zero fields.
func (c *Client) UpdateAPIKey(key Key, opts ...interface{}) (res UpdateKeyRes, err error) {
	if key.Value == "" {
		err = errs.ErrMissingKeyID
		return
	}
	path := c.path("/keys/%s", key.Value)
	err = c.transport.Request(&res, http.MethodPut, path, key, call.Write, opts...)
	res.wait = c.waitKeyHasChanged(key)
	return
}

// DeleteAPIKey deletes the API key for the given keyID.
//
// To restore a deleted key, you can use RestoreAPIKey with the same keyID.
func (c *Client) DeleteAPIKey(keyID string, opts ...interface{}) (res DeleteKeyRes, err error) {
	path := c.path("/keys/%s", keyID)
	err = c.transport.Request(&res, http.MethodDelete, path, nil, call.Write, opts...)
	res.wait = c.waitKeyIsNotAvailable(keyID)
	return
}

// RestoreAPIKey restores the API key for the given keyID if it ever existed.
func (c *Client) RestoreAPIKey(keyID string, opts ...interface{}) (res RestoreKeyRes, err error) {
	path := c.path("/keys/%s/restore", keyID)
	err = c.transport.Request(&res, http.MethodPost, path, nil, call.Write, opts...)
	res.wait = c.waitKeyIsAvailable(keyID)
	return
}

// ListAPIKeys list all the API keys of the application.
func (c *Client) ListAPIKeys(opts ...interface{}) (res ListAPIKeysRes, err error) {
	path := c.path("/keys")
	err = c.transport.Request(&res, http.MethodGet, path, nil, call.Read, opts...)
	return
}

func (c *Client) GetSecuredAPIKeyRemainingValidity(keyID string, opts ...interface{}) (v time.Duration, err error) {
	if len(keyID) == 0 {
		err = errs.ErrEmptySecuredAPIKey
		return
	}

	decoded, err := base64.StdEncoding.DecodeString(keyID)
	if err != nil {
		err = errs.ErrInvalidSecuredAPIKey
		return
	}

	submatch := regexp.MustCompile(`validUntil=(\d{10})`).FindSubmatch(decoded)

	if len(submatch) != 2 {
		err = errs.ErrValidUntilNotFound
		return
	}

	// Error checking is useless here since the previous regexp already matched
	// with an integer of maximum length 10.
	ts, _ := strconv.Atoi(string(submatch[1]))

	v = time.Until(time.Unix(int64(ts), 0))
	return
}

func (c *Client) waitKeyIsAvailable(keyID string) func() error {
	return func() error {
		return waitWithRetry(func() (bool, error) {
			_, err := c.GetAPIKey(keyID)
			if err == nil {
				return true, nil
			}
			if _, ok := errs.IsAlgoliaErrWithCode(err, http.StatusNotFound); ok {
				return false, nil
			}
			return true, err
		}, opt.DefaultWaitConfiguration())
	}
}

func (c *Client) waitKeyIsNotAvailable(keyID string) func() error {
	return func() error {
		return waitWithRetry(func() (bool, error) {
			_, err := c.GetAPIKey(keyID)
			if err == nil {
				return false, nil
			}
			if _, ok := errs.IsAlgoliaErrWithCode(err, http.StatusNotFound); ok {
				return true, nil
			}
			return true, err
		}, opt.DefaultWaitConfiguration())
	}
}

func (c *Client) waitKeyHasChanged(expected Key) func() error {
	return func() error {
		return waitWithRetry(func() (bool, error) {
			found, err := c.GetAPIKey(expected.Value)
			if err != nil {
				return true, err
			}
			return expected.Equal(found), nil
		}, opt.DefaultWaitConfiguration())
	}
}
