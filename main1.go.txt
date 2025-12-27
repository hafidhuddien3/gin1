package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func main() {
  router := gin.Default()
  router.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong789",
    })
  })
  router.GET("/pinghttp", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pongHTTP",
    })
  })
  router.Run() // listens on 0.0.0.0:8080 by default
}
//