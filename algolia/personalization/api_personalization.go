// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package personalization

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/algolia/algoliasearch-client-go/v4/algolia/transport"
	"github.com/algolia/algoliasearch-client-go/v4/algolia/utils"
)

type config struct {
	// -- Request options for API calls
	context      context.Context
	queryParams  url.Values
	headerParams map[string]string
	timeouts     transport.RequestConfiguration
}

type RequestOption interface {
	apply(*config)
}

type requestOption func(*config)

func (r requestOption) apply(c *config) {
	r(c)
}

func WithContext(ctx context.Context) requestOption {
	return requestOption(func(c *config) {
		c.context = ctx
	})
}

func WithHeaderParam(key string, value any) requestOption {
	return requestOption(func(c *config) {
		c.headerParams[key] = utils.ParameterToString(value)
	})
}

func WithQueryParam(key string, value any) requestOption {
	return requestOption(func(c *config) {
		c.queryParams.Set(utils.QueryParameterToString(key), utils.QueryParameterToString(value))
	})
}

func WithReadTimeout(timeout time.Duration) requestOption {
	return requestOption(func(c *config) {
		c.timeouts.ReadTimeout = &timeout
	})
}

func WithWriteTimeout(timeout time.Duration) requestOption {
	return requestOption(func(c *config) {
		c.timeouts.WriteTimeout = &timeout
	})
}

func WithConnectTimeout(timeout time.Duration) requestOption {
	return requestOption(func(c *config) {
		c.timeouts.ConnectTimeout = &timeout
	})
}

func (r *ApiCustomDeleteRequest) UnmarshalJSON(b []byte) error {
	req := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return fmt.Errorf("cannot unmarshal request: %w", err)
	}
	if v, ok := req["path"]; ok {
		err = json.Unmarshal(v, &r.path)
		if err != nil {
			err = json.Unmarshal(b, &r.path)
			if err != nil {
				return fmt.Errorf("cannot unmarshal path: %w", err)
			}
		}
	}
	if v, ok := req["parameters"]; ok {
		err = json.Unmarshal(v, &r.parameters)
		if err != nil {
			err = json.Unmarshal(b, &r.parameters)
			if err != nil {
				return fmt.Errorf("cannot unmarshal parameters: %w", err)
			}
		}
	}

	return nil
}

// ApiCustomDeleteRequest represents the request with all the parameters for the API call.
type ApiCustomDeleteRequest struct {
	path       string
	parameters map[string]any
}

// NewApiCustomDeleteRequest creates an instance of the ApiCustomDeleteRequest to be used for the API call.
func (c *APIClient) NewApiCustomDeleteRequest(path string) ApiCustomDeleteRequest {
	return ApiCustomDeleteRequest{
		path: path,
	}
}

// WithParameters adds the parameters to the ApiCustomDeleteRequest and returns the request for chaining.
func (r ApiCustomDeleteRequest) WithParameters(parameters map[string]any) ApiCustomDeleteRequest {
	r.parameters = parameters
	return r
}

/*
CustomDelete calls the API and returns the raw response from it.

	  This method lets you send requests to the Algolia REST API.


	Request can be constructed by NewApiCustomDeleteRequest with parameters below.
	  @param path string - Path of the endpoint, for example `1/newFeature`.
	  @param parameters map[string]any - Query parameters to apply to the current query.
	@param opts ...RequestOption - Optional parameters for the API call
	@return *http.Response - The raw response from the API
	@return []byte - The raw response body from the API
	@return error - An error if the API call fails
*/
func (c *APIClient) CustomDeleteWithHTTPInfo(r ApiCustomDeleteRequest, opts ...RequestOption) (*http.Response, []byte, error) {
	requestPath := "/{path}"
	requestPath = strings.ReplaceAll(requestPath, "{path}", utils.ParameterToString(r.path))

	if r.path == "" {
		return nil, nil, reportError("Parameter `path` is required when calling `CustomDelete`.")
	}

	conf := config{
		context:      context.Background(),
		queryParams:  url.Values{},
		headerParams: map[string]string{},
	}

	if !utils.IsNilOrEmpty(r.parameters) {
		for k, v := range r.parameters {
			conf.queryParams.Set(k, utils.QueryParameterToString(v))
		}
	}

	// optional params if any
	for _, opt := range opts {
		opt.apply(&conf)
	}

	var postBody any

	req, err := c.prepareRequest(conf.context, requestPath, http.MethodDelete, postBody, conf.headerParams, conf.queryParams)
	if err != nil {
		return nil, nil, err
	}

	return c.callAPI(req, false, conf.timeouts)
}

