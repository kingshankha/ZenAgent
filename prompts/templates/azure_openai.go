package templates

// Message represents a single message in a prompt.
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type AzureOpenAIPrompt struct {
	Messages []Message `json:"messages"`
}
