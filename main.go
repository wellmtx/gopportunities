package main

import (
	"gopportunities/config"
	"gopportunities/router"
)

func main() {
	logger := config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	router.Initialize()
}
