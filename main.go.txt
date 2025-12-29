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

    r.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "ping",
    })
    })

    r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "ping",
    })
    })

    // Routes
    r.POST("/books", handlers.CreateBook)
    r.GET("/books", handlers.GetBooks)
    r.GET("/books/:id", handlers.GetBook)
    r.PUT("/books/:id", handlers.UpdateBook)
    r.DELETE("/books/:id", handlers.DeleteBook)

// get ingredient v
// get product v
// get product_ingredient by productId
// post ingredient v
// post product v
// post product_ingredient v

    r.POST("/ingredients", handlers.CreateIngredient)
    r.GET("/ingredients", handlers.GetIngredients)
    // r.GET("/ingredients/:id", handlers.GetIngredient)
    r.PUT("/ingredients/:id", handlers.UpdateIngredient)
    r.DELETE("/ingredients/:id", handlers.DeleteIngredient)

    r.POST("/products", handlers.CreateProduct)
    r.GET("/products", handlers.GetProducts)
    r.PUT("/products/:id", handlers.UpdateProduct)
    r.DELETE("/products/:id", handlers.DeleteProduct)

    r.POST("/productIngredients", handlers.CreateProductIngredient)
    r.GET("/productIngredients/product/:productId", handlers.GetProductIngredientByProductId)
    r.PUT("/productIngredients/:id", handlers.UpdateProductIngredient)
    r.DELETE("/productIngredients/:id", handlers.DeleteProductIngredient)

    // Routes
    r.POST("/pattern", handlers.SearchPattern)

    // Swagger
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // r.Run(":8080")
    r.Run()
}
