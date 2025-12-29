package handler

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

var ginEngine *gin.Engine

// func init() {

// }

// Vercel entry point
func Handler(w http.ResponseWriter, r *http.Request) {
        ginEngine = gin.Default()
	ginEngine.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Hello from Gin on Vercel!"})
    })
    ginEngine.GET("/hello", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Hello from Gin on Vercel!"})
    })
	ginEngine.GET("/hello2", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Hello2 from Gin on Vercel!"})
    })
    
    ginEngine.ServeHTTP(w, r)
}