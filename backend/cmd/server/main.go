package main

import (
    "net/http"
    "log"
    "path/filepath"
)

func main() {
    // Construct the correct path to the build folder
    frontendBuildPath := filepath.Join("..", "frontend", "build")
    
    // Serve static files from the build directory
    fs := http.FileServer(http.Dir(frontendBuildPath))
    http.Handle("/", fs)

    // API endpoints
    http.HandleFunc("/llm", helloHandler)

    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{"message": "Hello from Go!"}`))
}

