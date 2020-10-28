package patchy

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
)

type GoGitPatcher struct{}

func (u *GoGitPatcher) Patch(request PatchRequest) (Patch, error) {

	// clone the repo to work with
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: request.Repository,
	})
	if err != nil {
		return Patch{}, fmt.Errorf("failed to clone repository %v: %s", request.Repository, err)
	}

	head, err := repo.Head()
	fmt.Printf("clone of %v successfull, now pointing to HEAD: %v\n", request.Repository, head)

	// checkout new tmp branch from target branch

	// merge originBranch onto tmp branch

	// squash commits

	// generate diff for one commit

	return Patch{}, nil
}
