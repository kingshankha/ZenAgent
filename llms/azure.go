package llms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/kingshankha/ZenAgent/models"
	"github.com/kingshankha/ZenAgent/prompts/templates"

	"github.com/go-playground/validator/v10"
)

type AzureOpenAI struct {
	endpoint      string `validate:"required"`
	api_key       string `validate:"required"`
	deployment_id string `validate:"required"`
	api_version   string `validate:"required"`
}

// Generate sends a prompt to the Azure OpenAI API using the provided configuration and temperature.
// It validates the AzureOpenAI struct fields, constructs the API request, and handles the response.
// On success, it returns the AI-generated response as a string. On failure, it returns an error and an empty string.
// Parameters:
//   - prompt: The prompt to send to the Azure OpenAI API, of type templates.AzureOpenAIPrompt.
//   - temperature: The sampling temperature for the model's output.
//
// Returns:
//   - err: An error if the request fails or the response cannot be parsed; otherwise, nil.
//   - ai_response: The AI-generated response as a string if successful; otherwise, an empty string.
func (llm_cfg AzureOpenAI) Generate(prompt templates.AzureOpenAIPrompt, temperature float32) (err error, ai_response string) {

	validator := validator.New(validator.WithPrivateFieldValidation())

	validation_err := validator.Struct(llm_cfg)

	if validation_err != nil {
		log.Printf("Missing Required LLM Configs from the azure openai llm instance \n err : %s", validation_err)
		return validation_err, ""

	}

	llm_cfg.endpoint = strings.TrimSuffix(llm_cfg.endpoint, "/")

	chat_completion_path := fmt.Sprintf("/openai/deployments/%s/chat/completions?api-version=%s", llm_cfg.deployment_id, llm_cfg.api_version)

	chat_completion_endpoint := llm_cfg.endpoint + chat_completion_path

	api_req := models.AzureOpenAIAPIReqPayload{
		AzureOpenAIPrompt: prompt,
		Temperature:       temperature,
	}

	json_payload, m_err := json.Marshal(api_req)

	if m_err != nil {
		log.Printf("error while encoding the azure api api request \n ,error = %s", m_err)
		return m_err, ""
	}

	req, err := http.NewRequest("POST", chat_completion_endpoint, bytes.NewBuffer(json_payload))

	if err != nil {
		log.Printf("error while creating a new request for the openai call")
		return err, ""
	}

	req.Header.Set("Content-Type", "application/json")

	req.Header.Set("api-key", string(llm_cfg.api_key))

	log.Printf("Calling Azure OpenAI API %s", time.Local)

	res, err := http.DefaultClient.Do(req)

	log.Printf("Response received for  Azure OpenAI API %s", time.Local)

	if err != nil {
		log.Printf("error while calling azure openai api \n err: %s", err)
		return err, ""
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	reader := strings.NewReader(string(body))

	json_decoder := json.NewDecoder(reader)

	if res.StatusCode == http.StatusOK {
		var parsedResp models.AzureOpenAIRespPayload
		decodeErr := json_decoder.Decode(&parsedResp)
		if decodeErr != nil {
			log.Printf("Error parsing successful Azure OpenAI response (possible schema change): %v", decodeErr)
			return decodeErr, ""
		}

		return nil, parsedResp.Choices[0].Message.Content
	} else {

		log.Print(string(body))

		var errResp models.AzureOpenAIErrRespPayload

		decodeErr := json_decoder.Decode(&errResp)
		if decodeErr != nil {
			log.Fatalf("Error parsing error response from Azure OpenAI (possible schema change): %v", decodeErr)
			return decodeErr, ""
		}

		log.Printf("Azure OpenAI API Call failed with status code: %d\nResponse body: %s", res.StatusCode, body)

		return fmt.Errorf("OpenAI API Call Error : Resp Code %d", res.StatusCode), ""
	}

}

func NewAzureOpenAI(endpoint string, api_key string, deployment_id string, api_version string) llm {

	var client llm = AzureOpenAI{
		endpoint:      endpoint,
		api_key:       api_key,
		deployment_id: deployment_id,
		api_version:   api_version,
	}

	return client
}
