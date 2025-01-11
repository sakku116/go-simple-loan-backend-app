package helper

import (
	"mime"
	"path/filepath"
	"strings"
)

func GetMimeType(filename string) string {
	return mime.TypeByExtension(filepath.Ext(filename))
}

func IsImage(filename string) bool {
	lowerFilename := strings.ToLower(filename)

	ext := filepath.Ext(lowerFilename)

	imageExtensions := []string{".jpg", ".jpeg", ".png"}

	for _, imgExt := range imageExtensions {
		if ext == imgExt {
			return true
		}
	}
	return false
}
