package main

import (
	"fmt"
	"os/exec"

	"github.com/alexandrelam/blame-detective-cli/pkg"
)

func main() {
	fmt.Println("Welcome to Blame Detective CLI!")

	// We parse the flags
	var parsedFlags = pkg.ParseFlags()

	// This generates a tmp file with all the diffs
	var tmpFilePath, tmpDir = pkg.GenerateTmpDiff(parsedFlags)

	// We parse this file to reconstruct the folder structure
	pkg.GenerateDiffFolder(tmpFilePath, tmpDir)

	// Open vscode with the folder
	exec.Command("code", tmpDir).Run()

	// We print the path of the tmp file
	fmt.Println(tmpFilePath)
}
