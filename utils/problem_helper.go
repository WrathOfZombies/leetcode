package utils

import (
	"io/fs"
	"os"
	"path"
	"strings"
	"unicode"
)

// Defining some constants for the problem space
const problems_dir = "./problems"
const solution_file = "solution.go"
const validation_file = "validation.yml"

// Structure definition for what a problem looks like
// (TIP: In golang anything captilizied is exported from the package)
type problem struct {
	Name       string
	Path       string
	Title      string
	Solution   string
	Validation string
}

// Helper to make the problem structure entry
func constructProblem(file fs.FileInfo) problem {
	var name = file.Name()
	var basePath = path.Join(problems_dir, name)
	var prettified = prettify(name)

	return problem{
		Name:       name,
		Title:      prettified,
		Path:       basePath,
		Solution:   path.Join(basePath, solution_file),
		Validation: path.Join(basePath, validation_file),
	}
}

// Prettifies the filename by split hyphens and converting them to more human readable format
// eg. 1-two-sum => 1. Two Sum
func prettify(fileName string) string {
	// Using string builder for performance
	var prettified strings.Builder

	// Use this bit to track if we need to capitalize
	captialize := false
	// Use this bit to identify when the hyphen following numbers shows up
	firstHypen := true

	// (TIP: A rune is a UTF-8 character in the string sequence)
	for _, rune := range fileName {
		if captialize {
			prettified.WriteRune(unicode.ToUpper(rune))
			captialize = false
		} else {
			// When we find an hyphen, mark captialize as true
			// If this is the first hyphen, then insert a '.'
			if rune == '-' {
				captialize = true
				if firstHypen {
					prettified.WriteRune('.')
					firstHypen = false
				}
				prettified.WriteRune(' ')
			} else {
				prettified.WriteRune(rune)
			}
		}
	}

	// Generate the string from the buffer
	return prettified.String()
}

// Helper to read the directory and load the problems
func LoadProblems() []problem {
	results := []problem{}

	// (TIP: Go doesn't have generics and hence this was a better approach instead of interface {}, which is similar to any)
	IterateFiles(problems_dir, func(file os.FileInfo) {
		results = append(results, constructProblem(file))
	})

	return results
}
