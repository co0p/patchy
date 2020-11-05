package application

import (
	"fmt"
	"github.com/co0p/patchy/internal/domain"
)

type GitOperator interface {
	Clone(string) (func(), error)
	Diff(string, string) (string, error)
}

type PatchUsecase struct {
	Git GitOperator
}

func (u *PatchUsecase) Patch(cmd domain.PatchRequest) (domain.Patch, error) {

	if valid := cmd.Valid(); valid == false {
		return domain.Patch{}, fmt.Errorf("invalid cmd given")
	}

	cleanup, err := u.Git.Clone(cmd.Repository)
	defer cleanup()

	if err != nil {
		return domain.Patch{}, fmt.Errorf("failed to clone repository: %v", err)
	}

	diff, err := u.Git.Diff(cmd.TargetBranch, cmd.OriginBranch)
	if err != nil {
		return domain.Patch{}, fmt.Errorf("failed to create diff: %v", err)
	}

	return domain.Patch{
		Diff: diff,
	}, nil
}
