package patchy

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type GitCliPatcher struct{}

func (cli *GitCliPatcher) Patch(cmd PatchRequest) (Patch, error) {
	workDir, err := ioutil.TempDir("", "repo")
	if err != nil {
		return Patch{}, errors.Wrap(err, "failed to create tmp dir")
	}
	defer os.RemoveAll(workDir)

	err = cli.checkoutRepository(cmd.Repository, workDir)
	if err != nil {
		return Patch{}, errors.Wrap(err, "failed to checkout repository")
	}


	err = cli.checkoutBranch(workDir, cmd.TargetBranch)
	if err != nil {
		return Patch{}, errors.Wrap(err, "failed to checkout branch")
	}

	err = cli.createNewBranch(workDir)
	if err != nil {
		return Patch{}, errors.Wrap(err, "failed to create tmp branch")
	}

	err = cli.squashCommits(workDir, cmd.OriginBranch)
	if err != nil {
		return Patch{}, errors.Wrap(err, "failed to squash commits")
	}

	err = cli.getDiff(workDir, cmd.TargetBranch)
	if err != nil {
		return Patch{}, errors.Wrap(err, "failed to get diff")
	}

	return Patch{}, nil
}

func (cli *GitCliPatcher) checkoutRepository(repo string, tmpDir string) error {
	command := fmt.Sprintf("git clone %s %s", repo, tmpDir)
	cmdArgs := strings.Split(command, " ")
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	return cmd.Run()
}


func (cli *GitCliPatcher) checkoutBranch(tmpDir string, branch string) error {
	command := fmt.Sprintf("git -C %s checkout %s", tmpDir, branch)
	cmdArgs := strings.Split(command, " ")
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	return cmd.Run()
}

func (cli *GitCliPatcher) createNewBranch(tmpDir string) error {
	tmpBranch := "squash_branch"
	command := fmt.Sprintf("git -C %s checkout -b %s", tmpDir, tmpBranch)
	cmdArgs := strings.Split(command, " ")
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	return cmd.Run()
}

func (cli *GitCliPatcher) squashCommits(tmpDir string, originBranch string) error {
	command := fmt.Sprintf("git -C %s merge --squash %s", tmpDir, originBranch)
	cmdArgs := strings.Split(command, " ")
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	return cmd.Run()
}

func (cli *GitCliPatcher) getDiff(tmpDir string, targetBranch string) error {
	command := fmt.Sprintf("git -C %s diff %s", tmpDir, targetBranch)
	cmdArgs := strings.Split(command, " ")
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
