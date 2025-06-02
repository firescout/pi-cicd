package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	api "github.com/firescout/repo-manager/restserver"
)

type HTTPHandler struct {
	repos Repos
}

type StdResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Repo struct {
	Name        string        `json:"name"`
	Url         string        `json:"url"`
	AfterScript []AfterScript `json:"after_script"`
	Path        string        `json:"path"`
}
type AfterScript struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

type Repos struct {
	Repos []Repo `json:"repos"`
}

func NewHandler() api.DefaultApiServicer {
	repos := new(Repos)
	file, err := os.ReadFile("firescout.json")
	if err != nil {
		panic("Failed to read repos.json: " + err.Error())
	}
	err = json.Unmarshal(file, repos)
	if err != nil {
		panic("Failed to Unmarshal or parse repos.json: " + err.Error())
	}
	return &HTTPHandler{
		repos: *repos,
	}
}

func (s *HTTPHandler) OnPush(ctx context.Context, repoName string) (api.ImplResponse, error) {
	res := StdResponse{
		Status:  "success",
		Message: "Received push for repository: " + repoName,
	}
	if repoName == "" {
		res.Status = "error"
		res.Message = "Repository name cannot be empty"
		return api.Response(http.StatusBadRequest, res), nil
	}

	err := s.handleOnPush(repoName)
	if err != nil {
		res.Status = "error"
		res.Message = "Error handling push for repository: " + repoName + " - " + err.Error()
		return api.Response(http.StatusInternalServerError, res), nil
	}

	return api.Response(http.StatusOK, res), nil
}

func (s *HTTPHandler) GetShutdown(ctx context.Context) (api.ImplResponse, error) {
	res := StdResponse{
		Status:  "success",
		Message: "Shutting down the system",
	}
	return api.Response(http.StatusOK, res), nil
}
