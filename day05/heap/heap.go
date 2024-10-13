package heap

import (
	"container/heap"
	"errors"
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
