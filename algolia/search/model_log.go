// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// Log struct for Log.
type Log struct {
	// Date and time of the API request, in RFC 3339 format.
	Timestamp string `json:"timestamp"`
	// HTTP method of the request.
	Method string `json:"method"`
	// HTTP status code of the response.
	AnswerCode string `json:"answer_code"`
	// Request body.
	QueryBody string `json:"query_body"`
	// Response body.
	Answer string `json:"answer"`
	// URL of the API endpoint.
	Url string `json:"url"`
	// IP address of the client that performed the request.
	Ip string `json:"ip"`
	// Request headers (API keys are obfuscated).
	QueryHeaders string `json:"query_headers"`
	// SHA1 signature of the log entry.
	Sha1 string `json:"sha1"`
	// Number of API requests.
	NbApiCalls string `json:"nb_api_calls"`
	// Processing time for the query in milliseconds. This doesn't include latency due to the network.
	ProcessingTimeMs string `json:"processing_time_ms"`
	// Index targeted by the query.
	Index *string `json:"index,omitempty"`
	// Query parameters sent with the request.
	QueryParams *string `json:"query_params,omitempty"`
	// Number of search results (hits) returned for the query.
	QueryNbHits *string `json:"query_nb_hits,omitempty"`
	// Queries performed for the given request.
	InnerQueries []LogQuery `json:"inner_queries,omitempty"`
}

type LogOption func(f *Log)

func WithLogIndex(val string) LogOption {
	return func(f *Log) {
		f.Index = &val
	}
}

func WithLogQueryParams(val string) LogOption {
	return func(f *Log) {
		f.QueryParams = &val
	}
}

func WithLogQueryNbHits(val string) LogOption {
	return func(f *Log) {
		f.QueryNbHits = &val
	}
}

func WithLogInnerQueries(val []LogQuery) LogOption {
	return func(f *Log) {
		f.InnerQueries = val
	}
}

// NewLog instantiates a new Log object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLog(timestamp string, method string, answerCode string, queryBody string, answer string, url string, ip string, queryHeaders string, sha1 string, nbApiCalls string, processingTimeMs string, opts ...LogOption) *Log {
	this := &Log{}
	this.Timestamp = timestamp
	this.Method = method
	this.AnswerCode = answerCode
	this.QueryBody = queryBody
	this.Answer = answer
	this.Url = url
	this.Ip = ip
	this.QueryHeaders = queryHeaders
	this.Sha1 = sha1
	this.NbApiCalls = nbApiCalls
	this.ProcessingTimeMs = processingTimeMs
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyLog return a pointer to an empty Log object.
func NewEmptyLog() *Log {
	return &Log{}
}

// GetTimestamp returns the Timestamp field value.
func (o *Log) GetTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *Log) GetTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value.
func (o *Log) SetTimestamp(v string) *Log {
	o.Timestamp = v
	return o
}

// GetMethod returns the Method field value.
func (o *Log) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *Log) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value.
func (o *Log) SetMethod(v string) *Log {
	o.Method = v
	return o
}

// GetAnswerCode returns the AnswerCode field value.
func (o *Log) GetAnswerCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AnswerCode
}

// GetAnswerCodeOk returns a tuple with the AnswerCode field value
// and a boolean to check if the value has been set.
func (o *Log) GetAnswerCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnswerCode, true
}

// SetAnswerCode sets field value.
func (o *Log) SetAnswerCode(v string) *Log {
	o.AnswerCode = v
	return o
}

// GetQueryBody returns the QueryBody field value.
func (o *Log) GetQueryBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryBody
}

// GetQueryBodyOk returns a tuple with the QueryBody field value
// and a boolean to check if the value has been set.
func (o *Log) GetQueryBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryBody, true
}

// SetQueryBody sets field value.
func (o *Log) SetQueryBody(v string) *Log {
	o.QueryBody = v
	return o
}

// GetAnswer returns the Answer field value.
func (o *Log) GetAnswer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Answer
}

