package routes

import (
	user "exerciciomongdb/internal/user"
	"github.com/gin-gonic/gin"
)

func NewRouter(userHandler *user.Handlers) *gin.Engine {
	router := gin.Default()
	
	router.POST("/users", userHandler.CreateUser)
	router.GET("/users", userHandler.GetAllUsers)
	router.GET("/users/dpp", userHandler.GetUsersByDpp)
	router.GET("/users/city", userHandler.GetUsersByCity)
	router.GET("/users/age", userHandler.GetUsersByAge)
	
	return router
}