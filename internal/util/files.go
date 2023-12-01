package util

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// ReadFileLines returns lines in the given file relative to the current working directory.
func ReadFileLines(filename string) ([]string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get working dir: %w", err)
	}

	checkFiles := []string{
		// Project root
		filepath.Join(pwd, filename),
		// Test and cmd default working directory is two levels deep.
		filepath.Join(pwd, "..", "..", filename),
	}

	var ferr error
	var file *os.File

	for _, checkFile := range checkFiles {
		file, ferr = os.Open(checkFile)
		if ferr == nil {
			break
		}
	}

	if ferr != nil {
		return nil, fmt.Errorf("failed to open file: %w", ferr)
	}

	defer func(file *os.File) {
		cerr := file.Close()
		if cerr != nil {
			fmt.Printf("failed to close file: %v\n", cerr)
		}
	}(file)

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		return nil, fmt.Errorf("failed to read lines from file: %w", err)
	}

	return lines, nil
}
