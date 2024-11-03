package ai

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"github.com/charmbracelet/log"
	"html/template"
	"os"
)

//go:embed prompt.tmpl
var promptTemplate string

/**
    Struct from following json
	{
	  "changes": [{"filename": "<name_of_file>", "content": "<desired_content>"}, ...]
	}
*/

type ChangesRequest struct {
	Changes []struct {
		Filename string `json:"filename"`
		Content  string `json:"content"`
	} `json:"changes"`
}

func (changeRequest ChangesRequest) Apply() {
	for _, change := range changeRequest.Changes {
		err := os.WriteFile(change.Filename, []byte(change.Content), 0644)
		if err != nil {
			log.Error("Failed to write file", "error", err)
		}
	}
	log.Debug("Applied changes to files", "files", changeRequest.Changes)
}

type File struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type Changes struct {
	Files       []File `json:"files"`
	TestCommand string `json:"test_command"`
	TestOutput  string `json:"test_output"`
}

func initChanges(testCommand string, testOutput string, fileList []string) Changes {
	changes := Changes{
		TestCommand: testCommand,
		Files:       make([]File, 0),
		TestOutput:  testOutput,
	}
	for _, filename := range fileList {
		log.Debug("Initial changes", changes.Files)
		content, err := os.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		changes.Files = append(changes.Files, File{Name: filename, Content: string(content)})
	}
	return changes
}

func GenerateProposedFixesPrompt(testCommand string, testOutput string, fileList []string) string {
	// Call test command and get output
	changes := initChanges(testCommand, testOutput, fileList)
	tmpl, err := template.New("prompt").Parse(promptTemplate)
	if err != nil {
		panic(err)
	}

	buffer := &bytes.Buffer{}
	err = tmpl.Execute(buffer, changes)
	if err != nil {
		panic(err)
	}
	proposedFixesPrompt := buffer.String()
	log.Debug("Proposing fixes", proposedFixesPrompt)
	return proposedFixesPrompt
}

func ParseChangesRequest(jsonString string) (ChangesRequest, error) {
	var changes ChangesRequest
	log.Debug("Parsing json: ", jsonString)
	err := json.Unmarshal([]byte(jsonString), &changes)
	if err != nil {
		log.Error("Failed to unmarshal JSON", err)
		return ChangesRequest{}, err
	}
	return changes, nil
}
