package fscomparer

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func Compare(newSnapshotName string, oldSnapshotName string) {
	newFiles, err := GetSnapshotFiles(newSnapshotName)
	if err != nil {
		fmt.Println(err)
		return
	}
	oldFiles, err := GetSnapshotFiles(oldSnapshotName)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, value := range newFiles {
		if !slices.Contains(oldFiles, value) {
			fmt.Println("REMOVED ", value)
		}
	}

	for _, value := range oldFiles {
		if !slices.Contains(newFiles, value) {
			fmt.Println("ADDED ", value)
		}
	}
}

func GetSnapshotFiles(filename string) (fileNames []string, err error) {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileName := scanner.Text()
		fileNames = append(fileNames, fileName)
	}
	err = scanner.Err()
	return
}
