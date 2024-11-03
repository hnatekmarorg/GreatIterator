package commands

import (
	"testing"
)

func Test_checkFiles(t *testing.T) {
	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("Calling app without arguments should be succesful")
	}
}

func Test_checkInvalidConfig(t *testing.T) {
	rootCmd.SetArgs([]string{
		"fix",
		"--openai-url",
		"https://clearly-invalid.url.foo.com",
		"-t", "invalid-token",
		"a",
		"b",
	})
	err := rootCmd.Execute()
	if err == nil {
		t.Errorf("Calling app with invalid config should fail")
	}
}
