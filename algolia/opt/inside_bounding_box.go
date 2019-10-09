package opt

import (
	"encoding/json"
	"reflect"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/errs"
)

// InsideBoundingBoxOption is a wrapper for an InsideBoundingBox option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type InsideBoundingBoxOption struct {
	boxes       [][4]float64
	coordinates string
}

// InsideBoundingBox returns an InsideBoundingBoxOption whose value is set to the given
// slice of rectangle coordinates.
func InsideBoundingBox(boxes [][4]float64) *InsideBoundingBoxOption {
	return &InsideBoundingBoxOption{boxes: boxes}
}

// InsideBoundingBoxFromCoordinates returns an InsideBoundingBoxOption whose value is to
// the given coordinates string.
func InsideBoundingBoxFromCoordinates(coordinates string) *InsideBoundingBoxOption {
	return &InsideBoundingBoxOption{coordinates: coordinates}
}

// Get retrieves the actual value of the option parameter.
func (o *InsideBoundingBoxOption) Get() ([][4]float64, string) {
	if o == nil {
		return nil, ""
	}
	return o.boxes, o.coordinates
}

// MarshalJSON implements the json.Marshaler interface for
// InsideBoundingBoxOption.
func (o InsideBoundingBoxOption) MarshalJSON() ([]byte, error) {
	if o.coordinates != "" {
		return json.Marshal(o.coordinates)
	}
	return json.Marshal(o.boxes)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// InsideBoundingBoxOption.
func (o *InsideBoundingBoxOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	var boxes [][4]float64
	if err := json.Unmarshal(data, &boxes); err == nil {
		o.boxes = boxes
		return nil
	}

	var coordinates string
	if err := json.Unmarshal(data, &coordinates); err == nil {
		o.coordinates = coordinates
		return nil
	}

	return errs.ErrJSONDecode(data, "InsideBoundingBox")
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *InsideBoundingBoxOption) Equal(o2 *InsideBoundingBoxOption) bool {
	if o == nil {
		return o2 == nil || len(o2.boxes) == 0 && len(o2.coordinates) == 0
	}
	if o2 == nil {
		return o == nil || len(o.boxes) == 0 && len(o.coordinates) == 0
	}
	return reflect.DeepEqual(o, o2)
}

// InsideBoundingBoxEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func InsideBoundingBoxEqual(o1, o2 *InsideBoundingBoxOption) bool {
	return o1.Equal(o2)
}
