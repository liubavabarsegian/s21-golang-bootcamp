package main

import (
	"day01/compareFS/config"
	fscomparer "day01/pkg/fs_comparer"
	"fmt"
)

func main() {
	newSnapshotName, oldSnapshotName, err := config.GetSnapshots()
	if err != nil {
		fmt.Println(err)
		return
	}

	fscomparer.Compare(newSnapshotName, oldSnapshotName)
}
