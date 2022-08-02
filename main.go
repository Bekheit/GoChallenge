package main

import (
	"example/go/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	routes.TransactionRoute(router)

	router.Run("localhost:8080")
}
