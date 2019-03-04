package opt

import (
	"encoding/json"

	"github.com/algolia/algoliasearch-client-go/algolia/errs"
)

type InsidePolygonOption struct {
	polygons    [][]float64
	coordinates string
}

func InsidePolygon(polygons [][]float64) *InsidePolygonOption {
	return &InsidePolygonOption{polygons: polygons}
}

func InsidePolygonFromCoordinates(coordinates string) *InsidePolygonOption {
	return &InsidePolygonOption{coordinates: coordinates}
}

func (o InsidePolygonOption) Get() ([][]float64, string) {
	return o.polygons, o.coordinates
}

func (o InsidePolygonOption) MarshalJSON() ([]byte, error) {
	if o.coordinates != "" {
		return json.Marshal(o.coordinates)
	}
	return json.Marshal(o.polygons)
}

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
