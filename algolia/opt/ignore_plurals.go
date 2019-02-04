package opt

import "encoding/json"

type IgnorePluralsOption struct {
	ignorePlurals bool
	languages     []string
}

func IgnorePlurals(languages ...string) IgnorePluralsOption {
	return IgnorePluralsOption{languages: languages}
}

func IgnorePluralsTrue() IgnorePluralsOption {
	return IgnorePluralsOption{ignorePlurals: true}
}

func IgnorePluralsFalse() IgnorePluralsOption {
	return IgnorePluralsOption{ignorePlurals: false}
}

func (o IgnorePluralsOption) Get() (bool, []string) {
	return o.ignorePlurals, o.languages
}

func (o IgnorePluralsOption) MarshalJSON() ([]byte, error) {
	if len(o.languages) > 0 {
		return json.Marshal(o.languages)
	}
	return json.Marshal(o.ignorePlurals)
}

func (o *IgnorePluralsOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	if err := json.Unmarshal(data, &o.languages); err == nil {
		return nil
	}

	return json.Unmarshal(data, &o.ignorePlurals)
}
