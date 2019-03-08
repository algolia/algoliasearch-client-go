package search

type Synonym interface {
	ObjectID() string
	Type() SynonymType
}

type SynonymType string

const (
	RegularSynonymType SynonymType = "synonym"
	OneWaySynonmType   SynonymType = "oneWaySynonym"
	AltCorrection1Type SynonymType = "altCorrection1"
	AltCorrection2Type SynonymType = "altCorrection2"
	PlaceholderType    SynonymType = "placeholder"
)
