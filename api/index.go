// package handler
 
// import (
//   "fmt"
//   "net/http"
// )
 
// func HandlerIndex(w http.ResponseWriter, r *http.Request) {
//   fmt.Fprintf(w, "<h1>Hello from index!</h1>")
//   // Set response header 
// //   w.Header().Set("Content-Type", "application/json") // Marshal the request struct into JSON 
// //   data, err := json.MarshalIndent(r, "", " ") if err != nil { http.Error(w, err.Error(), http.StatusInternalServerError) return } // Write JSON to response 
// //   w.Write(data)
// }

package handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    g := gin.Default()

    g.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Hello from Gin on root!"})
    })

	g.GET("/api", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Hello from Gin on api!"})
    })

    g.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

    g.ServeHTTP(w, r)
}
