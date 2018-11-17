package rbt

type NodeColor int

const (
	Red   Color = 0
	Black Color = 1
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
