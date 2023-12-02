package util

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// ParseInputLines applies the parseLine function to each line in the input file.
func ParseInputLines[T any](filename string, parseLine LineParser[T]) ([]T, error) {
	file, err := findInputFile(filename)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		cerr := file.Close()
		if cerr != nil {
			fmt.Printf("failed to close file: %v\n", cerr)
		}
	}(file)

	var parsedLines []T

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parsedLine, perr := parseLine(line)
		if perr != nil {
			return nil, fmt.Errorf("failed to parse line '%s': %w", line, perr)
		}

		parsedLines = append(parsedLines, parsedLine)
	}

	if scanner.Err() != nil {
		return nil, fmt.Errorf("failed to read lines from file: %w", err)
	}

	return parsedLines, nil
}

// ReadInputLines returns the raw lines in the input file.
func ReadInputLines(filename string) ([]string, error) {
	return ParseInputLines(filename, StringLineParser)
}

// findInputFile locates the given input file, checking a few likely locations.
func findInputFile(filename string) (*os.File, error) {
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

	return file, nil
}

// LineParser parses a generic type from the given line.
type LineParser[T any] func(string) (T, error)

// StringLineParser returns the given line as-is.
func StringLineParser(line string) (string, error) {
	return line, nil
}
