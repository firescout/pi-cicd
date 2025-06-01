package main

import (
	_ "embed"
	"log"
	"os"

	"github.com/firescout/repo-manager/service"
)

var (
	startCon  service.StartController
	sigNotify = make(chan os.Signal)
)

func main() {
	log.Println("Pipeline CI/CD via webhooks")
	log.Println("Manage Repo's over HTTP for Rasberry Pi OS")
	log.Println("Starting service...")

	go initApp()

	sig := <-sigNotify
	log.Println("[WARN] Caught signal", sig)
	if startCon != nil {
		startCon.Stop()
		log.Println("Service stopped")
	}
	os.Exit(0)
}

func initApp() {
	startCon = service.NewService()
	startCon.Start()
}
