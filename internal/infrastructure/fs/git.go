package fs

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type gitClient struct {
	tmpDir string
}

func NewGitClient() (*gitClient, error) {
	tmpDir, err := ioutil.TempDir(os.TempDir(), "patchy-")

	if err != nil {
		return &gitClient{}, fmt.Errorf("failed to init git client: %v", err)
	}

	return &gitClient{
		tmpDir: tmpDir,
	}, nil
}

func (g *gitClient) Clone(repo string) error {
	command := fmt.Sprintf("git clone %s %s", repo, g.tmpDir)
	cmdArgs := strings.Split(command, " ")
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	return cmd.Run()
}

func (g *gitClient) Diff(targetBranch string, originBranch string) (string, error) {
	command := fmt.Sprintf("git -C %s diff %s..%s", g.tmpDir, targetBranch, originBranch)
	cmdArgs := strings.Split(command, " ")
	out, err := exec.Command(cmdArgs[0], cmdArgs[1:]...).Output()

	if err != nil {
		return "", err
	}

	return string(out), nil
}
