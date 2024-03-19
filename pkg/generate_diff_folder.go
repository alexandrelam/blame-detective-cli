package pkg

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func GenerateDiffFolder(commitFolderPath string) {
	// Iterate over each file in the folder
	files, err := os.ReadDir(commitFolderPath)

	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		generateDiffFolderForCommit(commitFolderPath+"/"+file.Name(), commitFolderPath)
	}
}

func generateDiffFolderForCommit(filePath string, directory string) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		commitHash := scanner.Text()

		diffs, err := getDiffs(commitHash)
		if err != nil {
			fmt.Println("Error getting diffs:", err)
			continue
		}

		fmt.Println(diffs)
	}
}

func getDiffs(commitHash string) (string, error) {
	output, err := exec.Command("git", "show", commitHash).Output()

	if err != nil {
		return "", err
	}

	return string(output), nil
}

// func generateDiffFolderForFile(tmpFilePath string, tmpDir string) {
// 	// Open the file
// 	file, err := os.Open(tmpFilePath)
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	// Create a scanner to read the file line by line
// 	scanner := bufio.NewScanner(file)

// 	// Increase buffer size to 2mb
// 	buf := []byte{}
// 	scanner.Buffer(buf, 2048*1024)

// 	var filePath string = ""
// 	var commitHash string = ""
// 	var author string = ""
// 	var date string = ""

// 	// Iterate over each line in the file
// 	for scanner.Scan() {
// 		line := scanner.Text()

// 		if hasCommitHash(line) {
// 			commitHash = getCommitHash(line)
// 		}

// 		if hasAuthor(line) {
// 			author = getAuthor(line)
// 		}

// 		if hasFilePath(line) {
// 			filePath = getCommitFilePath(line)
// 		}

// 		if hasDate(line) {
// 			date = getDate(line)
// 		}

// 		if filePath != "" {
// 			// Write to file
// 			file, err := getWritefile(filePath, tmpDir)
// 			if err != nil || file == nil {
// 				fmt.Println("Error creating file:", err)
// 				continue
// 			}

// 			var lineToWrite string = line

// 			if isDiffTitle(line) && commitHash != "" && author != "" && date != "" {
// 				lineToWrite = line[:3] + " " + author + " | " + commitHash + " | " + date
// 			}

// 			_, err = file.WriteString(lineToWrite + "\n")
// 			if err != nil {
// 				fmt.Println("Error writing line to file:", err)
// 			}
// 		}
// 	}

// 	// Delete whole.diff
// 	err = os.Remove(tmpFilePath)
// 	if err != nil {
// 		fmt.Println("Error deleting whole.diff:", err)
// 	}

// 	// Check for any errors encountered during scanning
// 	if err := scanner.Err(); err != nil {
// 		fmt.Println("Error scanning file:", err)
// 	}
// }

// func hasDate(line string) bool {
// 	if len(line) < 5 {
// 		return false
// 	}
// 	return line[:5] == "Date:"
// }

// func getDate(line string) string {
// 	parts := strings.Split(line, "Date:")
// 	return strings.TrimSpace(parts[1])
// }

// func isDiffTitle(line string) bool {
// 	if len(line) < 3 {
// 		return false
// 	}
// 	return line[:3] == "+++" || line[:3] == "---"
// }

// func hasAuthor(line string) bool {
// 	if len(line) < 6 {
// 		return false
// 	}
// 	return line[:6] == "Author"
// }

// func getAuthor(line string) string {
// 	parts := strings.Split(line, "Author: ")
// 	return parts[1]
// }

// func hasFilePath(line string) bool {
// 	if len(line) < 4 {
// 		return false
// 	}
// 	return line[:4] == "diff"
// }

// func getCommitHash(line string) string {
// 	if len(line) < 6 {
// 		return ""
// 	}

// 	parts := strings.Split(line, "commit")
// 	return parts[1]
// }

// func hasCommitHash(line string) bool {
// 	if len(line) < 6 {
// 		return false
// 	}
// 	return line[:6] == "commit"
// }

// func getCommitFilePath(line string) string {
// 	parts := strings.Split(line, " ")
// 	if len(parts) == 4 {
// 		fromPath := strings.TrimPrefix(parts[2], "a/")
// 		return strings.TrimSpace(fromPath) + ".diff"
// 	}
// 	return ""
// }
