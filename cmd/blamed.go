package main

import (
	"fmt"

	"github.com/alexandrelam/blame-detective-cli/pkg"
)

func main() {
	fmt.Println("Welcome to Blame Detective CLI!")

	pkg.GenerateTmpDiff()
}
