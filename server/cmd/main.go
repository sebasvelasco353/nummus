package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sebasvelasco353/nummus/server/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "configuration error: %v\n", err)
		os.Exit(1)
	}

	server := gin.Default()
	server.Run()

	log.Printf("server listening on port %s", cfg.Server.Port)
}
