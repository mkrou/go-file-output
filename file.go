package go_file_output

import (
	"os"
)

type FileOutput string

func (f FileOutput) Write(p []byte) (n int, err error) {
	file, err := os.OpenFile(string(f), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return 0, nil
	}
	defer file.Close()
	return file.Write(p)
}
