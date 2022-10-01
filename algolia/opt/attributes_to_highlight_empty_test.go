package opt

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestAttributesToHighlightEmptyList(t *testing.T) {
	got, err := json.Marshal(AttributesToHighlightEmptyList())
	if err != nil {
		t.Fatal(err)
	}
	want := []byte("[]")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("AttributesToHighlightEmptyList() = %v, want %v", got, want)
	}
}