/*
CustomDelete casts the HTTP response body to a defined struct.

This method lets you send requests to the Algolia REST API.

Request can be constructed by NewApiCustomDeleteRequest with parameters below.

	@param path string - Path of the endpoint, for example `1/newFeature`.
	@param parameters map[string]any - Query parameters to apply to the current query.
	@return map[string]any
*/
func (c *APIClient) CustomDelete(r ApiCustomDeleteRequest, opts ...RequestOption) (*map[string]any, error) {
	var returnValue *map[string]any

	res, resBody, err := c.CustomDeleteWithHTTPInfo(r, opts...)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	if res.StatusCode >= 300 {
		return returnValue, c.decodeError(res, resBody)
	}

	err = c.decode(&returnValue, resBody)
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

func (r *ApiCustomGetRequest) UnmarshalJSON(b []byte) error {
	req := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return fmt.Errorf("cannot unmarshal request: %w", err)
	}
	if v, ok := req["path"]; ok {
		err = json.Unmarshal(v, &r.path)
		if err != nil {
			err = json.Unmarshal(b, &r.path)
			if err != nil {
				return fmt.Errorf("cannot unmarshal path: %w", err)
			}
		}
	}
	if v, ok := req["parameters"]; ok {
		err = json.Unmarshal(v, &r.parameters)
		if err != nil {
			err = json.Unmarshal(b, &r.parameters)
			if err != nil {
				return fmt.Errorf("cannot unmarshal parameters: %w", err)
			}
		}
	}

	return nil
}

// ApiCustomGetRequest represents the request with all the parameters for the API call.
type ApiCustomGetRequest struct {
	path       string
	parameters map[string]any
}

// NewApiCustomGetRequest creates an instance of the ApiCustomGetRequest to be used for the API call.
func (c *APIClient) NewApiCustomGetRequest(path string) ApiCustomGetRequest {
	return ApiCustomGetRequest{
		path: path,
	}
}

// WithParameters adds the parameters to the ApiCustomGetRequest and returns the request for chaining.
func (r ApiCustomGetRequest) WithParameters(parameters map[string]any) ApiCustomGetRequest {
	r.parameters = parameters
	return r
}

/*
CustomGet calls the API and returns the raw response from it.

	  This method lets you send requests to the Algolia REST API.


	Request can be constructed by NewApiCustomGetRequest with parameters below.
	  @param path string - Path of the endpoint, for example `1/newFeature`.
	  @param parameters map[string]any - Query parameters to apply to the current query.
	@param opts ...RequestOption - Optional parameters for the API call
	@return *http.Response - The raw response from the API
	@return []byte - The raw response body from the API
	@return error - An error if the API call fails
*/
func (c *APIClient) CustomGetWithHTTPInfo(r ApiCustomGetRequest, opts ...RequestOption) (*http.Response, []byte, error) {
	requestPath := "/{path}"
	requestPath = strings.ReplaceAll(requestPath, "{path}", utils.ParameterToString(r.path))

	if r.path == "" {
		return nil, nil, reportError("Parameter `path` is required when calling `CustomGet`.")
	}

	conf := config{
		context:      context.Background(),
		queryParams:  url.Values{},
		headerParams: map[string]string{},
	}

	if !utils.IsNilOrEmpty(r.parameters) {
		for k, v := range r.parameters {
			conf.queryParams.Set(k, utils.QueryParameterToString(v))
		}
	}

	// optional params if any
	for _, opt := range opts {
		opt.apply(&conf)
	}

	var postBody any

	req, err := c.prepareRequest(conf.context, requestPath, http.MethodGet, postBody, conf.headerParams, conf.queryParams)
	if err != nil {
		return nil, nil, err
	}

	return c.callAPI(req, false, conf.timeouts)
}

