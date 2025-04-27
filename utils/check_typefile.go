package utils

import (
	"path/filepath"
	"strings"
)

func GetAttachmentTypeFromFilename(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))

	switch ext {
	case ".png", ".jpg", ".jpeg", ".gif":
		return "image"
	case ".mp3", ".wav", ".ogg":
		return "sound"
	case ".mp4", ".avi", ".mov":
		return "video"
	case ".pdf":
		return "document"
	default:
		return "other"
	}
}
