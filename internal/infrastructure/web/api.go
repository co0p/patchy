package web

import (
	"encoding/json"
	"fmt"
	"github.com/co0p/patchy/internal/domain"
	"io/ioutil"
	"net/http"
)

type ApiHandler struct {
	Usecase domain.PatchUsecase
}

type ApiRequest struct {
	Repository   string `json:"repository"`
	OriginBranch string `json:"origin_branch"`
	TargetBranch string `json:"target_branch"`
}

type ApiResponse struct {
	Repository   string `json:"repository"`
	OriginBranch string `json:"origin_branch"`
	TargetBranch string `json:"target_branch"`
	Diff         string `json:"diff"`
}

func (h *ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	bytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "failed to get data", http.StatusBadRequest)
		return
	}

	var request ApiRequest
	err = json.Unmarshal(bytes, &request)
	if err != nil {
		http.Error(w, "failed to parse data", http.StatusBadRequest)
		return
	}

	patchRequest := domain.PatchRequest{
		Repository:   request.Repository,
		OriginBranch: request.OriginBranch,
		TargetBranch: request.TargetBranch,
	}

	if !patchRequest.Valid() {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	patch, err := h.Usecase.Patch(patchRequest)
	if err != nil {
		fmt.Errorf("failed to generate patch: %v", err)
		http.Error(w, "failed to generate patch", http.StatusBadRequest)
		return
	}

	response := ApiResponse{
		Repository:   request.Repository,
		OriginBranch: request.OriginBranch,
		TargetBranch: request.OriginBranch,
		Diff:         patch.Diff,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Errorf("failed to marshal response: %v", err)
		http.Error(w, "failed to generate path", http.StatusInternalServerError)
		return
	}

	header := w.Header()
	header.Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
