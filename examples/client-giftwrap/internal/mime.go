package internal

import (
	"net/http"
	"path/filepath"
	"strings"
)

var mimeExtensions = map[string]string{
	".txt":  "text/plain",
	".json": "application/json",
	".xml":  "application/xml",
	".html": "text/html",
	".css":  "text/css",
	".js":   "application/javascript",
	".png":  "image/png",
	".jpg":  "image/jpeg",
	".jpeg": "image/jpeg",
	".gif":  "image/gif",
	".pdf":  "application/pdf",
	".zip":  "application/zip",
	".tar":  "application/x-tar",
	".gz":   "application/gzip",
}

func DetectMimeType(filename string, data []byte) string {
	ext := strings.ToLower(filepath.Ext(filename))
	if mime, ok := mimeExtensions[ext]; ok {
		return mime
	}
	
	mime := http.DetectContentType(data)
	if mime != "" {
		return mime
	}
	
	return "application/octet-stream"
}
