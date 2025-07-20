package llms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/kingshankha/ZenAgent/models"
)

type AzureOpenAI struct {
	endpoint      string
	api_key       string
	deployment_id string
	api_version   string
}

type AzureOpenAIAPIReqPayload struct {
	Prompt      string  `json:"model"`
	Temperature float32 `json:"temperature"`
}

type AzureOpenAIAPIRespPayload struct {
}

func (llm_cfg AzureOpenAI) Generate(prompt string, temperature float32) string {

	chat_completion_path := fmt.Sprintf("/openai/deployments/%s/completions?api-version=%s", llm_cfg.deployment_id, llm_cfg.api_version)

	chat_completion_endpoint := llm_cfg.endpoint + chat_completion_path

	api_req := AzureOpenAIAPIReqPayload{
		Prompt:      prompt,
		Temperature: temperature,
	}

	josn_payload, m_err := json.Marshal(api_req)

	if m_err != nil {
		log.Fatalf("error while encoding the azure api api request \n ,error = %s", m_err)
	}

	req, err := http.NewRequest("POST", chat_completion_endpoint, bytes.NewBuffer(josn_payload))

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.Header.Set("api-key", string(llm_cfg.api_key))

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Panicf("error while calling azure openai api \n err: %s", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	reader := strings.NewReader(string(body))

	json_decoder := json.NewDecoder(reader)

	var parsed_resp models.AzureOpenAIResponse

	err = json_decoder.Decode(&parsed_resp)

	if err != nil {
		log.Fatalf("error while parsing the azurre ponai response check for change in schema in azure openai api resp \n%s", err)
	}

	ai_resp := parsed_resp.Output[0].Content[0].Text

	return ai_resp

}

func NewAzureOpenAI(endpoint string, api_key string, deployment_id string) llm {

	var client llm = AzureOpenAI{
		endpoint:      endpoint,
		api_key:       api_key,
		deployment_id: deployment_id,
	}

	return client
}
