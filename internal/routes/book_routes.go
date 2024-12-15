package route

import (
	"github.com/gin-gonic/gin"
	"TODOLIST/internal/handlers"
)


func RegisterRoutes(r *gin.Engine){
	bookGroup := r.Group("/books")
	{
		bookGroup.GET("/", handlers.GetBooks)
		bookGroup.GET("/:id", handlers.GetBookByID)
		bookGroup.POST("/", handlers.CreateBook)
		bookGroup.PUT("/:id", handlers.UpdateBook)
		bookGroup.DELETE("/:id", handlers.DeleteBook)
	}
}