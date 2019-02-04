package opt

import "encoding/json"

type OffsetOption struct {
	nbRecords int
}

func Offset(nbRecords int) OffsetOption {
	return OffsetOption{nbRecords}
}

func (o OffsetOption) Get() int {
	return o.nbRecords
}

func (o OffsetOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.nbRecords)
}

func (o *OffsetOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.nbRecords)
}
