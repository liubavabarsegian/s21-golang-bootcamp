package config

import (
	"errors"
	"flag"
	"os"
)

type ChosenFlags struct {
	OnlySymlinks    bool
	OnlyDirs        bool
	OnlyFiles       bool
	OnlySpecificExt bool
}

type Args struct {
	DirPath   string
	Extension string
}

var ErrWrongFlagsCombination = errors.New("flag -ext works only with flag -f")
var ErrNoSuchDirectory = errors.New("no such directory")
var ErrNoDirPassed = errors.New("no directory passed")

func GetArgs() (arguments Args, flags ChosenFlags, err error) {
	// Init flags and args
	flags = ChosenFlags{}
	arguments = Args{}

	flag.BoolVar(&flags.OnlySymlinks, "sl", false, "Print only symlinks")
	flag.BoolVar(&flags.OnlyDirs, "d", false, "Print only directories")
	flag.BoolVar(&flags.OnlyFiles, "f", false, "Print only files")
	flag.StringVar(&arguments.Extension, "ext", "", "Print only specific extension. Works only with -f")
	flag.Parse()

	if arguments.Extension != "" {
		flags.OnlySpecificExt = true
	}

	if flags.OnlySpecificExt && !flags.OnlyFiles {
		err = ErrWrongFlagsCombination
		return
	}

	if len(flag.Args()) > 0 {
		arguments.DirPath = flag.Arg(0)
	}

	if arguments.DirPath == "" {
		err = ErrNoDirPassed
		return
	}

	if _, err = os.Stat(arguments.DirPath); os.IsNotExist(err) {
		err = ErrNoSuchDirectory
		return
	}

	return
}
