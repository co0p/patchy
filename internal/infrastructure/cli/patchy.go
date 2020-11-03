package cli

import (
	"errors"
	"fmt"
	"github.com/co0p/patchy/internal/domain"
)

type Patchy struct {
	Usecase domain.Patcher
}

func (p *Patchy) Run(args []string) (string, error) {

	request, err := parseArgs(args)

	if err != nil {
		return "", fmt.Errorf("failed to parse args: %v", err)
	}

	res, err := p.Usecase.Patch(request)

	if err != nil {
		return "", fmt.Errorf("failed to create patch: %v", err)
	}

	return res.Diff, nil
}

func parseArgs(args []string) (domain.PatchRequest, error) {

	if len(args) < 3 {
		return domain.PatchRequest{}, errors.New("not enough arguments")
	}

	return domain.PatchRequest{
		Repository:   args[0],
		OriginBranch: args[1],
		TargetBranch: args[2],
	}, nil
}
