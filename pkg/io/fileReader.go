package io

import (
	"fmt"
	"io/fs"
	"strings"
)

func ReadLinesFromFile(f fs.FS, path string) ([]string, error) {
	content, err := fs.ReadFile(f, path)
	if err != nil {
		return nil, fmt.Errorf("read file: %v", err)
	}
	return strings.Split(string(content), "\n"), nil
}
