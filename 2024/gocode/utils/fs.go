package utils

import (
	"fmt"
	"os"
)

func ReadFileString(path string) string {
	return string(ReadFileBytes(path))
}

func ReadFileBytes(path string) []byte {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file: ", path)
	}
	return fileBytes
}
