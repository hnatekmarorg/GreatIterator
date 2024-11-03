package commands

import (
	"context"
	"errors"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/hnatekmarorg/GreatIterator/ai"
	"github.com/spf13/cobra"
	"github.com/tmc/langchaingo/llms"
	"os"
	"os/exec"
	"strings"
)

// checkFiles takes []string and returns true if all paths exists on filesystem
func checkFiles(files []string) bool {
	// For each file
	for _, fileName := range files {
		if fileInfo, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) || fileInfo.IsDir() {
			log.Errorf("File `%s` does not exist or is a directory", fileName)
			return false
		}
	}
	return true
}

// fixCmd represents the fix command
var fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "Fix action fixes file based on test case and 1 or more files it can change",
	Long:  `Fix action fixes file based on test case and 1 or more files it can change`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("fix requires at least two arguments <test_cmd> <files...>")
		}

		log.Debugf("Will use test command %s", args[0])

		if !checkFiles(args[1:]) {
			return fmt.Errorf("one or more files do not exist")
		}

		log.Debug("Initializing openai client")
		llm, err := GetOpenAIClient(cmd)
		if err != nil {
			return fmt.Errorf("openai client initialization failed. %s", err)
		}

		for {
			ctx := context.Background()
			cmdArgs := strings.Fields(args[0])
			testRun := exec.CommandContext(ctx, cmdArgs[0], cmdArgs[1:]...)

			testOut, err := testRun.CombinedOutput()

			if err == nil {
				log.Print("Testcase passing so it has been fixed!")
				return nil
			}

			fixPrompt := ai.GenerateProposedFixesPrompt(args[0], string(testOut), args[1:])

			completion, err := llm.Call(ctx, fixPrompt, llms.WithTemperature(0.0))
			if err != nil {
				return err
			}
			changeRequest, err := ai.ParseChangesRequest(completion)
			if err != nil {
				return err
			}
			changeRequest.Apply()
			if err != nil {
				return fmt.Errorf("invoking llm failed with %s", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(fixCmd)
	fixCmd.Flags().String("prompt-file", "", "Customize prompt default is embedded inside of application")
	fixCmd.Flags().Bool("dry-run", false, "do not modify files")
}
