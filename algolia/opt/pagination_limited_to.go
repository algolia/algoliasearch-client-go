package opt

import "encoding/json"

type PaginationLimitedToOption struct {
	nbRecords int
}

func PaginationLimitedTo(nbRecords int) PaginationLimitedToOption {
	return PaginationLimitedToOption{nbRecords}
}

func (o PaginationLimitedToOption) Get() int {
	return o.nbRecords
}

func (o PaginationLimitedToOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.nbRecords)
}

func (o *PaginationLimitedToOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.nbRecords)
}
