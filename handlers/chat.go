package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/kingshankha/ZenAgent/llms"
)

var azure_openai_api_key string = os.Getenv("AZURE_OPENAI_API_KEY")

var azure_openai_endpoint string = os.Getenv("AZURE_OPENAI_ENDPOINT")

/*
type ChatRequest struct {
	User    string `json:"user" validate:"required min=1 max=100"`
	Message string `json:"message" validate:"required min=1 max=500"`
}*/

// ChatReqPayload represents a sample chat message payload.
type ChatReqPayload struct {
	User    string `json:"user"`
	Message string `json:"message"`
}

// ChatPostHandler handles POST requests for chat messages.
func ChatPostHandler(w http.ResponseWriter, r *http.Request) {

	var chat_model_deployment string = "gpt-4o"

	var chat_model_temp float32 = 0.15

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var msg ChatReqPayload
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	prompt := fmt.Sprintf("You are a helpful assistant. the users question is: %s. Please answer clearly and concisely",
		msg.Message)

	azure_openai_client := llms.NewAzureOpenAI(azure_openai_endpoint, azure_openai_api_key, chat_model_deployment)

	azure_openai_client.Generate(prompt, chat_model_temp)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(msg)
}
