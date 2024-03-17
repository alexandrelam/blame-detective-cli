package pkg

import (
	"bufio"
	"fmt"
	"os"

	"github.com/schollz/progressbar/v3"
)

func Bar(tmpFilePath string) *progressbar.ProgressBar {
	// Get the number of lines in the file
	numLines, err := getNumberOfLinesInFile(tmpFilePath)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	// Create a progress bar
	bar := progressbar.Default(int64(numLines))

	return bar
}

func getNumberOfLinesInFile(filepath string) (int, error) {
	// Open the file
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error:", err)
		return 0, err
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Variable to count the lines
	count := 0

	// Read the file line by line and count the lines
	for scanner.Scan() {
		count++
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		return 0, err
	}

	return count, nil
}
