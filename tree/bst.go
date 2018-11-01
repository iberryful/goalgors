package tree

type Comparable interface {
	LessThan(comparable Comparable) bool
	Equal(comparable Comparable) bool
	LargerThan(comparable Comparable) bool
}

type BSTreeNode struct {
	Key    Comparable
	Left   *BSTreeNode
	Right  *BSTreeNode
	Parent *BSTreeNode
}

type BSTree struct {
	Root *BSTreeNode
}

func (tree *BSTree) Insert(key Comparable) {
	if tree == nil {
		node := &BSTreeNode{
			Key: key,
		}
		tree.Root = node
		return
	}
	Insert(tree.Root, key)
}

func Insert(root *BSTreeNode, key Comparable) {
	node := &BSTreeNode{
		Key: key,
	}

	parent := &BSTreeNode{}

	for root != nil {
		parent = root
		if key.LessThan(root.Key) {
			root = root.Left
		} else if key.LargerThan(root.Key) {
			root = root.Right
		} else {
			return
		}
	}

	if key.LessThan(parent.Key) {
		parent.Left = node
	} else if key.LargerThan(parent.Key) {
		parent.Right = node
	}
}
