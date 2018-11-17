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

func Search(root *BSTreeNode, key int) (*BSTreeNode, bool) {
	for root != nil && root.Key != key {
		if key < root.Key {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return root, root != nil
}

func (tree *BSTree) Search(key int) (*BSTreeNode, bool) {
	return Search(tree.Root, key)
}

func Min(root *BSTreeNode) (*BSTreeNode) {
	if root == nil {
		return nil
	}

	for root.Left != nil {
		root = root.Left
	}
	return root
}

func (tree *BSTree) Min() (*BSTreeNode) {
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

func (tree *BSTree) Max() (*BSTreeNode) {
	return Max(tree.Root)
}

func (root *BSTreeNode) Successor() (*BSTreeNode) {
	if root.Right != nil {
		return Min(root.Right)
	}

	p := root.Parent
	for p != nil && root == p.Right {
		root = p
		p = p.Parent
	}

	return p
}

func (root *BSTreeNode) Predecessor() (*BSTreeNode) {
	if root.Left != nil {
		return Max(root.Left)
	}

	p := root.Parent
	for p != nil && root == p.Left {
		root = p
		p = p.Parent
	}

	return p
}

func (tree *BSTree) Delete(key int) {
	node, ok := tree.Search(key)
	if !ok {
		return
	}

	tree.Root = Delete(tree.Root, node)
}

func Delete(root *BSTreeNode, node *BSTreeNode) (*BSTreeNode) {
	if node == nil {
		return root
	}

	if node.Key < root.Key {
		root.Left = Delete(root.Left, node)
		return root
	}

	if node.Key > root.Key {
		root.Right = Delete(root.Right, node)
		return root
	}

	if node.Left == nil && node.Right == nil {
		return nil
	}

	if node.Right != nil {
		k := Min(node.Right)
		node.Key = k.Key
		node.Right = Delete(node.Right, k)
		return root
	}

	root.Key = node.Left.Key
	root.Right = node.Left.Right
	root.Left = node.Left.Left

	if root.Left != nil {
		node.Left.Parent = root
	}
	if root.Right != nil {
		node.Right.Parent = root
	}

	return root
}

func Traverse(root *BSTreeNode) []int {
	var keys []int
	if root == nil {
		return keys
	}
	keys = append(keys, root.Key)
	keys = append(keys, Traverse(root.Left)...)
	keys = append(keys, Traverse(root.Right)...)
	return keys
}

func (tree *BSTree) Traverse() []int {
	return Traverse(tree.Root)
}
