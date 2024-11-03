/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/charmbracelet/log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "GreatIterator",
	Short: "GreatIterator is program for fixing issues via llm based on test cases",
	Long:  "GreatIterator is program for fixing issues via llm based on test cases",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() {
		debugEnabled, err := rootCmd.Flags().GetBool("debug")
		if err != nil {
			panic(err)
		}
		if debugEnabled {
			log.SetLevel(log.DebugLevel)
		}
	})
	if openaiURL := os.Getenv("OPENAI_URL"); openaiURL != "" {
		rootCmd.PersistentFlags().String("openai-url", openaiURL, "URL for openai")
	} else {
		rootCmd.PersistentFlags().String("openai-url", "https://api.openai.com/v1", "URL for openai")
	}
	if openaiToken := os.Getenv("OPENAI_TOKEN"); openaiToken != "" {
		rootCmd.PersistentFlags().String("openai-token", openaiToken, "URL for openai")
	} else {
		rootCmd.PersistentFlags().String("openai-token", "unknown", "URL for openai")
	}
	rootCmd.PersistentFlags().Bool("debug", false, "Print debug logs. WARNING: This could reveal sensitive information use with caution")
}
