package application_test

import (
	"errors"
	"github.com/co0p/patchy/internal/application"
	"github.com/co0p/patchy/internal/domain"
	"testing"
)

func Test_PatchUsecase_succeeds_given_validRequestAndOperator(t *testing.T) {

	request := domain.PatchRequest{"repo", "origin", "master"}
	usecase := application.PatchUsecase{Git: MockGitOperator{}}

	_, err := usecase.Patch(request)
	if err != nil {
		t.Errorf("expected err to be nil, got '%v'", err)
	}
}

func Test_PatchUsecase_fails_given_invalidRequest(t *testing.T) {

	cases := []struct {
		description    string
		invalidRequest domain.PatchRequest
	}{
		{"empty repo", domain.PatchRequest{"", "ob", "tb"}},
		{"empty origin", domain.PatchRequest{"", "", "tb"}},
		{"empty target", domain.PatchRequest{"repo", "ob", ""}},
	}

	for _, tt := range cases {
		usecase := application.PatchUsecase{Git: MockGitOperator{}}

		_, err := usecase.Patch(tt.invalidRequest)
		if err == nil {
			t.Errorf("expected err not to be nil, when '%v'", tt.description)
		}
	}
}

func Test_PatchUsecase_fails_given_failingGitOperation(t *testing.T) {

	cases := []struct {
		description     string
		failingOperator application.GitOperator
	}{
		{"clone fails", NewMockGitOperator(true, false)},
		{"diff fails", NewMockGitOperator(false, true)},
	}

	for _, tt := range cases {
		usecase := application.PatchUsecase{Git: tt.failingOperator}

		_, err := usecase.Patch(domain.PatchRequest{})
		if err == nil {
			t.Errorf("expected err not to be nil, when '%v'", tt.description)
		}
	}
}

type MockGitOperator struct {
	fail map[string]bool
}

func NewMockGitOperator(cloneFails, diffFails bool) MockGitOperator {
	var failureMap = make(map[string]bool)
	failureMap["clone"] = cloneFails
	failureMap["diff"] = diffFails
	return MockGitOperator{failureMap}
}

func (m MockGitOperator) Clone(s string) (func(), error) {

	return func() { /*noop*/ }, errWhenSet(m.fail, "clone")
}

func (m MockGitOperator) Diff(a, b string) (string, error) {
	return "", errWhenSet(m.fail, "diff")
}

func errWhenSet(failMap map[string]bool, key string) error {
	val, ok := failMap[key]
	if ok && val {
		return errors.New(key + " failed")
	}
	return nil
}
