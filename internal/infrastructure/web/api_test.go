package web_test

import (
	"bytes"
	"encoding/json"
	"github.com/co0p/patchy/internal/domain"
	"github.com/co0p/patchy/internal/infrastructure/web"
	"net/http"
	"net/http/httptest"
	"testing"
)

var validRequest = web.ApiRequest{
	Repository:   "repo",
	OriginBranch: "origin",
	TargetBranch: "target",
}

var invalidRequest = web.ApiRequest{
	OriginBranch: "origin",
	TargetBranch: "target",
}

func TestApiHandler_returns_diff_as_json_given_valid_request(t *testing.T) {

	apiHandler := web.ApiHandler{Usecase: &mockPatchUsecase{}}
	validPayload, _ := json.Marshal(validRequest)

	req, _ := http.NewRequest("POST", "/", bytes.NewReader(validPayload))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(apiHandler.ServeHTTP)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"repository":"repo","origin_branch":"origin","target_branch":"origin","diff":"some diff "}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

}

func TestApiHandler_returns_badRequest_given_invalid_request(t *testing.T) {

	apiHandler := web.ApiHandler{Usecase: &mockPatchUsecase{}}
	payload, _ := json.Marshal(invalidRequest)

	req, _ := http.NewRequest("POST", "/", bytes.NewReader(payload))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(apiHandler.ServeHTTP)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

type mockPatchUsecase struct{}

func (m mockPatchUsecase) Patch(request domain.PatchRequest) (domain.Patch, error) {
	return domain.Patch{
		RepositoryName: request.Repository,
		Diff:           "some diff ",
		TargetBranch:   request.TargetBranch,
		OriginBranch:   request.OriginBranch,
	}, nil
}
