package main

import (
	"github.com/lucassoaresfr/go-api.git/config"
	"github.com/lucassoaresfr/go-api.git/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	err := config.Init()

	if err != nil {
		logger.ErrorF("CONFIG INITIALIZATION ERROR: %v", err)
		return
	}

	router.Initialize()
}