/*
CustomGet casts the HTTP response body to a defined struct.

This method lets you send requests to the Algolia REST API.

Request can be constructed by NewApiCustomGetRequest with parameters below.

	@param path string - Path of the endpoint, for example `1/newFeature`.
	@param parameters map[string]any - Query parameters to apply to the current query.
	@return map[string]any
*/
func (c *APIClient) CustomGet(r ApiCustomGetRequest, opts ...RequestOption) (*map[string]any, error) {
	var returnValue *map[string]any

	res, resBody, err := c.CustomGetWithHTTPInfo(r, opts...)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	if res.StatusCode >= 300 {
		return returnValue, c.decodeError(res, resBody)
	}

	err = c.decode(&returnValue, resBody)
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

func (r *ApiCustomPostRequest) UnmarshalJSON(b []byte) error {
	req := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return fmt.Errorf("cannot unmarshal request: %w", err)
	}
	if v, ok := req["path"]; ok {
		err = json.Unmarshal(v, &r.path)
		if err != nil {
			err = json.Unmarshal(b, &r.path)
			if err != nil {
				return fmt.Errorf("cannot unmarshal path: %w", err)
			}
		}
	}
	if v, ok := req["parameters"]; ok {
		err = json.Unmarshal(v, &r.parameters)
		if err != nil {
			err = json.Unmarshal(b, &r.parameters)
			if err != nil {
				return fmt.Errorf("cannot unmarshal parameters: %w", err)
			}
		}
	}
	if v, ok := req["body"]; ok {
		err = json.Unmarshal(v, &r.body)
		if err != nil {
			err = json.Unmarshal(b, &r.body)
			if err != nil {
				return fmt.Errorf("cannot unmarshal body: %w", err)
			}
		}
	}

	return nil
}

// ApiCustomPostRequest represents the request with all the parameters for the API call.
type ApiCustomPostRequest struct {
	path       string
	parameters map[string]any
	body       map[string]any
}

// NewApiCustomPostRequest creates an instance of the ApiCustomPostRequest to be used for the API call.
func (c *APIClient) NewApiCustomPostRequest(path string) ApiCustomPostRequest {
	return ApiCustomPostRequest{
		path: path,
	}
}

// WithParameters adds the parameters to the ApiCustomPostRequest and returns the request for chaining.
func (r ApiCustomPostRequest) WithParameters(parameters map[string]any) ApiCustomPostRequest {
	r.parameters = parameters
	return r
}

// WithBody adds the body to the ApiCustomPostRequest and returns the request for chaining.
func (r ApiCustomPostRequest) WithBody(body map[string]any) ApiCustomPostRequest {
	r.body = body
	return r
}

/*
CustomPost calls the API and returns the raw response from it.

	  This method lets you send requests to the Algolia REST API.


	Request can be constructed by NewApiCustomPostRequest with parameters below.
	  @param path string - Path of the endpoint, for example `1/newFeature`.
	  @param parameters map[string]any - Query parameters to apply to the current query.
	  @param body map[string]any - Parameters to send with the custom request.
	@param opts ...RequestOption - Optional parameters for the API call
	@return *http.Response - The raw response from the API
	@return []byte - The raw response body from the API
	@return error - An error if the API call fails
*/
func (c *APIClient) CustomPostWithHTTPInfo(r ApiCustomPostRequest, opts ...RequestOption) (*http.Response, []byte, error) {
	requestPath := "/{path}"
	requestPath = strings.ReplaceAll(requestPath, "{path}", utils.ParameterToString(r.path))

	if r.path == "" {
		return nil, nil, reportError("Parameter `path` is required when calling `CustomPost`.")
	}

	conf := config{
		context:      context.Background(),
		queryParams:  url.Values{},
		headerParams: map[string]string{},
	}

	if !utils.IsNilOrEmpty(r.parameters) {
		for k, v := range r.parameters {
			conf.queryParams.Set(k, utils.QueryParameterToString(v))
		}
	}

	// optional params if any
	for _, opt := range opts {
		opt.apply(&conf)
	}

	var postBody any

	// body params
	if utils.IsNilOrEmpty(r.body) {
		postBody = "{}"
	} else {
		postBody = r.body
	}
	req, err := c.prepareRequest(conf.context, requestPath, http.MethodPost, postBody, conf.headerParams, conf.queryParams)
	if err != nil {
		return nil, nil, err
	}

	return c.callAPI(req, false, conf.timeouts)
}

