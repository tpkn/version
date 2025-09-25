package utils

import (
	"os"
)

// FileExists returns true is file exists
func FileExists(file_path string) bool {
	i, err := os.Stat(file_path)
	if os.IsNotExist(err) || i.IsDir() {
		return false
	}
	return true
}

// ReadFileString reads file content
func ReadFileString(file_path string) (string, error) {
	data, err := os.ReadFile(file_path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// WriteFile writes to a file
func WriteFile(file_path, content string) error {
	f, err := os.Create(file_path)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.WriteString(content); err != nil {
		return err
	}

	return nil
}
