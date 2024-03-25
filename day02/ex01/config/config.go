package config

import (
	"errors"
	"flag"
)

type Flags struct {
	Words      bool
	Lines      bool
	Characters bool
}

var ErrFlagsMutualExclusion = errors.New("flags must be mutually exclusive")
var ErrNoFilesGiven = errors.New("no files specified")

func GetArgs() (flags Flags, filenames []string, err error) {
	flags = Flags{}
	flag.BoolVar(&flags.Words, "w", false, "Count words in files")
	flag.BoolVar(&flags.Lines, "l", false, "Count lines in files")
	flag.BoolVar(&flags.Characters, "m", false, "Count characters in files")
	flag.Parse()

	filenames = flag.Args()

	if len(filenames) == 0 {
		err = ErrNoFilesGiven
		return
	}

	if !flags.Lines && !flags.Characters {
		flags.Words = true
	}

	if !flagsMutuallyExclusive(flags) {
		err = ErrFlagsMutualExclusion
		return
	}

	return
}

func flagsMutuallyExclusive(flags Flags) bool {
	return (flags.Lines != flags.Words) != flags.Characters
}
