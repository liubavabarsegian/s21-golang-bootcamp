package main

import (
	"log"
	"myRotate/config"
	archiver "myRotate/pkg"
	"sync"
	"time"
)

func main() {
	directory, filenames, err := config.GetArgs()
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	for _, filename := range filenames {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()

			time.Sleep(100 * time.Millisecond)
			archiver.CompressToTarGz(filename, directory)
		}(filename)
	}
	wg.Wait()
}
