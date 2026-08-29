package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sebasvelasco353/nummus/server/internal/config"
	"github.com/sebasvelasco353/nummus/server/internal/users"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "configuration error: %v\n", err)
		os.Exit(1)
	}

	// Connect to the database
	if err := config.ConnectDB(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "database error: %v\n", err)
		os.Exit(1)
	}
	defer config.DB.Close()

	server := gin.Default()
	users.RegisterRoutes(server)

	server.Run(":" + cfg.Server.Port)
}
