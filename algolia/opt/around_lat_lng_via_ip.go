package opt

import "encoding/json"

type AroundLatLngViaIPOption struct {
	value bool
}

func AroundLatLngViaIP(v bool) AroundLatLngViaIPOption {
	return AroundLatLngViaIPOption{v}
}

func (o AroundLatLngViaIPOption) Get() bool {
	return o.value
}

func (o AroundLatLngViaIPOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *AroundLatLngViaIPOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.value)
}
