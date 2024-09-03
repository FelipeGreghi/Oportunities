package main

import (
	"github.com/FelipeGreghi/Oportunities/config"
	"github.com/FelipeGreghi/Oportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("Oportunities")

	// Initialize Config
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}

	// Initializes the router
	router.Init()
}
