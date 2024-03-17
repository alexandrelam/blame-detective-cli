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

	bar := pkg.Bar(tmpFilePath)

	// We parse this file to reconstruct the folder structure
	pkg.GenerateDiffFolder(tmpFilePath, tmpDir, bar)

	// Open vscode with the folder
	exec.Command("code", tmpDir).Run()

	output, _ := exec.Command("du", "-sh", tmpDir).Output()
	fmt.Println("\nFolder size: " + string(output))
	fmt.Println("Done!")
}
