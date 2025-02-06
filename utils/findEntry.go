package utils

import (
	"log"
	"os"
	"strings"
	"sync"
)

func FindEntry(filepath string, subStr string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.ReadFile(filepath)

	if err != nil {
		log.Fatalln(err)
	}

	fileLines := strings.Split(string(file), "\n")

	for fileIdx, line := range fileLines {
		if strings.Contains(line, subStr) {
			log.Printf("%v : line %v\n", filepath, fileIdx+1)
		}
	}

}
