package pkg

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GenerateDiffFolder(rootDirectory string, commitFolderPath string) {
	// Iterate over each file in the folder
	files, err := os.ReadDir(commitFolderPath)

	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, commitFile := range files {
		if commitFile.IsDir() {
			continue
		}

		generateDiffFolderForCommit(commitFolderPath+"/"+commitFile.Name(), rootDirectory)
	}

	// Delete the commit folder
	err = os.RemoveAll(commitFolderPath)
	if err != nil {
		fmt.Println("Error deleting folder:", err)
	}
}

func generateDiffFolderForCommit(commitFilePath string, directory string) {
	// Open the file
	file, err := os.Open(commitFilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Increase buffer size to 2mb
	buf := []byte{}
	scanner.Buffer(buf, 2048*1024)

	for scanner.Scan() {
		commitHash := scanner.Text()

		diffs, err := getDiffs(commitHash)
		if err != nil {
			fmt.Println("Error getting diffs:", err)
			continue
		}

		var filePath string = ""
		var author string = ""
		var date string = ""

		for _, line := range diffs {
			if hasAuthor(line) {
				author = getAuthor(line)
			}

			if hasFilePath(line) {
				filePath = getCommitFilePath(line)
			}

			if hasDate(line) {
				date = getDate(line)
			}

			if filePath != "" {
				// Write to file
				file, err := getWritefile(filePath, directory)
				if err != nil || file == nil {
					fmt.Println("Error creating file:", err)
					continue
				}

				var lineToWrite string = line

				if isDiffTitle(line) && author != "" && date != "" {
					lineToWrite = line[:3] + " " + author + " | " + commitHash + " | " + date
				}

				_, err = file.WriteString(lineToWrite + "\n")
				if err != nil {
					fmt.Println("Error writing line to file:", err)
				}
			}
		}
	}

	// Check for any errors encountered during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
}

func getDiffs(commitHash string) ([]string, error) {
	output, err := exec.Command("git", "show", commitHash).Output()

	if err != nil {
		return nil, err
	}

	return strings.Split(string(output), "\n"), nil
}

func hasDate(line string) bool {
	if len(line) < 5 {
		return false
	}
	return line[:5] == "Date:"
}

func getDate(line string) string {
	parts := strings.Split(line, "Date:")
	return strings.TrimSpace(parts[1])
}

func isDiffTitle(line string) bool {
	if len(line) < 3 {
		return false
	}
	return line[:3] == "+++" || line[:3] == "---"
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

func getCommitFilePath(line string) string {
	parts := strings.Split(line, " ")
	if len(parts) == 4 {
		fromPath := strings.TrimPrefix(parts[2], "a/")
		return strings.TrimSpace(fromPath) + ".diff"
	}
	return ""
}
