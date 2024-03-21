package pkg

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/schollz/progressbar/v3"
)

func Bar(commitDirPath string) *progressbar.ProgressBar {
	files, err := os.ReadDir(commitDirPath)

	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	numLines := 0

	for _, file := range files {
		filePath := commitDirPath + "/" + file.Name()
		lines, err := getNumberOfLinesInFile(filePath)

		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}

		numLines += lines
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
	r := bufio.NewReader(file)

	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
