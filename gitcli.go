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
	tmpDir, err := createTempDir()
	if err != nil {
		return Patch{}, errors.Wrap(err, "failed to create tmp dir")
	}
	defer os.RemoveAll(tmpDir)

	err = cli.checkoutRepository(cmd.Repository, tmpDir)
	if err != nil {
		return Patch{}, errors.Wrap(err, "failed to checkout repository")
	}

	return Patch{}, nil
}

func (cli *GitCliPatcher) checkoutRepository(repo string, tmpDir string) error {

	command := fmt.Sprintf("git clone %s %s", repo, tmpDir)
	cmdArgs := strings.Split(command, " ")
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	return cmd.Run()
}

func createTempDir() (string, error) {
	return ioutil.TempDir("", "repo")
}
