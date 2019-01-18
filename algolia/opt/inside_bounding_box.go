package opt

import (
	"encoding/json"

	"github.com/algolia/algoliasearch-client-go/algolia/errs"
)

type InsideBoundingBoxOption struct {
	boxes       [][4]float64
	coordinates string
}

func InsideBoundingBox(boxes [][4]float64) InsideBoundingBoxOption {
	return InsideBoundingBoxOption{boxes: boxes}
}

func InsideBoundingBoxFromCoordinates(coordinates string) InsideBoundingBoxOption {
	return InsideBoundingBoxOption{coordinates: coordinates}
}

func (o InsideBoundingBoxOption) Get() ([][4]float64, string) {
	return o.boxes, o.coordinates
}

func (o InsideBoundingBoxOption) MarshalJSON() ([]byte, error) {
	if o.coordinates != "" {
		return json.Marshal(o.coordinates)
	}
	return json.Marshal(o.boxes)
}

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
