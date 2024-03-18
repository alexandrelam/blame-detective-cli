package pkg

import (
	"fmt"
	"os"
)

func GetWritefile(filepath string, tmpDir string) (*os.File, error) {
	newFile := tmpDir + "/" + filepath

	// Check if directory exists
	directory := getDirFromAbsolutepath(newFile)
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		// Create the directory
		err := os.MkdirAll(directory, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return nil, err
		}
	}

	// Check if the file exists
	if _, err := os.Stat(newFile); os.IsNotExist(err) {
		// Create the file
		_, err := os.Create(newFile)
		if err != nil {
			return nil, err
		}
	}

	// Otherwise return the file
	file, _ := os.OpenFile(newFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	return file, nil
}
