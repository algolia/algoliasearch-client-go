package search

import (
	"testing"
)

func TestIndex_GetAppID(t *testing.T) {
	tests := []struct {
		name  string
		appID string
		want  string
	}{
		{
			appID: "test_app_id",
			want:  "test_app_id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Index{
				appID: tt.appID,
			}
			if got := i.GetAppID(); got != tt.want {
				t.Errorf("GetAppID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_GetName(t *testing.T) {
	tests := []struct {
		name      string
		indexName string
		want      string
	}{
		{
			indexName: "test_index",
			want: "test_index",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Index{
				name: tt.indexName,
			}
			if got := i.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}
