package chat

import (
	"encoding/json"

	"github.com/kingshankha/ZenAgent/llms"
	"github.com/kingshankha/ZenAgent/prompts/templates"

	"net/http"
	"os"
)

var azure_openai_api_key string = os.Getenv("AZURE_OPENAI_API_KEY")

var azure_openai_endpoint string = os.Getenv("AZURE_OPENAI_ENDPOINT")

// ChatReqPayload represents a sample chat message payload.
type ChatReqPayload struct {
	User    string `json:"user" validate:"required min=1 max=100"`
	Message string `json:"message" validate:"required min=1 max=500"`
}

type ChatRespPayload struct {
	AI_Message string `json:"ai_message"`
}

// ChatPostHandler handles POST requests for chat messages.
func ChatPostHandler(w http.ResponseWriter, r *http.Request) {

	var chat_model_deployment string = "gpt-4o"

	var chat_model_temp float32 = 0.15
	var chat_model_api_version string = "2024-06-01"

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var msg ChatReqPayload
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	system_message := templates.Message{
		Role:    "system",
		Content: "You are a helpful assistant. given the user query. Please answer clearly and concisely",
	}

	user_message := templates.Message{
		Role:    "user",
		Content: msg.Message,
	}

	var prompt_messages []templates.Message = []templates.Message{system_message, user_message}

	prompt := templates.AzureOpenAIPrompt{
		Messages: prompt_messages,
	}

	azure_openai_client := llms.NewAzureOpenAI(azure_openai_endpoint, azure_openai_api_key, chat_model_deployment, chat_model_api_version)

	// code needs to be error message , handle properly in  azure_openai

	err, ai_message := azure_openai_client.Generate(prompt, chat_model_temp)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {

		http.Error(w, "Oops ! Something went wrong please try again!!", http.StatusInternalServerError)

	}

	resp_payload := ChatRespPayload{
		AI_Message: ai_message,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp_payload)
}
