package utils

import "log"

// A helper to check if the error is valid
func IsError(err error) {
	if err == nil {
		return
	}

	// TODO: Pretty print the error
	log.Fatal(err)
	panic(err)
}
