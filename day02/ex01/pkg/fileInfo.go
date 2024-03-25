package fileInfo

import (
	"bufio"
	"fmt"
	"myWc/config"
	"os"
	"time"
)

func ProcessFile(flags config.Flags, filename string) {

	time.Sleep(1000 * time.Millisecond)
	file, err := os.Open(filename)
	if err != nil {
		return
	} else {
		defer file.Close()
		fileScanner := bufio.NewScanner(file)
		switch {
		case flags.Words:
			fileScanner.Split(bufio.ScanWords)
		case flags.Lines:
			fileScanner.Split(bufio.ScanLines)
		case flags.Characters:
			//pass
		}

		count := 0
		for fileScanner.Scan() {
			count++
		}

		fmt.Println(count, "\t", filename)
	}

}
