package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GenerateDiffFolder(tmpFilePath string, tmpDir string) {
	// Open the file
	file, err := os.Open(tmpFilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var filePath string = ""
	var reset bool = true
	var tmpAppend string = ""

	// Iterate over each line in the file
	for scanner.Scan() {
		line := scanner.Text()

		if hasCommitHash(line) {
			reset = true
		}

		if hasFilePath(line) {
			filePath = getCommitFilePath(line)
			reset = false
		}

		if reset {
			tmpAppend += line + "\n"
		} else {
			// Write to file
			file, err := getWritefile(filePath, tmpDir)
			if err != nil {
				fmt.Println("Error creating file:", err)
				continue
			}

			if tmpAppend != "" {
				_, err = file.WriteString(tmpAppend)
				if err != nil {
					fmt.Println("Error writing tmpAppend to file:", err)
				}
				tmpAppend = ""
			}

			_, err = file.WriteString(line + "\n")
			if err != nil {
				fmt.Println("Error writing line to file:", err)
			}
		}

	}

	// Check for any errors encountered during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
}

func hasFilePath(line string) bool {
	if len(line) < 4 {
		return false
	}
	return line[:4] == "diff"
}

func hasCommitHash(line string) bool {
	if len(line) < 6 {
		return false
	}
	return line[:6] == "commit"
}

func getCommitFilePath(line string) string {
	parts := strings.Split(line, " ")
	if len(parts) == 4 {
		fromPath := strings.TrimPrefix(parts[2], "a/")
		return strings.TrimSpace(fromPath) + ".diff"
	}
	return ""
}

func getWritefile(filepath string, tmpDir string) (*os.File, error) {
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

func getDirFromAbsolutepath(filepath string) string {
	parts := strings.Split(filepath, "/")
	return strings.Join(parts[:len(parts)-1], "/")
}
