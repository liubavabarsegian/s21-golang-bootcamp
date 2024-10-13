package heap

import (
	"container/heap"
	"errors"
	"math"
	"reflect"
)

type Present struct {
	Value int
	Size  int
}

type PresentHeap []Present

func (p PresentHeap) Len() int { return len(p) }

func (p PresentHeap) Less(i, j int) bool {
	if p[i].Value == p[j].Value {
		return p[i].Size < p[j].Size
	}

	return p[i].Value > p[j].Value
}

func (p PresentHeap) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PresentHeap) Push(x interface{}) {
	*p = append(*p, x.(Present)) // Указатель на Present
}

func (p *PresentHeap) Pop() interface{} {
	old := *p
	oldSize := old.Len()
	new := old[oldSize-1]
	*p = old[:oldSize-1]
	return new
}

func (p *PresentHeap) GetNCoolestPresents(n int) ([]Present, error) {
	if n < 0 || n > p.Len() {
		return nil, errors.New("invalid number of presents")
	}

	// Инициализируем кучу
	h := &PresentHeap{}
	heap.Init(h)

	// Добавляем все подарки в кучу
	for _, present := range *p {
		heap.Push(h, present)
	}

	// Извлекаем n самых крутых подарков
	coolestPresents := make([]Present, 0, n)
	for i := 0; i < n; i++ {
		coolestPresents = append(coolestPresents, heap.Pop(h).(Present))
	}

	return coolestPresents, nil
}

func (p PresentHeap) GrabPresents(maxSize int) (profitPresents PresentHeap) {
	if maxSize <= 0 {
		return PresentHeap{}
	}

	weightsLen := p.Len()
	table := make([][]int, weightsLen+1)
	for i := range table {
		table[i] = make([]int, maxSize+1)
	}

	for i := 0; i <= weightsLen; i++ {
		for j := 0; j <= maxSize; j++ {
			if i == 0 || j == 0 {
				table[i][j] = 0
			} else {
				if p[i-1].Size > j {
					table[i][j] = table[i-1][j]
				} else {
					table[i][j] = max(table[i-1][j], table[i-1][j-p[i-1].Size]+p[i-1].Value)
				}
			}
		}
	}

	tracePresents(table, weightsLen, maxSize, &p, &profitPresents)
	return profitPresents
}

func tracePresents(table [][]int, i, j int, presents, profitPresents *PresentHeap) {
	if table[i][j] == 0 {
		return
	}

	if table[i-1][j] == table[i][j] {
		tracePresents(table, i-1, j, presents, profitPresents)
	} else {
		tracePresents(table, i-1, j-(*presents)[i-1].Size, presents, profitPresents)
		*profitPresents = append(*profitPresents, (*presents)[i-1])
	}
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func AreEqual(a, b PresentHeap) bool {
	if len(a) != len(b) {
		return false
	}

	// Create maps to count occurrences of each Present
	countMapA := make(map[Present]int)
	countMapB := make(map[Present]int)

	for _, present := range a {
		countMapA[present]++
	}

	for _, present := range b {
		countMapB[present]++
	}

	return reflect.DeepEqual(countMapA, countMapB)
}

// func main() {
// 	presents := PresentHeap{
// 		{Value: 5, Size: 1},
// 		{Value: 4, Size: 3},
// 		{Value: 3, Size: 2},
// 		{Value: 6, Size: 4},
// 	}

// 	capacity := 5
// 	result := presents.GrabPresents(capacity)

// 	fmt.Println("Выбранные подарки:")
// 	for _, p := range result {
// 		fmt.Printf("Value: %d, Size: %d\n", p.Value, p.Size)
// 	}
// }
