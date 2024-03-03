package app

import (
	"Ex00-Anscombe/internal/anscombe"
	"flag"
	"fmt"
	"strconv"
)

const (
	MAX_NUMBER = 100000
	MIN_NUMBER = -100000
)

type ChosenFlags struct {
	Mean              bool
	Median            bool
	Mode              bool
	StandartDeviation bool
}

func Call() {
	var flags ChosenFlags
	parseFlags(&flags)
	fmt.Println(noFlags(flags))
	if noFlags(flags) {
		chooseAllFlags(&flags)
	}

	numbers := scanNumbers()
	if len(numbers) > 0 {
		printStats(numbers, flags)
	}
}

func parseFlags(flags *ChosenFlags) {
	flag.BoolVar(&flags.Mean, "mean", false, "Mean Flag")
	flag.BoolVar(&flags.Mode, "mode", false, "Mode Flag")
	flag.BoolVar(&flags.Median, "median", false, "Median Flag")
	flag.BoolVar(&flags.StandartDeviation, "sd", false, "Standart Deviation Flag")
}

func noFlags(flags ChosenFlags) bool {
	return !flags.Mean && !flags.Mode && !flags.Median && !flags.StandartDeviation
}

func chooseAllFlags(flags *ChosenFlags) {
	flags.Mean = true
	flags.Mode = true
	flags.StandartDeviation = true
	flags.Median = true
}

func scanNumbers() []int {
	numbers := make([]int, 0)
	var buf string

	fmt.Println("Input numeric values [-100000, 100000] by Enter. Enter empty line to stop inputting args.")
	for {
		fmt.Scanln(&buf)
		if buf == "" {
			break
		}
		num, err := strconv.Atoi(buf)
		if err != nil {
			panic(err)
		}
		if num > MAX_NUMBER || num < MIN_NUMBER {
			fmt.Println("Input number out of range. Please, input numbers between -100000 and 100000")
			break
		}
		numbers = append(numbers, num)
		buf = ""
	}
	return numbers
}

func printStats(numbers []int, flags ChosenFlags) {
	if flags.Mean {
		fmt.Printf("Mean: %.2f\n", anscombe.Mean(numbers))
	}

	if flags.Median {
		fmt.Printf("Median: %.2f\n", anscombe.Median(numbers))
	}

	if flags.Mean {
		fmt.Printf("Mode: %d\n", anscombe.Mode(numbers))
	}

	if flags.Mean {
		fmt.Printf("Standart Deviation: %.2f\n", anscombe.StandartDeviation(numbers))
	}
}
