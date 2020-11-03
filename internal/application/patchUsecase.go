package application

import (
	"fmt"
	"github.com/co0p/patchy/internal/domain"
)

type GitOperator interface {
	Clone(string) error
	Checkout(string) error
	Branch(string) error
	Squash(string) error
	Diff(string) (string, error)
}

type PatchUsecase struct {
	Git GitOperator
}

func (u *PatchUsecase) Patch(cmd domain.PatchRequest) (domain.Patch, error) {

	if valid := cmd.Valid(); valid == false {
		return domain.Patch{}, fmt.Errorf("invalid cmd given")
	}

	err := u.Git.Clone(cmd.Repository)
	if err != nil {
		return domain.Patch{}, fmt.Errorf("failed to clone repository: %v", err)
	}

	err = u.Git.Checkout(cmd.TargetBranch)
	if err != nil {
		return domain.Patch{}, fmt.Errorf("failed to checkout branch: %v", err)
	}

	err = u.Git.Branch("tmp_branch")
	if err != nil {
		return domain.Patch{}, fmt.Errorf("failed to create tmp branch: %v", err)
	}

	err = u.Git.Squash(cmd.OriginBranch)
	if err != nil {
		return domain.Patch{}, fmt.Errorf("failed to squash commits: %v", err)
	}

	diff, err := u.Git.Diff(cmd.TargetBranch)
	if err != nil {
		return domain.Patch{}, fmt.Errorf("failed to create diff: %v", err)
	}

	return domain.Patch{
		Diff: diff,
	}, nil
}
