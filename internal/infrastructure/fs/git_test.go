// +build integration

package fs_test

import (
	"github.com/co0p/patchy/internal/infrastructure/fs"
	"io/ioutil"
	"testing"
)

// Constants must use point to the test repo and correct branch name
const (
	repoPath   = "../../../test/testRepository"
	branchName = "origin/testBranch"
	diffPath   = "../../../test/testRepository.diff"
)

func TestGitClient_Clone_fails(t *testing.T) {
	gitClient := fs.GitClient{}

	cleanup, err := gitClient.Clone("_some_invalid_path")
	if err == nil {
		t.Errorf("expected err not to be nil, got %v", err)
	}
	defer cleanup()
}

func TestGitClient_Clone_succeeds(t *testing.T) {
	gitClient := fs.GitClient{}
	cleanup, err := gitClient.Clone(repoPath)
	defer cleanup()

	if err != nil {
		t.Errorf("expected err to be nil, got %v", err)
	}
}

func TestGitClient_Diff_same_branch(t *testing.T) {
	gitClient := fs.GitClient{}
	gitClient.Clone(repoPath)

	diff, err := gitClient.Diff("master", "master")
	if err != nil {
		t.Errorf("expected err to be nil, got %v", err)
	}

	if len(diff) != 0 {
		t.Errorf("expected diff to be emtpy, got %v", diff)
	}
}

func TestGitClient_Diff_different_branches(t *testing.T) {
	gitClient := fs.GitClient{}
	gitClient.Clone(repoPath)

	diff, err := gitClient.Diff("origin/master", branchName)

	if err != nil {
		t.Errorf("expected err to be nil, got %v", err)
	}

	fixture, _ := ioutil.ReadFile(diffPath)
	if diff != string(fixture) {
		t.Errorf("expected diff to equal fixture, got %v", diff)
	}
}
