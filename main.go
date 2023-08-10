package main

import (
	"fmt"
	"go-restAPI/configs"

	"go-restAPI/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Get the MongoDB client from the config package
	client := configs.Client
	fmt.Println(client)
	router := gin.Default()
	routers.CartRoute(router)
	router.Run("localhost:6000")

}
