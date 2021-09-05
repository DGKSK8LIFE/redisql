package utils

import (
	"fmt"
	"os"
)

func ValidateFilePath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}

	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}

	return nil
}
