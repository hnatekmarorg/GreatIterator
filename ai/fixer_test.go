package ai

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func Test_ChangesRequest(t *testing.T) {
	t.Run("applies changes to allowed files", func(t *testing.T) {
		dir, err := ioutil.TempDir("", "fixer_test")
		assert.NoError(t, err)
		defer os.RemoveAll(dir)

		changeRequest := ChangesRequest{
			Changes: []struct {
				Filename string `json:"filename"`
				Content  string `json:"content"`
			}{
				{Filename: filepath.Join(dir, "allowed_file.txt"), Content: "Hello, world!"},
				{Filename: filepath.Join(dir, "another_allowed_file.txt"), Content: "Goodbye, world!"},
			},
		}

		allowedFiles := []string{filepath.Join(dir, "allowed_file.txt"), filepath.Join(dir, "another_allowed_file.txt")}

		changeRequest.Apply(allowedFiles)

		for _, change := range changeRequest.Changes {
			content, err := os.ReadFile(change.Filename)
			assert.NoError(t, err)
			assert.Equal(t, change.Content, string(content))
		}
	})

	t.Run("skips changes to disallowed files", func(t *testing.T) {
		dir, err := os.MkdirTemp("", "fixer_test")
		assert.NoError(t, err)
		defer os.RemoveAll(dir)

		changeRequest := ChangesRequest{
			Changes: []struct {
				Filename string `json:"filename"`
				Content  string `json:"content"`
			}{
				{Filename: filepath.Join(dir, "disallowed_file.txt"), Content: "This file should not be changed"},
			},
		}

		allowedFiles := []string{}

		changeRequest.Apply(allowedFiles)

		_, err = os.ReadFile(filepath.Join(dir, "disallowed_file.txt"))
		assert.Error(t, err)
	})
}
