package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sebasvelasco353/nummus/server/internal/config"
	"github.com/sebasvelasco353/nummus/server/internal/database"
)

func main() {
	fmt.Println("Hello from nummus server main")

	config.Load()
	database.InitDB()

	server := gin.Default()
	server.Run()
}
