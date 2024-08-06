package utils_test

import (
	"os"
	"testing"

	"rexai.com/helpers"
)

func TestFileToGithub(t *testing.T) {
	fileContent, _ := os.ReadFile("./file_to_github.go")
	uploader := helpers.Upload{}
	uri, err := uploader.UploadFile(fileContent, "test.go")
	if err != nil {
		t.Error(err)
	}
	t.Log(uri)
}
