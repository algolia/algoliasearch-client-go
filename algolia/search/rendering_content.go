package search

// Content defining how the search interface should be rendered.
// This is set via the settings for a default value and can be overridden via rules
type RenderingContent struct {
	FacetOrdering *FacetOrdering `json:"facetOrdering,omitempty"`
	Redirect      *Redirect      `json:"redirect,omitempty"`
}

// Facets and facets values ordering rules container
type FacetOrdering struct {
	// The ordering of facets.
	Facets *FacetsOrder `json:"facets"`

	// The ordering of facet values, within an individual list.
	Values map[string]FacetValuesOrder `json:"values,omitempty"`
}

// Redirect rule container
type Redirect struct {
	Url string `json:"url"`
}

// Facets ordering rule container
type FacetsOrder struct {
	// Pinned order of facet lists.
	Order []string `json:"order"`
}

// Facet values ordering rule container
type FacetValuesOrder struct {
	// Pinned order of facet values.
	Order []string `json:"order,omitempty"`

	// How to display the remaining items.
	SortRemainingBy *SortRule `json:"sortRemainingBy"`
}

func NewFacetValuesOrder(order []string, sortRemainingBy SortRule) FacetValuesOrder {
	return FacetValuesOrder{Order: order, SortRemainingBy: &sortRemainingBy}
}

// Rule defining the sort order of facet values.
type SortRule string

const (
	// alphabetical (ascending)
	Alpha SortRule = "alpha"

	// facet count (descending)
	Count SortRule = "count"

	// hidden (show only pinned values)
	Hidden SortRule = "hidden"
)
