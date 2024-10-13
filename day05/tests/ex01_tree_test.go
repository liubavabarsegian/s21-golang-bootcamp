package test

import (
	tree "day05/tree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeEasy(t *testing.T) {
	treeTest := &tree.Tree{}
	treeTest.Insert(true)
	treeTest.Root.InsertLeft(true)
	treeTest.Root.InsertRight(false)
	treeTest.Root.Left.InsertLeft(true)
	treeTest.Root.Left.InsertRight(false)
	treeTest.Root.Right.InsertLeft(true)
	treeTest.Root.Right.InsertRight(true)

	// Uncomment the method to display the tree
	// treeTest.Root.Print(0)

	expected := []bool{true, true, false, true, true, false, true}
	assert.Equal(t, expected, treeTest.UnrollGarland())
}

func TestTreeEasy01(t *testing.T) {
	treeTest := &tree.Tree{}
	treeTest.Insert(false)
	treeTest.Root.InsertLeft(false)
	treeTest.Root.InsertRight(true)
	treeTest.Root.Left.InsertLeft(true)
	treeTest.Root.Left.InsertRight(true)
	treeTest.Root.Left.Left.InsertLeft(true)
	treeTest.Root.Left.Left.InsertRight(false)
	treeTest.Root.Left.Right.InsertLeft(true)
	treeTest.Root.Left.Right.InsertRight(false)
	treeTest.Root.Left.Right.Left.InsertRight(true)

	// Uncomment the method to display the tree
	//treeTest.Root.Print(0)

	expected := []bool{false, false, true, true, true, true, false, true, false, true}
	assert.Equal(t, expected, treeTest.UnrollGarland())
}

func TestTreeEmpty(t *testing.T) {
	treeTest := &tree.Tree{}
	// Uncomment the method to display the tree
	//treeTest.Root.Print(0)

	expected := make([]bool, 0)
	assert.Equal(t, expected, treeTest.UnrollGarland())
}
