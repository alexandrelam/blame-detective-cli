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
	var commitHash string = ""
	var author string = ""

	// Iterate over each line in the file
	for scanner.Scan() {
		line := scanner.Text()

		if hasCommitHash(line) {
			commitHash = getCommitHash(line)
		}

		if hasAuthor(line) {
			author = getAuthor(line)
		}

		if hasFilePath(line) {
			filePath = getCommitFilePath(line)
		}

		if filePath != "" {
			// Write to file
			file, err := getWritefile(filePath, tmpDir)
			if err != nil || file == nil {
				fmt.Println("Error creating file:", err)
				continue
			}

			var lineToWrite string = line

			if isDiffTitle(line) && commitHash != "" && author != "" {
				lineToWrite = line + " | " + commitHash + " | " + author
			}

			_, err = file.WriteString(lineToWrite + "\n")
			if err != nil {
				fmt.Println("Error writing line to file:", err)
			}
		}
	}

	// Delete whole.diff
	err = os.Remove(tmpFilePath)
	if err != nil {
		fmt.Println("Error deleting whole.diff:", err)
	}

	// Check for any errors encountered during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
}

func isDiffTitle(line string) bool {
	if len(line) < 3 {
		return false
	}
	return line[:3] == "+++"
}

func hasAuthor(line string) bool {
	if len(line) < 6 {
		return false
	}
	return line[:6] == "Author"
}

func getAuthor(line string) string {
	parts := strings.Split(line, "Author: ")
	return parts[1]
}

func hasFilePath(line string) bool {
	if len(line) < 4 {
		return false
	}
	return line[:4] == "diff"
}

func getCommitHash(line string) string {
	if len(line) < 6 {
		return ""
	}

	parts := strings.Split(line, "commit")
	return parts[1]
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
