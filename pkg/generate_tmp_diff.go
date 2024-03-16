package pkg

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func GenerateTmpDiff() {
	var parsedFlags = parseFlags()

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
}

type parsedFlags struct {
	from   string
	to     string
	since  string
	until  string
	author string
}

func parseFlags() parsedFlags {
	var from string
	var to string
	var since string
	var until string
	var author string

	flag.StringVar(&from, "from", "", "The hash of the commit to start from")
	flag.StringVar(&to, "to", "", "The hash of the commit to end at")
	flag.StringVar(&since, "since", "", "The date to start from")
	flag.StringVar(&until, "until", "", "The date to end at")
	flag.StringVar(&author, "author", "", "The author to filter by")

	flag.Parse()

	if from == "" && since == "" {
		panic("You must provide either a from hash or a since date")
	}

	return parsedFlags{
		from:   from,
		to:     to,
		since:  since,
		until:  until,
		author: author,
	}
}

func generateTmpFilePath() string {
	tmpDir := os.TempDir() + "/blamed/" + GenerateRandomHash()
	err := os.MkdirAll(tmpDir, 0755)

	if err != nil {
		panic(err)
	}

	filepath := tmpDir + "/diff.txt"
	return filepath
}
