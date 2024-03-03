package anscombe

import (
	"math"
	"sort"
)

func Mean(numbers []int) float64 {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return float64(sum) / float64(len(numbers))
}

func Median(numbers []int) float64 {
	sort.Ints(numbers)
	if len(numbers)%2 == 1 {
		return float64(numbers[len(numbers)/2])
	} else {
		return (float64(numbers[len(numbers)/2]) + float64(numbers[len(numbers)/2-1])) / 2
	}
}

func Mode(numbers []int) int {
	numbers_occurences := make(map[int]int)
	for _, value := range numbers {
		if val, ok := numbers_occurences[value]; ok {
			numbers_occurences[value] = val + 1
		} else {
			numbers_occurences[value] = 1
		}
	}

	occurrence := numbers_occurences[numbers[0]]
	mode := numbers[0]
	for key, value := range numbers_occurences {
		if occurrence < value {
			occurrence = value
			mode = key
		}
	}
	return mode
}

func StandartDeviation(numbers []int) float64 {
	return math.Sqrt(variance(numbers))
}

func variance(numbers []int) float64 {
	var total float64
	for _, value := range numbers {
		total += math.Pow(float64(value)-Mean(numbers), 2)
	}
	return total / float64(len(numbers))
}
