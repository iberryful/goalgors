package main

import (
	"fmt"
	"github.com/iberryful/goalgors/tree/bst"
	"github.com/iberryful/goalgors/tree/rbt"
	"math/rand"
)

func assert(c bool) {
	if !c {
		panic("assert failed")
	}
}

func prepareRBTree() (*rbt.RBTree) {
	tree := &rbt.RBTree{}

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

func prepareRndTree(n int) (*rbt.RBTree) {
	tree := &rbt.RBTree{}
	for i := 0; i < n; i++ {
		tree.Insert(rand.Intn(n*100))
	}
	return tree
}

func checkRBTreeNode(node *rbt.RBTreeNode) error {
	if node == nil {
		return nil
	}

	assert(node.Color == rbt.Red || node.Color == rbt.Black)
	if node.Color == rbt.Red {
		assert(node.Left == nil || node.Left.Color == rbt.Black)
		assert(node.Right == nil || node.Right.Color == rbt.Black)
	}

	if node.Left != nil {
		assert(node.Left.Key < node.Key)
		assert(node.Left.Parent == node)
	}

	if node.Right != nil {
		assert(node.Right.Key > node.Key)
		assert(node.Right.Parent == node)
	}

	checkRBTreeNode(node.Right)
	checkRBTreeNode(node.Left)

	return nil
}

func CheckBSTdelete(tree *bst.BSTree) {
	keys := tree.Traverse()
	for _, v := range keys {
		node, ok := tree.Search(v)
		assert(node.Key == v)
		assert(ok == true)
		tree.Delete(v)
		tree.Search(v)
		_, ok = tree.Search(v)
		assert(ok == false)
		//checkBSTreeNode(tree.Root)
	}
}

func checkRBTree(tree *rbt.RBTree) error {
	if tree.Root == nil {
		return nil
	}

	checkRBTreeNode(tree.Root)
	//CheckBSTdelete(tree)

	return nil
}

func main() {
	tree := prepareRBTree()
	//prepareRBTree()
	checkRBTree(tree)

	rndTree := prepareRndTree(100)
	checkRBTree(rndTree)

	fmt.Println("tests passed")

	return
}
