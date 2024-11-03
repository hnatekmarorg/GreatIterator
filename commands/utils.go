package commands

import (
	"github.com/spf13/cobra"
	"github.com/tmc/langchaingo/llms/openai"
)

// GetOpenAIClient creates openai client based on information from supplied command
func GetOpenAIClient(cmd *cobra.Command) (*openai.LLM, error) {
	url, err := cmd.Flags().GetString("openai-url")
	if err != nil {
		panic(err)
	}
	token, err := cmd.Flags().GetString("openai-token")
	if err != nil {
		panic(err)
	}
	baseModel, err := cmd.Flags().GetString("base-model")
	if err != nil {
		panic(err)
	}
	llm, err := openai.New(openai.WithBaseURL(url), openai.WithToken(token), openai.WithModel(baseModel))
	return llm, err
}
