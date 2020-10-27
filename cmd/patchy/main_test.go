package main

import "testing"

func Test_ParseFlags_returns_error_on_missing_args(t *testing.T) {

	tt := []struct {
		desc  string
		input []string
	}{
		{"nil array", nil},
		{"empty array", []string{}},
		{"one element in array", []string{"df"}},
		{"two elements in array", []string{"df", "ddd"}},
	}

	for _, testCase := range tt {
		_, err := ParseFlags(testCase.input)

		if err == nil {
			t.Errorf("expected err == nil for case '%v'", testCase.desc)
		}
	}
}

func Test_ParseFlags_returns_cmd_on_valid_args(t *testing.T) {

	expRepository := "github.com/co0p/some/repo"
	expOriginBranch := "feature/branch"
	expTargetBranch := "master"
	args := []string{expRepository, expOriginBranch, expTargetBranch}

	res, err := ParseFlags(args)

	if err != nil {
		t.Errorf("expected err to be nil, got %v instead", err)
	}

	if res.Repository != expRepository {
		t.Errorf("expected %v, got %v instead", expRepository, res.Repository)
	}

	if res.OriginBranch != expOriginBranch {
		t.Errorf("expected %v, got %v instead", expOriginBranch, res.OriginBranch)
	}

	if res.TargetBranch != expTargetBranch {
		t.Errorf("expected %v, got %v instead", expTargetBranch, res.TargetBranch)
	}

}
