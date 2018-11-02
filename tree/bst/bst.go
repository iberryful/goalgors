package bst

type BSTreeNode struct {
	Key    int
	Left   *BSTreeNode
	Right  *BSTreeNode
	Parent *BSTreeNode
}

type BSTree struct {
	Root *BSTreeNode
}

func (tree *BSTree) Insert(key int) {
	node := Insert(tree.Root, key)
	if tree.Root == nil {
		tree.Root = node
	}
}

func Insert(root *BSTreeNode, key int) *BSTreeNode {
	node := &BSTreeNode{
		Key: key,
	}

	for root != nil {
		node.Parent = root
		switch {
		case key < root.Key:
			root = root.Left
		case key > root.Key:
			root = root.Right
		default:
			return node
		}
	}

	//empty tree
	if node.Parent == nil {
		return node
	}

	switch {
	case key < node.Parent.Key:
		node.Parent.Left = node
	case key > node.Parent.Key:
		node.Parent.Right = node
	}

	return node
}

func Search(root *BSTreeNode, key int) (*BSTreeNode, bool){
	for root != nil && root.Key != key {
		if key < root.Key{
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root, root != nil
}

func (tree *BSTree)Search(key int) (*BSTreeNode, bool) {
	return Search(tree.Root, key)
}

func Min(root *BSTreeNode) (*BSTreeNode) {

	for root.Left != nil {
		root = root.Left
	}
	return root
}

func (tree *BSTree)Min() (*BSTreeNode) {
	return Min(tree.Root)
}

func Max(root *BSTreeNode) (*BSTreeNode) {
	if root == nil {
		return nil
	}

	for root.Right != nil {
		root = root.Right
	}
	return root
}

func (tree *BSTree)Max() (*BSTreeNode) {
	return Max(tree.Root)
}