/*
CustomPost casts the HTTP response body to a defined struct.

This method lets you send requests to the Algolia REST API.

Request can be constructed by NewApiCustomPostRequest with parameters below.

	@param path string - Path of the endpoint, for example `1/newFeature`.
	@param parameters map[string]any - Query parameters to apply to the current query.
	@param body map[string]any - Parameters to send with the custom request.
	@return map[string]any
*/
func (c *APIClient) CustomPost(r ApiCustomPostRequest, opts ...RequestOption) (*map[string]any, error) {
	var returnValue *map[string]any

	res, resBody, err := c.CustomPostWithHTTPInfo(r, opts...)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	if res.StatusCode >= 300 {
		return returnValue, c.decodeError(res, resBody)
	}

	err = c.decode(&returnValue, resBody)
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

func (r *ApiCustomPutRequest) UnmarshalJSON(b []byte) error {
	req := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return fmt.Errorf("cannot unmarshal request: %w", err)
	}
	if v, ok := req["path"]; ok {
		err = json.Unmarshal(v, &r.path)
		if err != nil {
			err = json.Unmarshal(b, &r.path)
			if err != nil {
				return fmt.Errorf("cannot unmarshal path: %w", err)
			}
		}
	}
	if v, ok := req["parameters"]; ok {
		err = json.Unmarshal(v, &r.parameters)
		if err != nil {
			err = json.Unmarshal(b, &r.parameters)
			if err != nil {
				return fmt.Errorf("cannot unmarshal parameters: %w", err)
			}
		}
	}
	if v, ok := req["body"]; ok {
		err = json.Unmarshal(v, &r.body)
		if err != nil {
			err = json.Unmarshal(b, &r.body)
			if err != nil {
				return fmt.Errorf("cannot unmarshal body: %w", err)
			}
		}
	}

	return nil
}

// ApiCustomPutRequest represents the request with all the parameters for the API call.
type ApiCustomPutRequest struct {
	path       string
	parameters map[string]any
	body       map[string]any
}

// NewApiCustomPutRequest creates an instance of the ApiCustomPutRequest to be used for the API call.
func (c *APIClient) NewApiCustomPutRequest(path string) ApiCustomPutRequest {
	return ApiCustomPutRequest{
		path: path,
	}
}

// WithParameters adds the parameters to the ApiCustomPutRequest and returns the request for chaining.
func (r ApiCustomPutRequest) WithParameters(parameters map[string]any) ApiCustomPutRequest {
	r.parameters = parameters
	return r
}

// WithBody adds the body to the ApiCustomPutRequest and returns the request for chaining.
func (r ApiCustomPutRequest) WithBody(body map[string]any) ApiCustomPutRequest {
	r.body = body
	return r
}

/*
CustomPut calls the API and returns the raw response from it.

	  This method lets you send requests to the Algolia REST API.


	Request can be constructed by NewApiCustomPutRequest with parameters below.
	  @param path string - Path of the endpoint, for example `1/newFeature`.
	  @param parameters map[string]any - Query parameters to apply to the current query.
	  @param body map[string]any - Parameters to send with the custom request.
	@param opts ...RequestOption - Optional parameters for the API call
	@return *http.Response - The raw response from the API
	@return []byte - The raw response body from the API
	@return error - An error if the API call fails
*/
func (c *APIClient) CustomPutWithHTTPInfo(r ApiCustomPutRequest, opts ...RequestOption) (*http.Response, []byte, error) {
	requestPath := "/{path}"
	requestPath = strings.ReplaceAll(requestPath, "{path}", utils.ParameterToString(r.path))

	if r.path == "" {
		return nil, nil, reportError("Parameter `path` is required when calling `CustomPut`.")
	}

	conf := config{
		context:      context.Background(),
		queryParams:  url.Values{},
		headerParams: map[string]string{},
	}

	if !utils.IsNilOrEmpty(r.parameters) {
		for k, v := range r.parameters {
			conf.queryParams.Set(k, utils.QueryParameterToString(v))
		}
	}

	// optional params if any
	for _, opt := range opts {
		opt.apply(&conf)
	}

	var postBody any

	// body params
	if utils.IsNilOrEmpty(r.body) {
		postBody = "{}"
	} else {
		postBody = r.body
	}
	req, err := c.prepareRequest(conf.context, requestPath, http.MethodPut, postBody, conf.headerParams, conf.queryParams)
	if err != nil {
		return nil, nil, err
	}

	return c.callAPI(req, false, conf.timeouts)
}

