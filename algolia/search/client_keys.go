package search

import (
	"net/http"

	"github.com/algolia/algoliasearch-client-go/algolia/call"
	"github.com/algolia/algoliasearch-client-go/algolia/errs"
)

func (c *Client) GetAPIKey(keyID string, opts ...interface{}) (key Key, err error) {
	path := c.path("/keys/%s", keyID)
	err = c.transport.Request(&key, http.MethodGet, path, nil, call.Read, opts...)
	key.Value = keyID
	return
}

func (c *Client) AddAPIKey(key Key, opts ...interface{}) (res CreateKeyRes, err error) {
	path := c.path("/keys")
	err = c.transport.Request(&res, http.MethodPost, path, key, call.Write, opts...)
	res.wait = c.waitKeyIsAvailable(res.Key)
	return
}

func (c *Client) UpdateAPIKey(key Key, opts ...interface{}) (res UpdateKeyRes, err error) {
	path := c.path("/keys/%s", key.Value)
	err = c.transport.Request(&res, http.MethodPut, path, key, call.Write, opts...)
	res.wait = c.waitKeyHasChanged(key)
	return
}

func (c *Client) DeleteAPIKey(keyID string, opts ...interface{}) (res DeleteKeyRes, err error) {
	path := c.path("/keys/%s", keyID)
	err = c.transport.Request(&res, http.MethodDelete, path, nil, call.Write, opts...)
	res.wait = c.waitKeyIsNotAvailable(keyID)
	return
}

func (c *Client) RestoreAPIKey(keyID string, opts ...interface{}) (res RestoreKeyRes, err error) {
	path := c.path("/keys/%s/restore", keyID)
	err = c.transport.Request(&res, http.MethodPost, path, nil, call.Write, opts...)
	res.wait = c.waitKeyIsAvailable(keyID)
	return
}

func (c *Client) ListAPIKeys(opts ...interface{}) (res ListAPIKeysRes, err error) {
	path := c.path("/keys")
	err = c.transport.Request(&res, http.MethodGet, path, nil, call.Read, opts...)
	return
}

func (c *Client) waitKeyIsAvailable(keyID string) func() error {
	return func() error {
		return waitWithRetry(func() (bool, error) {
			_, err := c.GetAPIKey(keyID)
			if err == nil {
				return true, nil
			}
			if _, ok := errs.IsAlgoliaHTTPErrWithCode(err, http.StatusNotFound); ok {
				return false, nil
			}
			return true, err
		})
	}
}

func (c *Client) waitKeyIsNotAvailable(keyID string) func() error {
	return func() error {
		return waitWithRetry(func() (bool, error) {
			_, err := c.GetAPIKey(keyID)
			if err == nil {
				return false, nil
			}
			if _, ok := errs.IsAlgoliaHTTPErrWithCode(err, http.StatusNotFound); ok {
				return true, nil
			}
			return true, err
		})
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
		})
	}
}
