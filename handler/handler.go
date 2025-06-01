package handler

import (
	"context"
	"net/http"

	api "github.com/firescout/repo-manager/restserver"
)

type HTTPHandler struct {
}

type StdResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Repo struct {
	Url         string `json:"url"`
	AfterScript string `json:"after_script"`
}

type Repos struct {
	Repos []Repo `json:"repos"`
}

func NewHandler() api.DefaultApiServicer {
	return &HTTPHandler{}
}

func (s *HTTPHandler) OnPush(ctx context.Context, repoName string) (api.ImplResponse, error) {
	if repoName == "" {
		return api.Response(http.StatusBadRequest, StdResponse{
			Status:  "error",
			Message: "Repository name is required",
		}), nil
	}
	res := StdResponse{
		Status:  "success",
		Message: "Received push for repository: " + repoName,
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
