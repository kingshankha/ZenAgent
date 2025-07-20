package models

type AzureOpenAIResponse struct {
	ID                 string         `json:"id"`
	Object             string         `json:"object"`
	CreatedAt          int64          `json:"created_at"`
	Status             string         `json:"status"`
	Background         bool           `json:"background"`
	Error              *string        `json:"error"`
	IncompleteDetails  *string        `json:"incomplete_details"`
	Instructions       *string        `json:"instructions"`
	MaxOutputTokens    *int           `json:"max_output_tokens"`
	Model              string         `json:"model"`
	Output             []OutputItem   `json:"output"`
	ParallelToolCalls  bool           `json:"parallel_tool_calls"`
	PreviousResponseID *string        `json:"previous_response_id"`
	Reasoning          Reasoning      `json:"reasoning"`
	ServiceTier        string         `json:"service_tier"`
	Store              bool           `json:"store"`
	Temperature        float64        `json:"temperature"`
	Text               TextInfo       `json:"text"`
	TopP               float64        `json:"top_p"`
	Truncation         string         `json:"truncation"`
	Usage              Usage          `json:"usage"`
	User               *string        `json:"user"`
	Metadata           map[string]any `json:"metadata"`
}

type OutputItem struct {
	ID      string    `json:"id"`
	Type    string    `json:"type"`
	Status  string    `json:"status"`
	Content []Content `json:"content"`
	Role    string    `json:"role"`
}

type Content struct {
	Type        string `json:"type"`
	Annotations []any  `json:"annotations"` // [] for now; adjust if needed
	Text        string `json:"text"`
}

type Reasoning struct {
	Effort  *string `json:"effort"`
	Summary *string `json:"summary"`
}

type TextInfo struct {
	Format TextFormat `json:"format"`
}

type TextFormat struct {
	Type string `json:"type"`
}

type Usage struct {
	InputTokens         int         `json:"input_tokens"`
	InputTokensDetails  TokenDetail `json:"input_tokens_details"`
	OutputTokens        int         `json:"output_tokens"`
	OutputTokensDetails TokenDetail `json:"output_tokens_details"`
	TotalTokens         int         `json:"total_tokens"`
}

type TokenDetail struct {
	CachedTokens    *int `json:"cached_tokens,omitempty"`
	ReasoningTokens *int `json:"reasoning_tokens,omitempty"`
}

// You can use this if the structure of tools is unknown
type Tool map[string]any
