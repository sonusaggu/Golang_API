package main

import (
	"github.com/gin-gonic/gin"
	"TODOLIST/internal/routes"
	"TODOLIST/internal/middleware"
)

func main(){
	router := gin.Default()

	router.Use(middleware.Logger())

	route.RegisterRoutes(router)

	router.Run(":8081")
}