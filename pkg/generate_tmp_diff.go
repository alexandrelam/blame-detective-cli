package pkg

import (
	"fmt"
	"os"
	"os/exec"
)

func GenerateTmpDiff(parsedFlags ParsedFlags) string {
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

	tmpFilePath := generateTmpFilePath()
	command += " | while read -r commit_hash; do git show \"$commit_hash\"; done > " + tmpFilePath
	fmt.Println(command)
	exec.Command("sh", "-c", command).Run()

	return tmpFilePath
}

func generateTmpFilePath() string {
	tmpDir := os.TempDir() + "/blamed/" + GenerateRandomHash()
	err := os.MkdirAll(tmpDir, 0755)

	if err != nil {
		panic(err)
	}

	filepath := tmpDir + "/whole.diff"
	return filepath
}
