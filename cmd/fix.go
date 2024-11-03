package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// fixCmd represents the fix command
var fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "Fix action fixes file based on test case and 1 or more files it can change",
	Long:  `Fix action fixes file based on test case and 1 or more files it can change`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(fixCmd)
	if openaiURL := os.Getenv("OPENAI_URL"); openaiURL != "" {
		fixCmd.Flags().String("openai-url", openaiURL, "URL for openai")
	} else {
		fixCmd.Flags().String("openai-url", "https://api.openai.com/v1", "URL for openai")
	}

	if openaiToken := os.Getenv("OPENAI_TOKEN"); openaiToken != "" {
		fixCmd.Flags().String("openai-token", openaiToken, "URL for openai")
	} else {
		fixCmd.Flags().String("openai-token", "unkown", "URL for openai")
	}
}
