package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    // Get configuration from environment variables
    port := os.Getenv("PORT")
    secretMessage := os.Getenv("SECRET_MESSAGE")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello World!\nSecret message: %s", secretMessage)
    })

    fmt.Printf("Listening on port %s\n", port)
    http.ListenAndServe(":" + port, nil)
}