// GetAnswerOk returns a tuple with the Answer field value
// and a boolean to check if the value has been set.
func (o *Log) GetAnswerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Answer, true
}

// SetAnswer sets field value.
func (o *Log) SetAnswer(v string) *Log {
	o.Answer = v
	return o
}

// GetUrl returns the Url field value.
func (o *Log) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Log) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value.
func (o *Log) SetUrl(v string) *Log {
	o.Url = v
	return o
}

// GetIp returns the Ip field value.
func (o *Log) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *Log) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value.
func (o *Log) SetIp(v string) *Log {
	o.Ip = v
	return o
}

// GetQueryHeaders returns the QueryHeaders field value.
func (o *Log) GetQueryHeaders() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryHeaders
}

// GetQueryHeadersOk returns a tuple with the QueryHeaders field value
// and a boolean to check if the value has been set.
func (o *Log) GetQueryHeadersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryHeaders, true
}

// SetQueryHeaders sets field value.
func (o *Log) SetQueryHeaders(v string) *Log {
	o.QueryHeaders = v
	return o
}

// GetSha1 returns the Sha1 field value.
func (o *Log) GetSha1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sha1
}

// GetSha1Ok returns a tuple with the Sha1 field value
// and a boolean to check if the value has been set.
func (o *Log) GetSha1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sha1, true
}

// SetSha1 sets field value.
func (o *Log) SetSha1(v string) *Log {
	o.Sha1 = v
	return o
}

// GetNbApiCalls returns the NbApiCalls field value.
func (o *Log) GetNbApiCalls() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NbApiCalls
}

// GetNbApiCallsOk returns a tuple with the NbApiCalls field value
// and a boolean to check if the value has been set.
func (o *Log) GetNbApiCallsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NbApiCalls, true
}

// SetNbApiCalls sets field value.
func (o *Log) SetNbApiCalls(v string) *Log {
	o.NbApiCalls = v
	return o
}

// GetProcessingTimeMs returns the ProcessingTimeMs field value.
func (o *Log) GetProcessingTimeMs() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessingTimeMs
}

// GetProcessingTimeMsOk returns a tuple with the ProcessingTimeMs field value
// and a boolean to check if the value has been set.
func (o *Log) GetProcessingTimeMsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessingTimeMs, true
}

// SetProcessingTimeMs sets field value.
func (o *Log) SetProcessingTimeMs(v string) *Log {
	o.ProcessingTimeMs = v
	return o
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *Log) GetIndex() string {
	if o == nil || o.Index == nil {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetIndexOk() (*string, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *Log) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *Log) SetIndex(v string) *Log {
	o.Index = &v
	return o
}

// GetQueryParams returns the QueryParams field value if set, zero value otherwise.
func (o *Log) GetQueryParams() string {
	if o == nil || o.QueryParams == nil {
		var ret string
		return ret
	}
	return *o.QueryParams
}

// GetQueryParamsOk returns a tuple with the QueryParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetQueryParamsOk() (*string, bool) {
	if o == nil || o.QueryParams == nil {
		return nil, false
	}
	return o.QueryParams, true
}

// HasQueryParams returns a boolean if a field has been set.
func (o *Log) HasQueryParams() bool {
	if o != nil && o.QueryParams != nil {
		return true
	}

	return false
}

// SetQueryParams gets a reference to the given string and assigns it to the QueryParams field.
func (o *Log) SetQueryParams(v string) *Log {
	o.QueryParams = &v
	return o
}

// GetQueryNbHits returns the QueryNbHits field value if set, zero value otherwise.
func (o *Log) GetQueryNbHits() string {
	if o == nil || o.QueryNbHits == nil {
		var ret string
		return ret
	}
	return *o.QueryNbHits
}

// GetQueryNbHitsOk returns a tuple with the QueryNbHits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetQueryNbHitsOk() (*string, bool) {
	if o == nil || o.QueryNbHits == nil {
		return nil, false
	}
	return o.QueryNbHits, true
}

