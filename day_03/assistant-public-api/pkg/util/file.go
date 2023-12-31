package util

import (
	"fmt"
	"path"
	"path/filepath"
)

func ExtractFilePath(filePath string) (string, string, string, error) {
	if filePath == "" {
		return "", "", "", fmt.Errorf("File path is empty")
	}

	folder := path.Dir(filePath)
	ext := filepath.Ext(filePath)
	if len(ext) > 0 && ext[0] == '.' {
		ext = ext[1:]
	}
	fileName := path.Base(filePath)[0:(len(path.Base(filePath)) - len(filepath.Ext(filePath)))]

	return folder, fileName, ext, nil
}
