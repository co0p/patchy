package main

import (
	"fmt"
	"github.com/co0p/patchy/internal/application"
	"github.com/co0p/patchy/internal/infrastructure/fs"
	"github.com/co0p/patchy/internal/infrastructure/web"
	"net/http"
	"os"
	"strconv"
)

const defaultPort = 8080

func main() {
	port := getPort(os.Getenv("PORT"))

	gitClient := fs.GitClient{}
	usecase := application.PatchUsecase{Git: &gitClient}
	apiHandler := web.ApiHandler{Usecase: &usecase}

	mux := http.NewServeMux()
	mux.Handle("/api/diff", &apiHandler)

	addr := fmt.Sprintf(":%d", port)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Printf("failed to start server: %v", err)
	}
}

func getPort(env string) int {
	i, err := strconv.Atoi(env)
	if err != nil {
		return defaultPort
	}
	return i
}
