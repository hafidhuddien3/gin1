
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

func HandlerSwagger(w http.ResponseWriter, r *http.Request) {

    db.InitDB()

    g := gin.Default()

    g.GET("/api/", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "ping",
    })
    })

    g.GET("/api/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "ping",
    })
    })

    // Routes
	// if r.Method == "POST" {
	// } else {
	// }
	g.POST("/api/books", handlers.CreateBook)
	g.GET("/api/books", handlers.GetBooks)
    g.GET("/api/books/:id", handlers.GetBook)
    g.PUT("/api/books/:id", handlers.UpdateBook)
    g.DELETE("/api/books/:id", handlers.DeleteBook)

    g.POST("/api/ingredients", handlers.CreateIngredient)
    g.GET("/api/ingredients", handlers.GetIngredients)
    g.PUT("/api/ingredients/:id", handlers.UpdateIngredient)
    g.DELETE("/api/ingredients/:id", handlers.DeleteIngredient)

    g.POST("/api/products", handlers.CreateProduct)
    g.GET("/api/products", handlers.GetProducts)
    g.PUT("/api/products/:id", handlers.UpdateProduct)
    g.DELETE("/api/products/:id", handlers.DeleteProduct)

    g.POST("/api/productIngredients", handlers.CreateProductIngredient)
    g.GET("/api/productIngredients/product/:productId", handlers.GetProductIngredientByProductId)
    g.PUT("/api/productIngredients/:id", handlers.UpdateProductIngredient)
    g.DELETE("/api/productIngredients/:id", handlers.DeleteProductIngredient)

    // Routes
    g.POST("/api/pattern", handlers.SearchPattern)

    // Swagger
    g.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // r.Run(":8080")
    g.ServeHTTP(w, r)
}