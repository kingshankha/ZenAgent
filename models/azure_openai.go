package models

import "github.com/kingshankha/ZenAgent/prompts/templates"

type AzureOpenAIErrRespPayload struct {
	Error APIError `json:"error"`
}

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type AzureOpenAIAPIReqPayload struct {
	templates.AzureOpenAIPrompt
	Temperature float32 `json:"temperature"`
}

type AzureOpenAIRespPayload struct {
	Choices             []Choice             `json:"choices"`
	Created             int64                `json:"created"`
	ID                  string               `json:"id"`
	Model               string               `json:"model"`
	Object              string               `json:"object"`
	PromptFilterResults []PromptFilterResult `json:"prompt_filter_results"`
	SystemFingerprint   string               `json:"system_fingerprint"`
	Usage               Usage                `json:"usage"`
}

type Choice struct {
	ContentFilterResults ContentFilterResults `json:"content_filter_results"`
	FinishReason         string               `json:"finish_reason"`
	Index                int                  `json:"index"`
	Logprobs             any                  `json:"logprobs"` // nullable
	Message              Message              `json:"message"`
}

type ContentFilterResults struct {
	Hate                  SeverityFilter   `json:"hate"`
	ProtectedMaterialCode DetectionFilter  `json:"protected_material_code,omitempty"`
	ProtectedMaterialText DetectionFilter  `json:"protected_material_text,omitempty"`
	SelfHarm              SeverityFilter   `json:"self_harm"`
	Sexual                SeverityFilter   `json:"sexual"`
	Violence              SeverityFilter   `json:"violence"`
	CustomBlocklists      *FilteredFlag    `json:"custom_blocklists,omitempty"`
	Jailbreak             *DetectionFilter `json:"jailbreak,omitempty"`
	Profanity             *DetectionFilter `json:"profanity,omitempty"`
}

type SeverityFilter struct {
	Filtered bool   `json:"filtered"`
	Severity string `json:"severity"`
}

type DetectionFilter struct {
	Filtered bool `json:"filtered"`
	Detected bool `json:"detected"`
}

type FilteredFlag struct {
	Filtered bool `json:"filtered"`
}

type Message struct {
	Annotations []any  `json:"annotations"` // could be refined if structure is known
	Content     string `json:"content"`
	Refusal     any    `json:"refusal"` // nullable
	Role        string `json:"role"`
}

type PromptFilterResult struct {
	PromptIndex          int                  `json:"prompt_index"`
	ContentFilterResults ContentFilterResults `json:"content_filter_results"`
}

type Usage struct {
	CompletionTokens        int                    `json:"completion_tokens"`
	CompletionTokensDetails CompletionTokensDetail `json:"completion_tokens_details"`
	PromptTokens            int                    `json:"prompt_tokens"`
	PromptTokensDetails     PromptTokensDetail     `json:"prompt_tokens_details"`
	TotalTokens             int                    `json:"total_tokens"`
}

type CompletionTokensDetail struct {
	AcceptedPredictionTokens int `json:"accepted_prediction_tokens"`
	AudioTokens              int `json:"audio_tokens"`
	ReasoningTokens          int `json:"reasoning_tokens"`
	RejectedPredictionTokens int `json:"rejected_prediction_tokens"`
}

type PromptTokensDetail struct {
	AudioTokens  int `json:"audio_tokens"`
	CachedTokens int `json:"cached_tokens"`
}
