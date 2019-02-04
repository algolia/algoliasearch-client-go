package opt

import "encoding/json"

type MinWordSizeFor1TypoOption struct {
	minWordSize int
}

func MinWordSizeFor1Typo(minWordSize int) MinWordSizeFor1TypoOption {
	return MinWordSizeFor1TypoOption{minWordSize}
}

func (o MinWordSizeFor1TypoOption) Get() int {
	return o.minWordSize
}

func (o MinWordSizeFor1TypoOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.minWordSize)
}

func (o *MinWordSizeFor1TypoOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.minWordSize)
}
