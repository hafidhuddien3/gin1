package handler
 
import (
  // "fmt"
  "net/http"
   "encoding/json"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Marshal the request into JSON
    data, err := json.MarshalIndent(r, "", "  ")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Write JSON to response
    w.Write(data)
}
