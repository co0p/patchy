package main

import (
	"fmt"
	"github.com/co0p/patchy/internal/application"
	"github.com/co0p/patchy/internal/infrastructure/cli"
	"github.com/co0p/patchy/internal/infrastructure/fs"
	"os"
)

func main() {

	client := fs.GitClient{}
	usecase := application.PatchUsecase{Git: &client}
	cli := cli.Patchy{Usecase: &usecase}

	str, err := cli.Run(os.Args[1:])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(str)
}
