package utils

import (
	"log"
	"os"
	"path"
	"sync"
)

func FetchFiles(dirPath string, files *[]string, mutex *sync.Mutex) {

	entries, err := os.ReadDir(dirPath)

	if err != nil {
		log.Panicln(err)
	}

	for _, entry := range entries {

		if entry.IsDir() {
			FetchFiles(path.Join(dirPath, entry.Name()), files, mutex)
		} else {
			mutex.Lock()
			*files = append(*files, path.Join(dirPath, entry.Name()))
			mutex.Unlock()
		}

	}

}
