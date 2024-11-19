package main

import (
	"ecommers/router"
	"ecommers/util"
	"fmt"
	"log"
	"net/http"

	"go.uber.org/zap"
)

func main() {

	r, logger, err := router.InitRouter()
	if err != nil {
		log.Fatalf("Failed to initialize router: %v", err)
	}
	defer logger.Sync()

	config, err := util.ReadConfiguration()
	if err != nil {
		logger.Fatal("Failed to read configuration", zap.Error(err))
	}

	port := config.Port
	logger.Info("Starting server", zap.String("port", port))
	addr := fmt.Sprintf(":%s", port)

	if err := http.ListenAndServe(addr, r); err != nil {
		logger.Fatal("Server failed to start", zap.Error(err))
	}
}
