package main

import (
	"fmt"
	"net/http"
	"github.com/ovrdis/goforex/pkg/handlers"
)

func main() {
	http.HandleFunc("/llm", handlers.LLMHandler)

	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

