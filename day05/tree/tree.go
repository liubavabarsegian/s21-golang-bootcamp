package tree

import (
	"container/list"
	"fmt"
	"strings"
)

type Tree struct {
	Root *TreeNode
}

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func (t *Tree) Insert(value bool) {
	if t.Root == nil {
		t.Root = &TreeNode{HasToy: value, Left: nil, Right: nil}
	}
}

func (t *Tree) AreToysBalanced() bool {
	if t.Root == nil {
		return false
	}

	return t.Root.Left.CountToys() == t.Root.Right.CountToys()
}

func (t *Tree) UnrollGarland() []bool {
	if t.Root == nil {
		return make([]bool, 0)
	}

	return t.Root.GetBoolValues()
}

func (n *TreeNode) GetBoolValues() []bool {
	var results []bool
	queue := list.New()
	isOdd := true

	queue.PushBack(n)
	for queue.Len() > 0 {
		size := queue.Len()
		levelNodes := make([]bool, size)

		for i := 0; i < size; i++ {
			nodeElement := queue.Front()
			queue.Remove(nodeElement)

			// при извлечении достается элемент типа *list.Element, нужно привести к нужному типу
			node := nodeElement.Value.(*TreeNode)

			if isOdd {
				levelNodes[size-1-i] = node.HasToy
			} else {
				levelNodes[i] = node.HasToy
			}

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		results = append(results, levelNodes...)
		isOdd = !isOdd
	}

	return results
}

func (n *TreeNode) Print(identation int) {
	if n == nil {
		return
	}

	if n.Right != nil {
		n.Right.Print(identation + 2)
	}

	if n.HasToy == true {
		fmt.Println(strings.Repeat("  ", identation), 1)
	} else if n.HasToy == false {
		fmt.Println(strings.Repeat("  ", identation), 0)
	}

	if n.Left != nil {
		n.Left.Print(identation + 2)
	}
}

func (n *TreeNode) Insert(value bool, side string) {
	if n == nil {
		return
	}

	switch side {
	case "left":
		n.InsertLeft(value)
	case "right":
		n.InsertRight(value)
	}
}

func (n *TreeNode) InsertLeft(value bool) {
	if n.Left == nil {
		n.Left = &TreeNode{HasToy: value}
	}
}

func (n *TreeNode) InsertRight(value bool) {
	if n.Right == nil {
		n.Right = &TreeNode{HasToy: value}
	}
}

func (n *TreeNode) CountToys() (sum int) {
	sum = 0
	if n.HasToy == true {
		sum += 1
	}

	if n.Left != nil {
		sum += n.Left.CountToys()
	}
	if n.Right != nil {
		sum += n.Right.CountToys()
	}

	return
}

// EX00 MAIN
// func main() {
// 	treeTest := &Tree{}
// 	treeTest.Insert(false)
// 	treeTest.Root.InsertLeft(false)
// 	treeTest.Root.InsertRight(true)
// 	treeTest.Root.Left.InsertLeft(false)
// 	treeTest.Root.Left.InsertRight(true)

// 	treeTest.Root.Print(0)
// 	fmt.Println(treeTest.Root.Right.CountToys())
// }

// EX01 MAIN
// func main() {
// 	//                    	  1
// 	//                    	/  \
// 	//                    1     0
// 	//                    / \   / \
// 	//                   1   0 1   1

// 	//  [true, true, false, true, true, false, true]

// 	treeTest := &Tree{}
// 	treeTest.Insert(true)
// 	treeTest.Root.InsertLeft(true)
// 	treeTest.Root.InsertRight(false)
// 	treeTest.Root.Left.InsertLeft(true)
// 	treeTest.Root.Left.InsertRight(false)
// 	treeTest.Root.Right.InsertLeft(true)
// 	treeTest.Root.Right.InsertRight(true)

// 	treeTest.Root.Print(0)
// 	fmt.Println(treeTest.UnrollGarland())
// }
