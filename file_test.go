package go_file_output

import (
	"testing"
	"fmt"
	"io/ioutil"
	"os"
)

var variants = []struct {
	path string
	text string
}{
	{
		"test1.txt",
		"string1",
	},
	{
		"test2.text",
		"string \n 2\n",
	},
}

func TestFileOutput_Write(t *testing.T) {

	for _, variant := range variants {
		checkFile(variant.path, variant.text, variant.text, t)
		checkFile(variant.path, variant.text, variant.text+variant.text, t)
		deleteFile(variant.path, t)
	}
}
func checkFile(path, text, expected string, t *testing.T) {
	n, err := FileOutput(path).Write([]byte(text))
	if err != nil {
		t.Error(fmt.Sprintf("FileOutput(\"%s\").Write(\"%s\"): error expected nil, actual \"%s\"", path, text, err))
	}
	if n != len(text) {
		t.Error(fmt.Sprintf("FileOutput(\"%s\").Write(\"%s\"): n expected %d, actual %d", path, text, len(text), n))
	}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		t.Error(err)
	}
	if expected != string(file) {
		t.Error(fmt.Sprintf("File %s: expected \"%s\", actual \"%s\"", path, text, string(file)))
	}
}
func deleteFile(path string, t *testing.T) {
	err := os.Remove(path)
	if err != nil {
		t.Error(err)
	}
}
