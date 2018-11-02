package main

import (
	"fmt"
	"github.com/iberryful/goalgors/tree/bst"
)

func main() {
	tree := &bst.BSTree{}

	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(8)
	tree.Insert(4)
	tree.Insert(9)
	tree.Insert(7)
	tree.Insert(11)
	tree.Insert(1)
	tree.Insert(12)

	n, _ := tree.Search(11)
	fmt.Println(n)

	max := tree.Max()
	fmt.Println(max.Key == 12)

	return
}
