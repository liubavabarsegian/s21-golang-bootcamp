package fscomparer

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func Compare(newSnapshotName string, oldSnapshotName string) {
	newFiles := GetSnapshotFiles(newSnapshotName)
	oldFiles := GetSnapshotFiles(oldSnapshotName)

	for _, value := range newFiles {
		if !slices.Contains(oldFiles, value) {
			fmt.Println("ADDED ", value)
		}
	}

	for _, value := range oldFiles {
		if !slices.Contains(newFiles, value) {
			fmt.Println("REMOVED ", value)
		}
	}
}

func GetSnapshotFiles(filename string) (fileNames []string) {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileName := scanner.Text()
		fileNames = append(fileNames, fileName)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading new file:", err)
	}
	return
}
