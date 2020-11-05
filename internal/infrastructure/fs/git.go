package fs

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type GitClient struct {
	tmpDir string
}

var noop = func() { /*nothing to do*/ }

func (g *GitClient) Clone(repo string) (func(), error) {

	tmpDir, err := ioutil.TempDir(os.TempDir(), "patchy-")
	if err != nil {
		return noop, fmt.Errorf("failed to create tmp dir: %v", err)
	}

	g.tmpDir = tmpDir
	cleanup := func() {
		os.RemoveAll(tmpDir)
	}

	command := fmt.Sprintf("git clone %s %s", repo, g.tmpDir)
	cmdArgs := strings.Split(command, " ")
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	return cleanup, cmd.Run()
}

func (g *GitClient) Diff(targetBranch string, originBranch string) (string, error) {

	if len(g.tmpDir) == 0 {
		return "", fmt.Errorf("working dir has not been initialized, run clone first")
	}

	command := fmt.Sprintf("git -C %s diff %s..%s", g.tmpDir, targetBranch, originBranch)
	cmdArgs := strings.Split(command, " ")
	out, err := exec.Command(cmdArgs[0], cmdArgs[1:]...).Output()

	if err != nil {
		return "", err
	}

	return string(out), nil
}
