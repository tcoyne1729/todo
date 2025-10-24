package storage

import (
	"log"
	"os"
	"path/filepath"
)

func DefaultDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Join(home, ".localtodo")
	os.MkdirAll(dir, 0755)
	return dir
}
