package rbtree

// Color is a typedef for booleans
type Color bool

// Red is a basic coloring constant equaling true
const Red = true

// Black is a basic coloring constant equaling false
const Black = false

// Node are nodes of base BST
type Node struct {
	key    string
	value  interface{}
	left   *Node
	right  *Node
	color  Color
	parent *Node
}

// Compare compares the keys of the two nodes
// returns 1 if n's is greater than other's
// returns 0 if n's key is equal to other's
// returns -1 if n's key is less than other's
func (n Node) Compare(other Node) int {
	if n.key > other.key {
		return 1
	} else if n.key < other.key {
		return -1
	}
	return 0
}

// Grandparent returns the grandparent of the current node
func (n *Node) Grandparent() *Node {
	if n.parent == nil || n.parent.parent == nil {
		panic("There is no grandparent!")
	}
	return n.parent.parent
}

// LeftUncle goes to find the right uncle and panics if not
func (n *Node) LeftUncle() *Node {
	if n.parent.parent == nil {
		panic("There is no left uncle because there is no grandparent!")
	}
	return n.parent.parent.left
}

// RightUncle goes to find the right uncle and panics if not
func (n *Node) RightUncle() *Node {
	if n.parent.parent == nil {
		panic("There is no right uncle because there is no grandparent!")
	}
	return n.parent.parent.right
}
