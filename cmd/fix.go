package cmd

import (
	"context"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/tmc/langchaingo/llms"
)

// fixCmd represents the fix command
var fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "Fix action fixes file based on test case and 1 or more files it can change",
	Long:  `Fix action fixes file based on test case and 1 or more files it can change`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("Initializing openai client")
		llm, err := GetOpenAIClient(cmd)
		if err != nil {
			log.Errorf("Openai client initialization failed. %s", err)
		}
		ctx := context.Background()
		completion, err := llm.Call(ctx, "Example prompt", llms.WithTemperature(0.0))
		if err != nil {
			log.Errorf("Invoking openai failed with %s", err)
		}
		log.Debug(completion)
	},
}

func init() {
	rootCmd.AddCommand(fixCmd)
	fixCmd.Flags().String("prompt-file", "", "Customize prompt default is embedded inside of application")
}