// HasQueryNbHits returns a boolean if a field has been set.
func (o *Log) HasQueryNbHits() bool {
	if o != nil && o.QueryNbHits != nil {
		return true
	}

	return false
}

// SetQueryNbHits gets a reference to the given string and assigns it to the QueryNbHits field.
func (o *Log) SetQueryNbHits(v string) *Log {
	o.QueryNbHits = &v
	return o
}

// GetInnerQueries returns the InnerQueries field value if set, zero value otherwise.
func (o *Log) GetInnerQueries() []LogQuery {
	if o == nil || o.InnerQueries == nil {
		var ret []LogQuery
		return ret
	}
	return o.InnerQueries
}

// GetInnerQueriesOk returns a tuple with the InnerQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetInnerQueriesOk() ([]LogQuery, bool) {
	if o == nil || o.InnerQueries == nil {
		return nil, false
	}
	return o.InnerQueries, true
}

// HasInnerQueries returns a boolean if a field has been set.
func (o *Log) HasInnerQueries() bool {
	if o != nil && o.InnerQueries != nil {
		return true
	}

	return false
}

// SetInnerQueries gets a reference to the given []LogQuery and assigns it to the InnerQueries field.
func (o *Log) SetInnerQueries(v []LogQuery) *Log {
	o.InnerQueries = v
	return o
}

func (o Log) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["answer_code"] = o.AnswerCode
	}
	if true {
		toSerialize["query_body"] = o.QueryBody
	}
	if true {
		toSerialize["answer"] = o.Answer
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["ip"] = o.Ip
	}
	if true {
		toSerialize["query_headers"] = o.QueryHeaders
	}
	if true {
		toSerialize["sha1"] = o.Sha1
	}
	if true {
		toSerialize["nb_api_calls"] = o.NbApiCalls
	}
	if true {
		toSerialize["processing_time_ms"] = o.ProcessingTimeMs
	}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.QueryParams != nil {
		toSerialize["query_params"] = o.QueryParams
	}
	if o.QueryNbHits != nil {
		toSerialize["query_nb_hits"] = o.QueryNbHits
	}
	if o.InnerQueries != nil {
		toSerialize["inner_queries"] = o.InnerQueries
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal Log: %w", err)
	}

	return serialized, nil
}

func (o Log) String() string {
	out := ""
	out += fmt.Sprintf("  timestamp=%v\n", o.Timestamp)
	out += fmt.Sprintf("  method=%v\n", o.Method)
	out += fmt.Sprintf("  answer_code=%v\n", o.AnswerCode)
	out += fmt.Sprintf("  query_body=%v\n", o.QueryBody)
	out += fmt.Sprintf("  answer=%v\n", o.Answer)
	out += fmt.Sprintf("  url=%v\n", o.Url)
	out += fmt.Sprintf("  ip=%v\n", o.Ip)
	out += fmt.Sprintf("  query_headers=%v\n", o.QueryHeaders)
	out += fmt.Sprintf("  sha1=%v\n", o.Sha1)
	out += fmt.Sprintf("  nb_api_calls=%v\n", o.NbApiCalls)
	out += fmt.Sprintf("  processing_time_ms=%v\n", o.ProcessingTimeMs)
	out += fmt.Sprintf("  index=%v\n", o.Index)
	out += fmt.Sprintf("  query_params=%v\n", o.QueryParams)
	out += fmt.Sprintf("  query_nb_hits=%v\n", o.QueryNbHits)
	out += fmt.Sprintf("  inner_queries=%v\n", o.InnerQueries)
	return fmt.Sprintf("Log {\n%s}", out)
}

type NullableLog struct {
	value *Log
	isSet bool
}

func (v NullableLog) Get() *Log {
	return v.value
}

func (v *NullableLog) Set(val *Log) {
	v.value = val
	v.isSet = true
}

func (v NullableLog) IsSet() bool {
	return v.isSet
}

func (v *NullableLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLog(val *Log) *NullableLog {
	return &NullableLog{value: val, isSet: true}
}

func (v NullableLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
