// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package abtesting

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"slices"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/algolia/algoliasearch-client-go/v4/algolia/call"
	"github.com/algolia/algoliasearch-client-go/v4/algolia/compression"
	"github.com/algolia/algoliasearch-client-go/v4/algolia/transport"
)

// APIClient manages communication with the A/B Testing API API v2.0.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	appID     string
	cfg       *AbtestingConfiguration
	transport *transport.Transport
}

// NewClient creates a new API client with appID, apiKey and region.
func NewClient(appID, apiKey string, region Region) (*APIClient, error) {
	return NewClientWithConfig(AbtestingConfiguration{
		Configuration: transport.Configuration{
			AppID:         appID,
			ApiKey:        apiKey,
			DefaultHeader: make(map[string]string),
			UserAgent:     getUserAgent(),
			Requester:     transport.NewDefaultRequester(nil),
		},
		Region: region,
	})
}

// NewClientWithConfig creates a new API client with the given configuration to fully customize the client behaviour.
func NewClientWithConfig(cfg AbtestingConfiguration) (*APIClient, error) {
	if cfg.AppID == "" {
		return nil, errors.New("`appId` is missing.")
	}
	if cfg.ApiKey == "" {
		return nil, errors.New("`apiKey` is missing.")
	}
	if len(cfg.Hosts) == 0 {
		if cfg.Region != "" && !slices.Contains(allowedRegions[:], string(cfg.Region)) {
			return nil, fmt.Errorf("`region` must be one of the following: %s", strings.Join(allowedRegions[:], ", "))
		}
		cfg.Hosts = getDefaultHosts(cfg.Region)
	}
	if cfg.UserAgent == "" {
		cfg.UserAgent = getUserAgent()
	}

	return &APIClient{
		appID: cfg.AppID,
		cfg:   &cfg,
		transport: transport.New(
			cfg.Configuration,
		),
	}, nil
}

func getDefaultHosts(r Region) []transport.StatefulHost {
	if r == "" {
		return []transport.StatefulHost{transport.NewStatefulHost("https", "analytics.algolia.com", call.IsReadWrite)}
	}

	return []transport.StatefulHost{transport.NewStatefulHost("https", strings.ReplaceAll("analytics.{region}.algolia.com", "{region}", string(r)), call.IsReadWrite)}
}

func getUserAgent() string {
	return fmt.Sprintf("Algolia for Go (4.0.0-beta.29); Go (%s); Abtesting (4.0.0-beta.29)", runtime.Version())
}

// AddDefaultHeader adds a new HTTP header to the default header in the request.
func (c *APIClient) AddDefaultHeader(key string, value string) {
	c.cfg.DefaultHeader[key] = value
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request, useReadTransporter bool) (*http.Response, []byte, error) {
	callKind := call.Write
	if useReadTransporter || request.Method == http.MethodGet {
		callKind = call.Read
	}

	resp, body, err := c.transport.Request(request.Context(), request, callKind)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to do request: %w", err)
	}

	return resp, body, nil
}

// Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior.
func (c *APIClient) GetConfiguration() *AbtestingConfiguration {
	return c.cfg
}

// prepareRequest build the request.
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody any,
	headerParams map[string]string,
	queryParams url.Values,
) (req *http.Request, err error) {
	body, err := setBody(postBody, c.cfg.Compression)
	if err != nil {
		return nil, fmt.Errorf("failed to set the body: %w", err)
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the path: %w", err)
	}

	var queryString []string
	for k, v := range queryParams {
		for _, value := range v {
			queryString = append(queryString, k+"="+value)
		}
	}

	url.RawQuery = strings.Join(queryString, "&")

	// Generate a new request

	// weird nil typing
	var bodyReader io.Reader
	if body != nil {
		bodyReader = body
	}
	req, err = http.NewRequest(method, url.String(), bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		for h, v := range headerParams {
			req.Header.Add(h, v)
		}
	}

	contentType := "application/json"

	// Add the user agent to the request.
	req.Header.Add("User-Agent", c.cfg.UserAgent)
	req.Header.Add("X-Algolia-Application-Id", c.cfg.AppID)
	req.Header.Add("X-Algolia-API-Key", c.cfg.ApiKey)
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Accept", contentType)

	if ctx != nil {
		// add context to the request
		req = req.WithContext(ctx)
	}

	for header, value := range c.cfg.DefaultHeader {
		req.Header.Add(header, value)
	}

	return req, nil
}

func (c *APIClient) decode(v any, b []byte) error {
	if len(b) == 0 {
		return nil
	}
	if s, ok := v.(*string); ok {
		*s = string(b)
		return nil
	}

	if actualObj, ok := v.(interface{ GetActualInstance() any }); ok { // oneOf schemas
		if unmarshalObj, ok := actualObj.(interface{ UnmarshalJSON([]byte) error }); ok { // make sure it has UnmarshalJSON defined
			if err := unmarshalObj.UnmarshalJSON(b); err != nil {
				return fmt.Errorf("failed to unmarshal one of in response body: %w", err)
			}
		} else {
			return errors.New("Unknown type with GetActualInstance but no unmarshalObj.UnmarshalJSON defined")
		}
	} else if err := json.Unmarshal(b, v); err != nil { // simple model
		return fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return nil
}

// Prevent trying to import "fmt".
func reportError(format string, a ...any) error {
	return fmt.Errorf(format, a...)
}

// A wrapper for strict JSON decoding.
func newStrictDecoder(data []byte) *json.Decoder { 
	dec := json.NewDecoder(bytes.NewBuffer(data))
	dec.DisallowUnknownFields()
	return dec
}

// A wrapper for validating a struct, returns nil if value is not a struct.
func validateStruct(v any) error { 
	err := validator.New().Struct(v)
	validationErrors, ok := err.(validator.ValidationErrors)
	if ok && len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}

// Set request body from an any.
func setBody(body any, c compression.Compression) (*bytes.Buffer, error) {
	if body == nil {
		return nil, nil
	}

	bodyBuf := &bytes.Buffer{}
	var err error

	switch c {
	case compression.GZIP:
		gzipWriter := gzip.NewWriter(bodyBuf)
		defer gzipWriter.Close()
		err = json.NewEncoder(gzipWriter).Encode(body)
	default:
		if reader, ok := body.(io.Reader); ok {
			_, err = bodyBuf.ReadFrom(reader)
		} else if b, ok := body.([]byte); ok {
			_, err = bodyBuf.Write(b)
		} else if s, ok := body.(string); ok {
			_, err = bodyBuf.WriteString(s)
		} else if s, ok := body.(*string); ok {
			_, err = bodyBuf.WriteString(*s)
		} else {
			err = json.NewEncoder(bodyBuf).Encode(body)
		}
	}

	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %w", err)
	}

	if bodyBuf.Len() == 0 {
		return nil, errors.New("Invalid body type, or empty body")
	}
	return bodyBuf, nil
}

type APIError struct {
	Message string
	Status  int
}

func (e APIError) Error() string {
	return fmt.Sprintf("API error [%d] %s", e.Status, e.Message)
}
