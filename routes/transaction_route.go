package routes

import (
	"example/go/controllers"
	"github.com/gin-gonic/gin"
)

func TransactionRoute(router *gin.Engine) {
	router.GET("/transactions", controllers.GetAll())
	router.POST("/transactions", controllers.CreateTransaction())
}
