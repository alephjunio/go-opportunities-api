package main

import (
	"github.com/alephjunio/go-opportunities-api/config"
	"github.com/alephjunio/go-opportunities-api/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("Config Initialization error: %v", err)
		return
	}

	router.Inizialize()

}
