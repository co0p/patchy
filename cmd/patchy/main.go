package main

import (
	"errors"
	"fmt"
	"github.com/co0p/patchy"
	"os"
)

const ExitError = 1

func main() {

	patchUsecase := patchy.PatchUsecase{}

	req, err := ParseFlags(os.Args[1:])
	if err != nil {
		fmt.Printf("usage: <path/to/remote/repo> <origin branch> <target branch>\n")
		os.Exit(ExitError)
	}

	patch, err := patchUsecase.Generate(req)
	if err != nil {
		fmt.Printf("failed to generate patch: %s\n", err)
		os.Exit(ExitError)
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
