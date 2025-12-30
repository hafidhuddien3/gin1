// package main
package handler

import (
    _ "gin-quickstart/docs"
    "gin-quickstart/db"
    "gin-quickstart/handlers"

    "github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
      "net/http"
)

// func main() {
func HandlerMain(w http.ResponseWriter, r *http.Request) {
    db.InitDB()

    g := gin.Default()

    g.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "ping",
    })
    })

    g.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "ping",
    })
    })

    // Routes
    g.POST("/books", handlers.CreateBook)
    g.GET("/books", handlers.GetBooks)
	g.GET("/api/books", handlers.GetBooks)
    g.GET("/books/:id", handlers.GetBook)
    g.PUT("/books/:id", handlers.UpdateBook)
    g.DELETE("/books/:id", handlers.DeleteBook)

    g.POST("/ingredients", handlers.CreateIngredient)
    g.GET("/ingredients", handlers.GetIngredients)
    g.PUT("/ingredients/:id", handlers.UpdateIngredient)
    g.DELETE("/ingredients/:id", handlers.DeleteIngredient)

    g.POST("/products", handlers.CreateProduct)
    g.GET("/products", handlers.GetProducts)
    g.PUT("/products/:id", handlers.UpdateProduct)
    g.DELETE("/products/:id", handlers.DeleteProduct)

    g.POST("/productIngredients", handlers.CreateProductIngredient)
    g.GET("/productIngredients/product/:productId", handlers.GetProductIngredientByProductId)
    g.PUT("/productIngredients/:id", handlers.UpdateProductIngredient)
    g.DELETE("/productIngredients/:id", handlers.DeleteProductIngredient)

    // Routes
    g.POST("/pattern", handlers.SearchPattern)

    // Swagger
    g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // r.Run(":8080")
    // r.Run()
    g.ServeHTTP(w, r)
}

// func Handler(w http.ResponseWriter, r *http.Request) {
//     var ginEngine *gin.Engine
//     ginEngine = gin.Default()
// 	  ginEngine.GET("/", func(c *gin.Context) {
//         c.JSON(200, gin.H{"message": "Hello from Gin on Vercel!"})
//     })
//     // ginEngine.GET("/hello", func(c *gin.Context) {
//     //     c.JSON(200, gin.H{"message": "Hello from Gin on Vercel!"})
//     // })
// 	  ginEngine.GET("/hello2", func(c *gin.Context) {
//         c.JSON(200, gin.H{"message": "Hello2 from Gin on Vercel!"})
//     })

//     ginEngine.ServeHTTP(w, r)
// }