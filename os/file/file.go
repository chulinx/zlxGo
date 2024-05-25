package file

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

var (
	errorNotFound = errors.New("not found")
)

func ErrorNotFound() error {
	return errorNotFound
}

func GrepFile(filepath, query string) (string, error) {
	file, err := os.Open(filepath)
	defer file.Close()

	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(file)
	for i := 1; scanner.Scan(); i++ {
		if strings.Contains(scanner.Text(), query) {
			return scanner.Text(), nil
		}
	}
	return "", ErrorNotFound()
}
