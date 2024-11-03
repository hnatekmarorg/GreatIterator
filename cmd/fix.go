package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

// fixCmd represents the fix command
var fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "Fix action fixes file based on test case and 1 or more files it can change",
	Long:  `Fix action fixes file based on test case and 1 or more files it can change`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("Running in debug mode")
	},
}

func init() {
	rootCmd.AddCommand(fixCmd)
	fixCmd.Flags().String("prompt-file", "", "Customize prompt default is embedded inside of application")
}
