package handler
 
import (
  "fmt"
  "net/http"
)
 
func Handler2(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>Hello from 2!</h1>")
  // Set response header 
//   w.Header().Set("Content-Type", "application/json") // Marshal the request struct into JSON 
//   data, err := json.MarshalIndent(r, "", " ") if err != nil { http.Error(w, err.Error(), http.StatusInternalServerError) return } // Write JSON to response 
//   w.Write(data)
}