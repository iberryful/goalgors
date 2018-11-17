package rbt

type NodeColor int

const (
	Red   NodeColor = 1
	Black NodeColor = 0
)

type RBTreeNode struct {
	Key    int
	Color  NodeColor
	Left   *RBTreeNode
	Right  *RBTreeNode
	Parent *RBTreeNode
}

type RBTree struct {
	Root *RBTreeNode
}

func (tree *RBTree) Insert(key int) *RBTreeNode {
	root := tree.Root
	node := &RBTreeNode{
		Key:   key,
		Color: Red,
	}

	for root != nil {
		node.Parent = root
		switch {
		case key < root.Key:
			root = root.Left
		case key > root.Key:
			root = root.Right
		default:
			return root
		}
	}

	//empty tree
	if node.Parent == nil {
		node.Color = Black
		tree.Root = node
		return node
	}

	switch {
	case key < node.Parent.Key:
		node.Parent.Left = node
	case key > node.Parent.Key:
		node.Parent.Right = node
	}

	return InsertFix(tree, node)
}

func InsertFix(tree *RBTree, node *RBTreeNode) *RBTreeNode {
	for node.Parent != nil && node.Parent.Color == Red {
		if node.uncle() != nil && node.uncle().Color == Red {
			node.Parent.Color = Black
			node.uncle().Color = Black
			node.grandParent().Color = Red
			node = node.grandParent()
		} else {
			if node.Parent == node.grandParent().Left {
				if node == node.Parent.Right {
					node = node.Parent
					leftRotate(tree, node)
				} else {
					node.Parent.Color = Black
					node.grandParent().Color = Red
					rightRotate(tree, node.grandParent())
				}
			} else {
				if node == node.Parent.Left {
					node = node.Parent
					rightRotate(tree, node)
				} else {
					node.Parent.Color = Black
					node.grandParent().Color = Red
					leftRotate(tree, node.grandParent())
				}
			}
		}
	}
	tree.Root.Color = Black
	return node
}

func (node *RBTreeNode) grandParent() *RBTreeNode {
	return node.Parent.Parent
}

func (node *RBTreeNode) uncle() *RBTreeNode {
	grandParent := node.grandParent()
	if grandParent.Left == node.Parent {
		return grandParent.Right
	} else {
		return grandParent.Left
	}
}

func leftRotate(tree *RBTree, node *RBTreeNode) {
	newRoot := node.Right
	node.Right = newRoot.Left
	if newRoot.Left != nil {
		newRoot.Left.Parent = node
	}
	if node.Parent == nil {
		tree.Root = newRoot
	} else {
		if node == node.Parent.Left {
			node.Parent.Left = newRoot
		} else {
			node.Parent.Right = newRoot
		}
	}
	newRoot.Parent = node.Parent
	newRoot.Left = node
	node.Parent = newRoot
}

func rightRotate(tree *RBTree, node *RBTreeNode) {
	newRoot := node.Left
	node.Left = newRoot.Right
	if newRoot.Right != nil {
		newRoot.Right.Parent = node
	}

	if node.Parent == nil {
		tree.Root = newRoot
	} else {
		if node == node.Parent.Left {
			node.Parent.Left = newRoot
		} else {
			node.Parent.Right = newRoot
		}
	}
	newRoot.Parent = node.Parent
	newRoot.Right = node
	node.Parent = newRoot
}