/*
CustomPut casts the HTTP response body to a defined struct.

This method lets you send requests to the Algolia REST API.

Request can be constructed by NewApiCustomPutRequest with parameters below.

	@param path string - Path of the endpoint, for example `1/newFeature`.
	@param parameters map[string]any - Query parameters to apply to the current query.
	@param body map[string]any - Parameters to send with the custom request.
	@return map[string]any
*/
func (c *APIClient) CustomPut(r ApiCustomPutRequest, opts ...RequestOption) (*map[string]any, error) {
	var returnValue *map[string]any

	res, resBody, err := c.CustomPutWithHTTPInfo(r, opts...)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	if res.StatusCode >= 300 {
		return returnValue, c.decodeError(res, resBody)
	}

	err = c.decode(&returnValue, resBody)
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

func (r *ApiDeleteUserProfileRequest) UnmarshalJSON(b []byte) error {
	req := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return fmt.Errorf("cannot unmarshal request: %w", err)
	}
	if v, ok := req["userToken"]; ok {
		err = json.Unmarshal(v, &r.userToken)
		if err != nil {
			err = json.Unmarshal(b, &r.userToken)
			if err != nil {
				return fmt.Errorf("cannot unmarshal userToken: %w", err)
			}
		}
	}

	return nil
}

// ApiDeleteUserProfileRequest represents the request with all the parameters for the API call.
type ApiDeleteUserProfileRequest struct {
	userToken string
}

// NewApiDeleteUserProfileRequest creates an instance of the ApiDeleteUserProfileRequest to be used for the API call.
func (c *APIClient) NewApiDeleteUserProfileRequest(userToken string) ApiDeleteUserProfileRequest {
	return ApiDeleteUserProfileRequest{
		userToken: userToken,
	}
}

/*
DeleteUserProfile calls the API and returns the raw response from it.

	Deletes a user profile.

The response includes a date and time when the user profile can safely be considered deleted.

	    Required API Key ACLs:
	    - recommendation

	Request can be constructed by NewApiDeleteUserProfileRequest with parameters below.
	  @param userToken string - Unique identifier representing a user for which to fetch the personalization profile.
	@param opts ...RequestOption - Optional parameters for the API call
	@return *http.Response - The raw response from the API
	@return []byte - The raw response body from the API
	@return error - An error if the API call fails
*/
func (c *APIClient) DeleteUserProfileWithHTTPInfo(r ApiDeleteUserProfileRequest, opts ...RequestOption) (*http.Response, []byte, error) {
	requestPath := "/1/profiles/{userToken}"
	requestPath = strings.ReplaceAll(requestPath, "{userToken}", url.PathEscape(utils.ParameterToString(r.userToken)))

	if r.userToken == "" {
		return nil, nil, reportError("Parameter `userToken` is required when calling `DeleteUserProfile`.")
	}

	conf := config{
		context:      context.Background(),
		queryParams:  url.Values{},
		headerParams: map[string]string{},
	}

	// optional params if any
	for _, opt := range opts {
		opt.apply(&conf)
	}

	var postBody any

	req, err := c.prepareRequest(conf.context, requestPath, http.MethodDelete, postBody, conf.headerParams, conf.queryParams)
	if err != nil {
		return nil, nil, err
	}

	return c.callAPI(req, false, conf.timeouts)
}

