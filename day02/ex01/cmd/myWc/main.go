package main

import (
	"fmt"
	"myWc/config"
	fileInfo "myWc/pkg"
	"sync"
)

func main() {
	flags, filenames, err := config.GetArgs()
	if err != nil {
		fmt.Println(err)
	}

	var wg sync.WaitGroup

	for _, value := range filenames {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fileInfo.ProcessFile(flags, value)
		}()
	}
	wg.Wait()

}
