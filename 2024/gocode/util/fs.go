package utils

import (
	"fmt"
	"os"
)

func ReadFileString(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file: ", path)
	}
	return string(file)
}
