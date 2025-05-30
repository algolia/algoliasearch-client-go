// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// SynonymHit Synonym object.
type SynonymHit struct {
	// Unique identifier of a synonym object.
	ObjectID string      `json:"objectID"`
	Type     SynonymType `json:"type"`
	// Words or phrases considered equivalent.
	Synonyms []string `json:"synonyms,omitempty"`
	// Word or phrase to appear in query strings (for [`onewaysynonym`s](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/adding-synonyms/in-depth/one-way-synonyms/)).
	Input *string `json:"input,omitempty"`
	// Word or phrase to appear in query strings (for [`altcorrection1` and `altcorrection2`](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/adding-synonyms/in-depth/synonyms-alternative-corrections/)).
	Word *string `json:"word,omitempty"`
	// Words to be matched in records.
	Corrections []string `json:"corrections,omitempty"`
	// [Placeholder token](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/adding-synonyms/in-depth/synonyms-placeholders/) to be put inside records.
	Placeholder *string `json:"placeholder,omitempty"`
	// Query words that will match the [placeholder token](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/adding-synonyms/in-depth/synonyms-placeholders/).
	Replacements []string `json:"replacements,omitempty"`
}

type SynonymHitOption func(f *SynonymHit)

func WithSynonymHitSynonyms(val []string) SynonymHitOption {
	return func(f *SynonymHit) {
		f.Synonyms = val
	}
}

func WithSynonymHitInput(val string) SynonymHitOption {
	return func(f *SynonymHit) {
		f.Input = &val
	}
}

func WithSynonymHitWord(val string) SynonymHitOption {
	return func(f *SynonymHit) {
		f.Word = &val
	}
}

func WithSynonymHitCorrections(val []string) SynonymHitOption {
	return func(f *SynonymHit) {
		f.Corrections = val
	}
}

func WithSynonymHitPlaceholder(val string) SynonymHitOption {
	return func(f *SynonymHit) {
		f.Placeholder = &val
	}
}

func WithSynonymHitReplacements(val []string) SynonymHitOption {
	return func(f *SynonymHit) {
		f.Replacements = val
	}
}

// NewSynonymHit instantiates a new SynonymHit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSynonymHit(objectID string, type_ SynonymType, opts ...SynonymHitOption) *SynonymHit {
	this := &SynonymHit{}
	this.ObjectID = objectID
	this.Type = type_
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptySynonymHit return a pointer to an empty SynonymHit object.
func NewEmptySynonymHit() *SynonymHit {
	return &SynonymHit{}
}

// GetObjectID returns the ObjectID field value.
func (o *SynonymHit) GetObjectID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectID
}

// GetObjectIDOk returns a tuple with the ObjectID field value
// and a boolean to check if the value has been set.
func (o *SynonymHit) GetObjectIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectID, true
}

// SetObjectID sets field value.
func (o *SynonymHit) SetObjectID(v string) *SynonymHit {
	o.ObjectID = v
	return o
}

// GetType returns the Type field value.
func (o *SynonymHit) GetType() SynonymType {
	if o == nil {
		var ret SynonymType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SynonymHit) GetTypeOk() (*SynonymType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SynonymHit) SetType(v SynonymType) *SynonymHit {
	o.Type = v
	return o
}

// GetSynonyms returns the Synonyms field value if set, zero value otherwise.
func (o *SynonymHit) GetSynonyms() []string {
	if o == nil || o.Synonyms == nil {
		var ret []string
		return ret
	}
	return o.Synonyms
}

// GetSynonymsOk returns a tuple with the Synonyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SynonymHit) GetSynonymsOk() ([]string, bool) {
	if o == nil || o.Synonyms == nil {
		return nil, false
	}
	return o.Synonyms, true
}

// HasSynonyms returns a boolean if a field has been set.
func (o *SynonymHit) HasSynonyms() bool {
	if o != nil && o.Synonyms != nil {
		return true
	}

	return false
}

