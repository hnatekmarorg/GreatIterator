package cmd

import (
	"context"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/tmc/langchaingo/llms"
)

// fixCmd represents the fix command
var fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "Fix action fixes file based on test case and 1 or more files it can change",
	Long:  `Fix action fixes file based on test case and 1 or more files it can change`,
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Debug("Initializing openai client")
		llm, err := GetOpenAIClient(cmd)
		if err != nil {
			return fmt.Errorf("openai client initialization failed. %s", err)
		}
		ctx := context.Background()
		completion, err := llm.Call(ctx, "Example prompt", llms.WithTemperature(0.0))
		if err != nil {
			return fmt.Errorf("invoking llm failed with %s", err)
		}
		log.Debug(completion)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(fixCmd)
	fixCmd.Flags().String("prompt-file", "", "Customize prompt default is embedded inside of application")
}
