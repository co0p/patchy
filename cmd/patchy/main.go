package main

import (
	"errors"
	"fmt"
	"github.com/co0p/patchy"
	"os"
)

const ExitErrorCode = 1

func main() {

	patcher := patchy.GitCliPatcher{}

	patchCmd, err := ParseFlags(os.Args[1:])
	if err != nil {
		fmt.Printf("usage: <path/to/remote/repo> <origin branch> <target branch>\n")
		os.Exit(ExitErrorCode)
	}

	patch, err := patcher.Patch(patchCmd)
	if err != nil {
		fmt.Printf("failed to generate patch: %s\n", err)
		os.Exit(ExitErrorCode)
	}

	fmt.Printf(patch.Diff)
}

func ParseFlags(args []string) (patchy.PatchRequest, error) {

	if len(args) < 3 {
		return patchy.PatchRequest{}, errors.New("not enough arguments")
	}

	return patchy.PatchRequest{
		Repository:   args[0],
		OriginBranch: args[1],
		TargetBranch: args[2],
	}, nil
}
