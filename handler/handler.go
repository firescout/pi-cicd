package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"

	api "github.com/firescout/repo-manager/restserver"
)

func NewHandler() api.DefaultApiServicer {
	repos := new(Repos)
	file, err := os.ReadFile("settings.json")
	if err != nil {
		panic("Failed to read repos.json: " + err.Error())
	}
	err = json.Unmarshal(file, repos)
	if err != nil {
		panic("Failed to Unmarshal or parse repos.json: " + err.Error())
	}
	if runtime.GOOS != repos.System {
		panic(fmt.Sprintf("System mismatch: Expected %s but got %s", runtime.GOOS, repos.System))
	}
	return &HTTPHandler{
		repoCloneDir:  repos.ClonePath,
		pathSeparator: string(os.PathSeparator),
		repos:         *repos,
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
