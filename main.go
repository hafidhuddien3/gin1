package main

import (
    _ "gin-quickstart/docs"
    "gin-quickstart/db"
    "gin-quickstart/handlers"

    "github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
)

func main() {
    db.InitDB()

    r := gin.Default()

    // Routes
    r.POST("/books", handlers.CreateBook)
    r.GET("/books", handlers.GetBooks)
    r.GET("/books/:id", handlers.GetBook)
    r.PUT("/books/:id", handlers.UpdateBook)
    r.DELETE("/books/:id", handlers.DeleteBook)

    // Swagger
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    r.Run(":8080")
}
