package pkg

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func GenerateSplitCommitsFolder(parsedFlags ParsedFlags) (string, string) {
	tmpCommitFilePath, tmpDir := generateTmpFilePath()

	// Generate a file with all the commits
	generateAllCommits(parsedFlags, tmpCommitFilePath)

	// Split this file into multiple files with commits
	// Each one of those file will be picked up by a goroutine
	// to be processed in parallel
	generateSplitCommits(tmpCommitFilePath, tmpDir)

	// Delete original commits file
	err := os.Remove(tmpCommitFilePath)
	if err != nil {
		fmt.Println("Error deleting file:", err)
	}

	return tmpCommitFilePath, tmpDir
}

func generateSplitCommits(tmpCommitFilePath string, tmpDir string) {
	// Read tmpCommitFilePath
	file, err := os.Open(tmpCommitFilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	count := 0
	fileIndex := 0

	for scanner.Scan() {
		line := scanner.Text()

		file, err := getWritefile(fmt.Sprintf("part_%d", fileIndex), fmt.Sprintf("%s/blamed_commits", tmpDir))
		if err != nil || file == nil {
			fmt.Println("Error creating file:", err)
			continue
		}

		_, err = file.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}

		count++

		if count == 10 {
			count = 0
			fileIndex++
		}
	}

}

func generateAllCommits(parsedFlags ParsedFlags, tmpFilePath string) {
	var command string = "git log  --pretty=format:%H"

	if parsedFlags.from != "" {
		command += " " + parsedFlags.from
		if parsedFlags.to != "" {
			command += ".." + parsedFlags.to
		} else {
			command += "..HEAD"
		}
	}

	if parsedFlags.since != "" {
		command += " --since=" + parsedFlags.since
		if parsedFlags.until != "" {
			command += " --until=" + parsedFlags.until
		} else {
			command += " --until=now"
		}
	}

	if parsedFlags.author != "" {
		command += " --author=" + parsedFlags.author
	}

	command += " > " + tmpFilePath
	fmt.Println(command)
	exec.Command("sh", "-c", command).Run()
}

func generateTmpFilePath() (string, string) {
	tmpDir := os.TempDir() + "/blamed/" + GenerateRandomHash()
	tmpDirCommit := tmpDir + "/blamed_commits"
	err := os.MkdirAll(tmpDirCommit, 0755)

	if err != nil {
		panic(err)
	}

	filepath := tmpDirCommit + "/blamed_commits"
	return filepath, tmpDir
}
