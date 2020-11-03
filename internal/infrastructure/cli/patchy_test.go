package cli_test

import (
	"errors"
	"github.com/co0p/patchy/internal/domain"
	"github.com/co0p/patchy/internal/infrastructure/cli"
	"testing"
)

func Test_Run_returnsError_given_invalidArguments(t *testing.T) {

	cases := []struct {
		description string
		input       []string
	}{
		{"empty args", []string{}},
		{"args with only repo", []string{"single_entry"}},
		{"args with repo and origin branch", []string{"repo", "originBranch"}},
	}

	cli := cli.Patchy{
		Usecase: &PatchUsecaseMock{false},
	}

	for _, tt := range cases {

		_, err := cli.Run(tt.input)
		if err == nil {
			t.Errorf("expected cli to return err when '%v', got not nil instead ", tt.description)
		}
	}
}

func Test_Run_returnsDiff_given_noError(t *testing.T) {

	validArgs := []string{"repo", "origin", "master"}

	cli := cli.Patchy{
		Usecase: &PatchUsecaseMock{false},
	}

	str, err := cli.Run(validArgs)
	if err != nil {
		t.Errorf("expected err to be nil, got %v", err)
	}

	if len(str) == 0 {
		t.Errorf("expected str to have data, got empty str instead")
	}
}

type PatchUsecaseMock struct {
	Fail bool
}

func (p PatchUsecaseMock) Patch(request domain.PatchRequest) (domain.Patch, error) {
	if p.Fail {
		return domain.Patch{}, errors.New("failed because mock was asked to")
	}

	return domain.Patch{
		RepositoryName: request.Repository,
		Diff:           "some Diff",
		TargetBranch:   request.TargetBranch,
		OriginBranch:   request.OriginBranch,
	}, nil
}
