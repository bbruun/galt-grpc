package initializers

import (
	"fmt"
	"log"
)

func GenerateKeys() error {
	path, err := findKeys()
	if err != nil {
		log.Fatalf("no keys could be found for some reason: %v\n", err)
	}
	fmt.Printf("Path to look for keys: %s\n", path)

	return nil
}

func findKeys() (string, error) {
	// Find keys on disk
	return "", nil
}
