package middleware

import (
	"time"
	"log"
	"github.com/gin-gonic/gin"
)


func Logger() gin.HandlerFunc{
	return func(c *gin.Context){
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		log.Printf("latency: %s path: %s method: %s", latency,
		c.Request.URL.Path, c.Request.Method)
		}	
}