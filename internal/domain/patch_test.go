package domain_test

import (
	"github.com/co0p/patchy/internal/domain"
	"testing"
)

func Test_PatchRequest_Valid_succeeds(t *testing.T) {

	request := domain.PatchRequest{"repo", "origin", "target"}
	if !request.Valid() {
		t.Errorf("expect request to be valid, got %v", request.Valid())
	}
}

func Test_PatchRequest_Valid_fails(t *testing.T) {
	type fields struct {
		Repository   string
		OriginBranch string
		TargetBranch string
	}

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"all empty", fields{"", "", ""}, false},
		{"repo missing", fields{"", "df", "df"}, false},
		{"origin missing", fields{"repo", "", "df"}, false},
		{"target missing", fields{"", "", "df"}, false},
	}
	for _, tt := range tests {
		r := domain.PatchRequest{
			Repository:   tt.fields.Repository,
			OriginBranch: tt.fields.OriginBranch,
			TargetBranch: tt.fields.TargetBranch,
		}
		if got := r.Valid(); got != tt.want {
			t.Errorf("Valid() = %v, want %v", got, tt.want)
		}
	}
}
