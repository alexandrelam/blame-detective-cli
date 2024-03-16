package main

import (
	"fmt"

	"github.com/alexandrelam/blame-detective-cli/pkg"
)

func main() {
	fmt.Println("Welcome to Blame Detective CLI!")

	// We parse the flags
	var parsedFlags = pkg.ParseFlags()

	// This generates a tmp file with all the diffs
	pkg.GenerateTmpDiff(parsedFlags)

	// We parse this file to reconstruct the folder structure
}
