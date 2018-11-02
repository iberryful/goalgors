package main

import (
	"fmt"
	"github.com/iberryful/goalgors/tree/bst"
)

func assert(c bool) {
	if !c {
		panic("assert failed")
	}
}

func prepareBSTree() (*bst.BSTree) {
	tree := &bst.BSTree{}

	tree.Insert(4)
	tree.Insert(8)
	tree.Insert(3)
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(16)
	tree.Insert(10)
	tree.Insert(9)
	tree.Insert(14)

	return tree
}



func checkBSTreeNode(node *bst.BSTreeNode) error {
	if node == nil {
		return nil
	}
	if node.Left != nil {
		assert(node.Left.Key < node.Key)
		assert(node.Left.Parent == node)
		checkBSTreeNode(node.Left)
	}
	if node.Right != nil {
		assert(node.Right.Key > node.Key)
		assert(node.Right.Parent == node)
		checkBSTreeNode(node.Right)
	}
	return nil
}

func checkBSTree(tree *bst.BSTree) error {
	if tree.Root == nil {
		return nil
	}

	checkBSTreeNode(tree.Root)

	return nil
}

func main() {
	tree := prepareBSTree()

	checkBSTree(tree)

	fmt.Println("tests passed")

	s := []int{1,2,3}
	for i,v := range s {
		fmt.Println(i,v)
		s = append(s, i+1)
	}

	return
}
