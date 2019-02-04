package opt

import "encoding/json"

type MinWordSizeFor2TyposOption struct {
	minWordSize int
}

func MinWordSizeFor2Typos(minWordSize int) MinWordSizeFor2TyposOption {
	return MinWordSizeFor2TyposOption{minWordSize}
}

func (o MinWordSizeFor2TyposOption) Get() int {
	return o.minWordSize
}

func (o MinWordSizeFor2TyposOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.minWordSize)
}

func (o *MinWordSizeFor2TyposOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.minWordSize)
}
