package routes

import (
	"TODOLIST/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, handler *handlers.BookHandler) {

	bookGroup := r.Group("/books")
	{
		bookGroup.GET("/", handler.GetBooks)
		bookGroup.GET("/:id", handlers.GetBookByID)
		bookGroup.POST("/", handler.CreateBook)
		bookGroup.PUT("/:id", handlers.UpdateBook)
		bookGroup.DELETE("/:id", handlers.DeleteBook)
	}
}
