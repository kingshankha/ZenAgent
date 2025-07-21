package llms

import "github.com/kingshankha/ZenAgent/prompts/templates"

type llm interface {
	Generate(prompt templates.AzureOpenAIPrompt, temperature float32) (err error, ai_response string)
}
