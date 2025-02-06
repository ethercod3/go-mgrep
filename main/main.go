package main

import (
	"bufio"
	"fmt"
	"mgrep/utils"
	"os"
	"sync"
)

func main() {

	stdin := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter substring to search : ")
	subStr, _ := stdin.ReadString('\n')
	fmt.Printf("Enter dirpath : ")
	dirPath, _ := stdin.ReadString('\n')

	files := make([]string, 0)
	utils.FetchFiles(utils.CleanStr(dirPath), &files)

	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go utils.FindEntry(file, utils.CleanStr(subStr), &wg)
	}

	wg.Wait()

}
