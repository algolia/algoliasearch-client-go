package opt

import "encoding/json"

type LengthOption struct {
	nbRecords int
}

func Length(nbRecords int) LengthOption {
	return LengthOption{nbRecords}
}

func (o LengthOption) Get() int {
	return o.nbRecords
}

func (o LengthOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.nbRecords)
}

func (o *LengthOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.nbRecords)
}
