package llms

type llm interface {
	Generate(prompt string, temperature float32) (response string)
}
