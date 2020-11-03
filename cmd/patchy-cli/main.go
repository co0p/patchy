package main

import (
	"fmt"
	"github.com/co0p/patchy/internal/application"
	"github.com/co0p/patchy/internal/infrastructure/cli"
	"github.com/co0p/patchy/internal/infrastructure/fs"
	"os"
)

const ExitErrorCode = 1

func main() {

	gitOperator, err := fs.NewGitClient()
	if err != nil {
		fmt.Println(err)
		os.Exit(ExitErrorCode)
	}

	patchUsecase := application.PatchUsecase{Git: gitOperator}
	cli := cli.Patchy{Usecase: &patchUsecase}

	str, err := cli.Run(os.Args[1:])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(str)
}
