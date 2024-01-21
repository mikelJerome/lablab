package test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test_Output(t *testing.T) {
	// Set the directory you want to document here
	dirToDocument := "D://lab" // Use the path to your project directory

	// Set the Markdown file you want to output to
	outputFile := "directory_structure.md"

	// Create or truncate the output file
	file, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write the header of the Markdown file
	file.WriteString("# Project Directory Structure\n\n")

	// A helper function to write the directory structure recursively
	var writeDir func(writer *os.File, path string, indentLevel int)
	writeDir = func(writer *os.File, path string, indentLevel int) {
		entries, err := os.ReadDir(path)
		if err != nil {
			panic(err)
		}

		for _, entry := range entries {
			// Create indentation
			indent := strings.Repeat("  ", indentLevel)
			if entry.IsDir() {
				// Write directory name in bold
				writer.WriteString(fmt.Sprintf("%s- **%s/**\n", indent, entry.Name()))
				// Recurse into subdirectory
				writeDir(writer, filepath.Join(path, entry.Name()), indentLevel+1)
			} else {
				// Write file name
				writer.WriteString(fmt.Sprintf("%s- `%s`\n", indent, entry.Name()))
			}
		}
	}

	// Start the recursive directory writing process
	writeDir(file, dirToDocument, 0)

	fmt.Printf("Directory structure written to %s\n", outputFile)
}
