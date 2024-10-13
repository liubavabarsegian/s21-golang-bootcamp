package test

import (
	hp "day05/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKnapsackEasy(t *testing.T) {
	items := hp.PresentHeap{
		{Value: 5, Size: 3},
		{Value: 3, Size: 2},
		{Value: 4, Size: 1},
	}

	expected := hp.PresentHeap{{Value: 4, Size: 1}, {Value: 5, Size: 3}}
	result := hp.AreEqual(expected, items.GrabPresents(5))
	assert.Equal(t, result, true)
}

func TestKnapsackZeroCapacity(t *testing.T) {
	items := hp.PresentHeap{
		{Value: 5, Size: 3},
		{Value: 3, Size: 2},
		{Value: 4, Size: 1},
	}

	expected := hp.PresentHeap{}
	assert.Equal(t, expected, items.GrabPresents(0))
}

func TestKnapsackNoItem(t *testing.T) {
	items := hp.PresentHeap{}

	expected := hp.PresentHeap{}
	assert.Equal(t, expected, items.GrabPresents(0))
}
