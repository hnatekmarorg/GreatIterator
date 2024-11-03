/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package commands

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
		token, err := rootCmd.Flags().GetString("openai-token")
		if err != nil {
			panic(err)
		}
		url, err := rootCmd.Flags().GetString("openai-url")
		if err != nil {
			panic(err)
		}
		log.Debugf("Using following config OPENAI_URL: %s OPENAI_TOKEN: %s", url, token)
	})
	usage := "URL for openai can be set from OPENAI_URL environment variable"
	if openaiURL := os.Getenv("OPENAI_URL"); openaiURL != "" {
		rootCmd.PersistentFlags().String("openai-url", openaiURL, usage)
	} else {
		rootCmd.PersistentFlags().String("openai-url", "https://api.openai.com/v1", usage)
	}
	usage = "Token for openai can be set from OPENAI_TOKEN environment variable"
	if openaiToken := os.Getenv("OPENAI_TOKEN"); openaiToken != "" {
		rootCmd.PersistentFlags().StringP("openai-token", "t", openaiToken, usage)
	} else {
		rootCmd.PersistentFlags().StringP("openai-token", "t", "unknown", usage)
	}
	rootCmd.PersistentFlags().StringP("base-model", "m", "gpt-4", "Model to use for completions")
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Print debug logs. WARNING: This could reveal sensitive information use with caution")
}
