package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/ovrdos/goforex/internal/llm"
)

type LLMRequest struct {
	Input string `json:"input"`
}

type LLMResponse struct {
	Output string `json:"output"`
}

func LLMHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LLMRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	output := llm.GenerateResponse(req.Input)

	res := LLMResponse{Output: output}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

