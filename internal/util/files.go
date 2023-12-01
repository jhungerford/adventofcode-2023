package util

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// ReadInputLines returns lines in the given file.  Checks a few locations for the input file.
func ReadInputLines(filename string) ([]string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get working dir: %w", err)
	}

	// IntelliJ uses the project root as the working directory for commands, and the internal directory
	// where the test is run for tests.  Check a couple of likely locations for input files.
	checkFiles := []string{
		// Project root for commands, test location for tests
		filepath.Join(pwd, filename),
		// Input from the project root
		filepath.Join(pwd, "input", filename),
		// Test working directory is two levels deep.
		filepath.Join(pwd, "..", "..", filename),
		// Input from the test working directory
		filepath.Join(pwd, "..", "..", "input", filename),
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
		return nil, fmt.Errorf("failed to open file: %w\nchecked: %+v", errors.Unwrap(ferr), checkFiles)
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
