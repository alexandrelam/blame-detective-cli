package main

import (
	"fmt"

	"github.com/alexandrelam/blame-detective-cli/pkg"
)

func main() {
	fmt.Println("Welcome to Blame Detective CLI!")
	// startTime := time.Now()

	// We parse the flags
	var parsedFlags = pkg.ParseFlags()

	// This generates a tmp file with all the diffs
	var tmpFilePath, tmpDir = pkg.GenerateTmpDiff(parsedFlags)

	fmt.Println(tmpFilePath, tmpDir)

	// bar := pkg.Bar(tmpFilePath)

	// // We parse this file to reconstruct the folder structure
	// pkg.GenerateDiffFolder(tmpFilePath, tmpDir, bar)

	// // Open vscode with the folder
	// exec.Command("code", tmpDir).Run()

	// // Size of the folder
	// output, _ := exec.Command("du", "-sh", tmpDir).Output()
	// fmt.Println("\nFolder size: " + string(output))

	// // Measure time taken in seconds
	// elapsedTime := time.Since(startTime).Seconds()
	// fmt.Printf("Execution time: %.2f seconds\n", elapsedTime)
	// fmt.Println("Done!")
}
