package opt

import "encoding/json"

type AroundLatLngOption struct {
	latLng string
}

func AroundLatLng(latLng string) AroundLatLngOption {
	return AroundLatLngOption{latLng}
}

func (o AroundLatLngOption) Get() string {
	return o.latLng
}

func (o AroundLatLngOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.latLng)
}

func (o *AroundLatLngOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.latLng)
}
