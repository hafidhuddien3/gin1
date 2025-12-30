package handler

import (
    "fmt"
    "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    // Vercel exposes the dynamic param as a query value
    id := r.URL.Query().Get("id")
    fmt.Fprintf(w, "<h1>Book ID: %s</h1>", id)
}