/*
DeleteUserProfile casts the HTTP response body to a defined struct.

Deletes a user profile.

The response includes a date and time when the user profile can safely be considered deleted.

Required API Key ACLs:
  - recommendation

Request can be constructed by NewApiDeleteUserProfileRequest with parameters below.

	@param userToken string - Unique identifier representing a user for which to fetch the personalization profile.
	@return DeleteUserProfileResponse
*/
func (c *APIClient) DeleteUserProfile(r ApiDeleteUserProfileRequest, opts ...RequestOption) (*DeleteUserProfileResponse, error) {
	var returnValue *DeleteUserProfileResponse

	res, resBody, err := c.DeleteUserProfileWithHTTPInfo(r, opts...)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	if res.StatusCode >= 300 {
		return returnValue, c.decodeError(res, resBody)
	}

	err = c.decode(&returnValue, resBody)
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

/*
GetPersonalizationStrategy calls the API and returns the raw response from it.

	  Retrieves the current personalization strategy.

	    Required API Key ACLs:
	    - recommendation

	Request can be constructed by NewApiGetPersonalizationStrategyRequest with parameters below.
	@param opts ...RequestOption - Optional parameters for the API call
	@return *http.Response - The raw response from the API
	@return []byte - The raw response body from the API
	@return error - An error if the API call fails
*/
func (c *APIClient) GetPersonalizationStrategyWithHTTPInfo(opts ...RequestOption) (*http.Response, []byte, error) {
	requestPath := "/1/strategies/personalization"

	conf := config{
		context:      context.Background(),
		queryParams:  url.Values{},
		headerParams: map[string]string{},
	}

	// optional params if any
	for _, opt := range opts {
		opt.apply(&conf)
	}

	var postBody any

	req, err := c.prepareRequest(conf.context, requestPath, http.MethodGet, postBody, conf.headerParams, conf.queryParams)
	if err != nil {
		return nil, nil, err
	}

	return c.callAPI(req, false, conf.timeouts)
}

/*
GetPersonalizationStrategy casts the HTTP response body to a defined struct.

Retrieves the current personalization strategy.

Required API Key ACLs:
  - recommendation

Request can be constructed by NewApiGetPersonalizationStrategyRequest with parameters below.

	@return PersonalizationStrategyParams
*/
func (c *APIClient) GetPersonalizationStrategy(opts ...RequestOption) (*PersonalizationStrategyParams, error) {
	var returnValue *PersonalizationStrategyParams

	res, resBody, err := c.GetPersonalizationStrategyWithHTTPInfo(opts...)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	if res.StatusCode >= 300 {
		return returnValue, c.decodeError(res, resBody)
	}

	err = c.decode(&returnValue, resBody)
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

func (r *ApiGetUserTokenProfileRequest) UnmarshalJSON(b []byte) error {
	req := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return fmt.Errorf("cannot unmarshal request: %w", err)
	}
	if v, ok := req["userToken"]; ok {
		err = json.Unmarshal(v, &r.userToken)
		if err != nil {
			err = json.Unmarshal(b, &r.userToken)
			if err != nil {
				return fmt.Errorf("cannot unmarshal userToken: %w", err)
			}
		}
	}

	return nil
}

// ApiGetUserTokenProfileRequest represents the request with all the parameters for the API call.
type ApiGetUserTokenProfileRequest struct {
	userToken string
}

// NewApiGetUserTokenProfileRequest creates an instance of the ApiGetUserTokenProfileRequest to be used for the API call.
func (c *APIClient) NewApiGetUserTokenProfileRequest(userToken string) ApiGetUserTokenProfileRequest {
	return ApiGetUserTokenProfileRequest{
		userToken: userToken,
	}
}

/*
GetUserTokenProfile calls the API and returns the raw response from it.

	  Retrieves a user profile and their affinities for different facets.

	    Required API Key ACLs:
	    - recommendation

	Request can be constructed by NewApiGetUserTokenProfileRequest with parameters below.
	  @param userToken string - Unique identifier representing a user for which to fetch the personalization profile.
	@param opts ...RequestOption - Optional parameters for the API call
	@return *http.Response - The raw response from the API
	@return []byte - The raw response body from the API
	@return error - An error if the API call fails
*/
func (c *APIClient) GetUserTokenProfileWithHTTPInfo(r ApiGetUserTokenProfileRequest, opts ...RequestOption) (*http.Response, []byte, error) {
	requestPath := "/1/profiles/personalization/{userToken}"
	requestPath = strings.ReplaceAll(requestPath, "{userToken}", url.PathEscape(utils.ParameterToString(r.userToken)))

	if r.userToken == "" {
		return nil, nil, reportError("Parameter `userToken` is required when calling `GetUserTokenProfile`.")
	}

	conf := config{
		context:      context.Background(),
		queryParams:  url.Values{},
		headerParams: map[string]string{},
	}

	// optional params if any
	for _, opt := range opts {
		opt.apply(&conf)
	}

	var postBody any

	req, err := c.prepareRequest(conf.context, requestPath, http.MethodGet, postBody, conf.headerParams, conf.queryParams)
	if err != nil {
		return nil, nil, err
	}

	return c.callAPI(req, false, conf.timeouts)
}

/*
GetUserTokenProfile casts the HTTP response body to a defined struct.

Retrieves a user profile and their affinities for different facets.

Required API Key ACLs:
  - recommendation

Request can be constructed by NewApiGetUserTokenProfileRequest with parameters below.

	@param userToken string - Unique identifier representing a user for which to fetch the personalization profile.
	@return GetUserTokenResponse
*/
func (c *APIClient) GetUserTokenProfile(r ApiGetUserTokenProfileRequest, opts ...RequestOption) (*GetUserTokenResponse, error) {
	var returnValue *GetUserTokenResponse

	res, resBody, err := c.GetUserTokenProfileWithHTTPInfo(r, opts...)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	if res.StatusCode >= 300 {
		return returnValue, c.decodeError(res, resBody)
	}

	err = c.decode(&returnValue, resBody)
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

func (r *ApiSetPersonalizationStrategyRequest) UnmarshalJSON(b []byte) error {
	req := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return fmt.Errorf("cannot unmarshal request: %w", err)
	}
	if v, ok := req["personalizationStrategyParams"]; ok {
		err = json.Unmarshal(v, &r.personalizationStrategyParams)
		if err != nil {
			err = json.Unmarshal(b, &r.personalizationStrategyParams)
			if err != nil {
				return fmt.Errorf("cannot unmarshal personalizationStrategyParams: %w", err)
			}
		}
	} else {
		err = json.Unmarshal(b, &r.personalizationStrategyParams)
		if err != nil {
			return fmt.Errorf("cannot unmarshal body parameter personalizationStrategyParams: %w", err)
		}
	}

	return nil
}

// ApiSetPersonalizationStrategyRequest represents the request with all the parameters for the API call.
type ApiSetPersonalizationStrategyRequest struct {
	personalizationStrategyParams *PersonalizationStrategyParams
}

// NewApiSetPersonalizationStrategyRequest creates an instance of the ApiSetPersonalizationStrategyRequest to be used for the API call.
func (c *APIClient) NewApiSetPersonalizationStrategyRequest(personalizationStrategyParams *PersonalizationStrategyParams) ApiSetPersonalizationStrategyRequest {
	return ApiSetPersonalizationStrategyRequest{
		personalizationStrategyParams: personalizationStrategyParams,
	}
}

/*
SetPersonalizationStrategy calls the API and returns the raw response from it.

	  Creates a new personalization strategy.

	    Required API Key ACLs:
	    - recommendation

	Request can be constructed by NewApiSetPersonalizationStrategyRequest with parameters below.
	  @param personalizationStrategyParams PersonalizationStrategyParams
	@param opts ...RequestOption - Optional parameters for the API call
	@return *http.Response - The raw response from the API
	@return []byte - The raw response body from the API
	@return error - An error if the API call fails
*/
func (c *APIClient) SetPersonalizationStrategyWithHTTPInfo(r ApiSetPersonalizationStrategyRequest, opts ...RequestOption) (*http.Response, []byte, error) {
	requestPath := "/1/strategies/personalization"

	if r.personalizationStrategyParams == nil {
		return nil, nil, reportError("Parameter `personalizationStrategyParams` is required when calling `SetPersonalizationStrategy`.")
	}

	conf := config{
		context:      context.Background(),
		queryParams:  url.Values{},
		headerParams: map[string]string{},
	}

	// optional params if any
	for _, opt := range opts {
		opt.apply(&conf)
	}

	var postBody any

	// body params
	postBody = r.personalizationStrategyParams
	req, err := c.prepareRequest(conf.context, requestPath, http.MethodPost, postBody, conf.headerParams, conf.queryParams)
	if err != nil {
		return nil, nil, err
	}

	return c.callAPI(req, false, conf.timeouts)
}

/*
SetPersonalizationStrategy casts the HTTP response body to a defined struct.

Creates a new personalization strategy.

Required API Key ACLs:
  - recommendation

Request can be constructed by NewApiSetPersonalizationStrategyRequest with parameters below.

	@param personalizationStrategyParams PersonalizationStrategyParams
	@return SetPersonalizationStrategyResponse
*/
func (c *APIClient) SetPersonalizationStrategy(r ApiSetPersonalizationStrategyRequest, opts ...RequestOption) (*SetPersonalizationStrategyResponse, error) {
	var returnValue *SetPersonalizationStrategyResponse

	res, resBody, err := c.SetPersonalizationStrategyWithHTTPInfo(r, opts...)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	if res.StatusCode >= 300 {
		return returnValue, c.decodeError(res, resBody)
	}

	err = c.decode(&returnValue, resBody)
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}
