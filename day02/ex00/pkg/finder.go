package finder

import (
	"fmt"
	"myFind/config"
	"os"
	"path/filepath"
)

func IterateOverEntities(arguments config.Args, flags config.ChosenFlags) {
	err := filepath.Walk(arguments.DirPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			switch {
			case flags.OnlySymlinks && isSymlink(info):
				printPath(path, info)
			case flags.OnlyDirs && info.IsDir():
				printPath(path, info)
			case flags.OnlySpecificExt && filepath.Ext(path) == ("."+arguments.Extension):
				printPath(path, info)
			case !flags.OnlySpecificExt && flags.OnlyFiles && isFile(info):
				printPath(path, info)
			case noFlags(flags):
				printPath(path, info)
			}
			return nil
		})
	if err != nil {
		fmt.Println(err)
	}
}

func isFile(info os.FileInfo) bool {
	return info.Mode().IsRegular()
}

func isSymlink(info os.FileInfo) bool {
	return info.Mode()&os.ModeSymlink == os.ModeSymlink
}

func noFlags(flags config.ChosenFlags) bool {
	return !flags.OnlyDirs && !flags.OnlyFiles && !flags.OnlySpecificExt && !flags.OnlySymlinks
}

func printPath(path string, info os.FileInfo) {
	if info.Mode()&os.ModeSymlink == os.ModeSymlink {
		targetPath, err := os.Readlink(path)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(path, "->", targetPath)
		}
	} else {
		fmt.Println(path)
	}
}