// SetSynonyms gets a reference to the given []string and assigns it to the Synonyms field.
func (o *SynonymHit) SetSynonyms(v []string) *SynonymHit {
	o.Synonyms = v
	return o
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *SynonymHit) GetInput() string {
	if o == nil || o.Input == nil {
		var ret string
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SynonymHit) GetInputOk() (*string, bool) {
	if o == nil || o.Input == nil {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *SynonymHit) HasInput() bool {
	if o != nil && o.Input != nil {
		return true
	}

	return false
}

// SetInput gets a reference to the given string and assigns it to the Input field.
func (o *SynonymHit) SetInput(v string) *SynonymHit {
	o.Input = &v
	return o
}

// GetWord returns the Word field value if set, zero value otherwise.
func (o *SynonymHit) GetWord() string {
	if o == nil || o.Word == nil {
		var ret string
		return ret
	}
	return *o.Word
}

// GetWordOk returns a tuple with the Word field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SynonymHit) GetWordOk() (*string, bool) {
	if o == nil || o.Word == nil {
		return nil, false
	}
	return o.Word, true
}

// HasWord returns a boolean if a field has been set.
func (o *SynonymHit) HasWord() bool {
	if o != nil && o.Word != nil {
		return true
	}

	return false
}

// SetWord gets a reference to the given string and assigns it to the Word field.
func (o *SynonymHit) SetWord(v string) *SynonymHit {
	o.Word = &v
	return o
}

// GetCorrections returns the Corrections field value if set, zero value otherwise.
func (o *SynonymHit) GetCorrections() []string {
	if o == nil || o.Corrections == nil {
		var ret []string
		return ret
	}
	return o.Corrections
}

// GetCorrectionsOk returns a tuple with the Corrections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SynonymHit) GetCorrectionsOk() ([]string, bool) {
	if o == nil || o.Corrections == nil {
		return nil, false
	}
	return o.Corrections, true
}

// HasCorrections returns a boolean if a field has been set.
func (o *SynonymHit) HasCorrections() bool {
	if o != nil && o.Corrections != nil {
		return true
	}

	return false
}

// SetCorrections gets a reference to the given []string and assigns it to the Corrections field.
func (o *SynonymHit) SetCorrections(v []string) *SynonymHit {
	o.Corrections = v
	return o
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise.
func (o *SynonymHit) GetPlaceholder() string {
	if o == nil || o.Placeholder == nil {
		var ret string
		return ret
	}
	return *o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SynonymHit) GetPlaceholderOk() (*string, bool) {
	if o == nil || o.Placeholder == nil {
		return nil, false
	}
	return o.Placeholder, true
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *SynonymHit) HasPlaceholder() bool {
	if o != nil && o.Placeholder != nil {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given string and assigns it to the Placeholder field.
func (o *SynonymHit) SetPlaceholder(v string) *SynonymHit {
	o.Placeholder = &v
	return o
}

// GetReplacements returns the Replacements field value if set, zero value otherwise.
func (o *SynonymHit) GetReplacements() []string {
	if o == nil || o.Replacements == nil {
		var ret []string
		return ret
	}
	return o.Replacements
}

// GetReplacementsOk returns a tuple with the Replacements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SynonymHit) GetReplacementsOk() ([]string, bool) {
	if o == nil || o.Replacements == nil {
		return nil, false
	}
	return o.Replacements, true
}

// HasReplacements returns a boolean if a field has been set.
func (o *SynonymHit) HasReplacements() bool {
	if o != nil && o.Replacements != nil {
		return true
	}

	return false
}

// SetReplacements gets a reference to the given []string and assigns it to the Replacements field.
func (o *SynonymHit) SetReplacements(v []string) *SynonymHit {
	o.Replacements = v
	return o
}

func (o SynonymHit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	toSerialize["objectID"] = o.ObjectID
	toSerialize["type"] = o.Type
	if o.Synonyms != nil {
		toSerialize["synonyms"] = o.Synonyms
	}
	if o.Input != nil {
		toSerialize["input"] = o.Input
	}
	if o.Word != nil {
		toSerialize["word"] = o.Word
	}
	if o.Corrections != nil {
		toSerialize["corrections"] = o.Corrections
	}
	if o.Placeholder != nil {
		toSerialize["placeholder"] = o.Placeholder
	}
	if o.Replacements != nil {
		toSerialize["replacements"] = o.Replacements
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal SynonymHit: %w", err)
	}

	return serialized, nil
}

func (o SynonymHit) String() string {
	out := ""
	out += fmt.Sprintf("  objectID=%v\n", o.ObjectID)
	out += fmt.Sprintf("  type=%v\n", o.Type)
	out += fmt.Sprintf("  synonyms=%v\n", o.Synonyms)
	out += fmt.Sprintf("  input=%v\n", o.Input)
	out += fmt.Sprintf("  word=%v\n", o.Word)
	out += fmt.Sprintf("  corrections=%v\n", o.Corrections)
	out += fmt.Sprintf("  placeholder=%v\n", o.Placeholder)
	out += fmt.Sprintf("  replacements=%v\n", o.Replacements)
	return fmt.Sprintf("SynonymHit {\n%s}", out)
}
