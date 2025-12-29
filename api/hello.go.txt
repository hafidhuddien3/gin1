package handler
 
import (
  // "fmt"
  "net/http"
   "encoding/json"
   "io"
)
 
type RequestDump struct {
    Method string              `json:"method"`
    URL    string              `json:"url"`
    Header map[string][]string `json:"header"`
    Body   string              `json:"body,omitempty"`
}

func HandlerHello(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Read body (if any)
    bodyBytes, _ := io.ReadAll(r.Body)

    dump := RequestDump{
        Method: r.Method,
        URL:    r.URL.String(),
        Header: r.Header,
        Body:   string(bodyBytes),
    }

    // Marshal into JSON
    data, err := json.MarshalIndent(dump, "", "  ")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Write(data)
}


// {"seq":[2,4,8]}

// {
//   "method": "POST",
//   "url": "/pattern",
//   "header": {
//     "Content-Type": ["application/json"],
//     "User-Agent": ["curl/7.64.1"]
//   },
//   "body": "{\"seq\":[2,4,8]}"
// }
