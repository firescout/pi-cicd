package service

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/firescout/repo-manager/handler"
	"github.com/firescout/repo-manager/restserver"
)

type StartController interface {
	Start()
	Stop()
}

type Service struct {
	server *http.Server
}

func NewService() StartController {
	return &Service{
		server: &http.Server{
			Addr: "0.0.0.0:9090",
		},
	}
}

func (n *Service) Start() {
	log.Println(strings.Repeat("=", 10) + " START " + strings.Repeat("=", 10))
	log.Println("Starting service...")
	router := restserver.NewDefaultApiController(handler.NewHandler())
	muxRouter := restserver.NewRouter(router)

	n.server.Handler = muxRouter

	go n.bootServer()

	log.Println("listening on...", n.server.Addr)
}

func (n *Service) bootServer() {
	err := n.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

func (n *Service) Stop() {
	log.Println(strings.Repeat("=", 10) + " STOP " + strings.Repeat("=", 10))
	if err := n.server.Shutdown(context.Background()); err != nil {
		panic(err)
	}
}
