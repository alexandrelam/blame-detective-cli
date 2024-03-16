package pkg

import (
	"bufio"
	"fmt"
	"os"
)

func GenerateDiffFolder(tmpFilePath string) {
	// Open the file
	file, err := os.Open(tmpFilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Iterate over each line in the file
	for scanner.Scan() {
		line := scanner.Text()
		// Process each line here, for example, you can print it
		fmt.Println(line)
	}

	// Check for any errors encountered during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
}
