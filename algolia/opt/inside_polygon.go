package opt

import (
	"encoding/json"
	"reflect"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/errs"
)

// InsidePolygonOption is a wrapper for an InsidePolygon option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type InsidePolygonOption struct {
	polygons    [][]float64
	coordinates string
}

// InsidePolygon returns an InsidePolygonOption whose value is set to the given
// slice of rectangle coordinates.
func InsidePolygon(polygons [][]float64) *InsidePolygonOption {
	return &InsidePolygonOption{polygons: polygons}
}

// InsidePolygonFromCoordinates returns an InsidePolygonOption whose value is to
// the given coordinates string.
func InsidePolygonFromCoordinates(coordinates string) *InsidePolygonOption {
	return &InsidePolygonOption{coordinates: coordinates}
}

// Get retrieves the actual value of the option parameter.
func (o *InsidePolygonOption) Get() ([][]float64, string) {
	if o == nil {
		return nil, ""
	}
	return o.polygons, o.coordinates
}

// MarshalJSON implements the json.Marshaler interface for
// InsidePolygonOption.
func (o InsidePolygonOption) MarshalJSON() ([]byte, error) {
	if o.coordinates != "" {
		return json.Marshal(o.coordinates)
	}
	return json.Marshal(o.polygons)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// InsidePolygonOption.
func (o *InsidePolygonOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	var polygons [][]float64
	if err := json.Unmarshal(data, &polygons); err == nil {
		o.polygons = polygons
		return nil
	}

	var coordinates string
	if err := json.Unmarshal(data, &coordinates); err == nil {
		o.coordinates = coordinates
		return nil
	}

	return errs.ErrJSONDecode(data, "InsidePolygon")
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *InsidePolygonOption) Equal(o2 *InsidePolygonOption) bool {
	if o == nil {
		return o2 == nil || len(o.polygons) == 0 && len(o.coordinates) == 0
	}
	if o2 == nil {
		return o == nil || len(o.polygons) == 0 && len(o.coordinates) == 0
	}
	return reflect.DeepEqual(o, o2)
}

// InsidePolygonEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func InsidePolygonEqual(o1, o2 *InsidePolygonOption) bool {
	return o1.Equal(o2)
}
