package main

import (
	"github.com/psenna/isup/api"
	"github.com/psenna/isup/dependencies"
)

func main() {

	dependencies.InitDependencies()

	defer dependencies.CloseDependencies()

	server := api.GetApiServer()

	server.Run("0.0.0.0:8080")
}
