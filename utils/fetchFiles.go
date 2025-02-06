package utils

import (
	"log"
	"os"
	"path"
)

func FetchFiles(dirPath string, files *[]string) {

	entries, err := os.ReadDir(dirPath)

	if err != nil {
		log.Panicln(err)
	}

	for _, entry := range entries {

		if entry.IsDir() {
			FetchFiles(path.Join(dirPath, entry.Name()), files)
		} else {
			*files = append(*files, path.Join(dirPath, entry.Name()))
		}

	}

}
