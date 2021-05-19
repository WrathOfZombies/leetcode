package utils

import "os"

// A callback that gets the file that was iterated on
type callbackFn func(file os.FileInfo)

// A helper that iterates the files in the current directly. Doesn't traverse nested directories
func IterateFiles(dir string, callback callbackFn) {
	// Get a pointer to the directory
	folder, err := os.Open(dir)
	IsError(err)

	// Iterate the files
	files, err := folder.Readdir(-1)
	// Remember to dispose the directory pointer
	folder.Close()
	IsError(err)

	// Iterate the files
	// TODO: Can this be lazy?
	for _, file := range files {
		callback(file)
	}
}
