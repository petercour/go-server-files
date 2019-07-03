// golang static file web server
// https://golangr.com
package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, r.URL.Path[1:])
    })

    log.Fatal(http.ListenAndServe(":8082", nil))
}