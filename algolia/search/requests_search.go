package search

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	iopt "github.com/algolia/algoliasearch-client-go/v3/algolia/internal/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
)

type searchReq struct {
	Params string `json:"params"`
}

type searchParams struct {
	Query string `json:"query"`
	QueryParams
	ExtraParams map[string]interface{} `json:"-"`
}

func (p searchParams) MarshalJSON() ([]byte, error) {

	// First, if there is no extra parameter, the structure is
	// serialized as-is, using a type alias to avoid infinite
	// recursion of MarshalJSON calls.

	if len(p.ExtraParams) == 0 {
		type searchParamsAlias searchParams
		return json.Marshal(searchParamsAlias(p))
	}

	// By there, because we do have extra search parameters, it
	// is needed to flatten them in the same JSON payload, along
	// with the known ones.
	//
	// To do that, we create a new type at runtime using the reflect
	// package and populate it with:
	//
	//   1. the known fields: "query" + any other fields from
	//      QueryParams if they are present
	//   2. the unknown fields from the ExtraParams map
	//
	// Once the new type has been created, we will populate it with the
	// actual values from the initial searchParams.

	var fields []reflect.StructField

	t := reflect.TypeOf(p)
	for i := 0; i < t.NumField(); i++ {
		fields = append(fields, t.Field(i))
	}

	for k, v := range p.ExtraParams {
		fields = append(fields, reflect.StructField{
			Name: strings.ToTitle(k),
			Type: reflect.TypeOf(v),
			Tag:  reflect.StructTag(fmt.Sprintf(`json:"%s"`, k)),
		})
	}

	e := reflect.New(reflect.StructOf(fields)).Elem()
	e.FieldByName("Query").Set(reflect.ValueOf(p.Query))
	e.FieldByName("QueryParams").Set(reflect.ValueOf(p.QueryParams))
	for k, v := range p.ExtraParams {
		e.FieldByName(strings.ToTitle(k)).Set(reflect.ValueOf(v))
	}

	return json.Marshal(e.Interface())
}

func newSearchParams(query string, opts ...interface{}) searchParams {
	return searchParams{
		Query:       query,
		QueryParams: newQueryParams(opts...),
		ExtraParams: iopt.ExtractExtraOptions(opts...).Get(),
	}
}

type searchForFacetValuesParams struct {
	FacetQuery string           `json:"facetQuery"`
	Query      *opt.QueryOption `json:"query"`
	QueryParams
}

func newSearchForFacetValuesParams(query string, opts ...interface{}) searchForFacetValuesParams {
	return searchForFacetValuesParams{
		FacetQuery:  query,
		Query:       iopt.ExtractQuery(opts...),
		QueryParams: newQueryParams(opts...),
	}
}
