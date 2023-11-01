package routes

import (
	"github.com/ashikask2002/ecomerce.git/pkg/api/handler"
	"github.com/gin-gonic/gin"
)

func UserRoutes(engin *gin.RouterGroup, userHandler *handler.UserHandler) {
	engin.POST("/signup", userHandler.UserSignUp)
	engin.POST("/login", userHandler.LoginHandler)
}
