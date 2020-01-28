package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/psenna/isup/api"
	"github.com/psenna/isup/dependencies"
	"github.com/psenna/isup/models"
)

func main() {

	setupCloseHandler()

	// handler configuration definition errors
	config, _ := models.CreateConfig("{}")

	dependencies.InitDependencies(config.ToDependenciesCOnfigurations())

	server := api.GetApiServer()

	server.Run("0.0.0.0:8080")
}

func setupCloseHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	signal.Notify(c, os.Interrupt, syscall.SIGQUIT)
	go func() {
		<-c
		fmt.Println("\nClosing and stopping app...")
		dependencies.CloseDependencies()
		os.Exit(0)
	}()
}
