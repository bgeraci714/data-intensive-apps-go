package bstree

// Comparer interface for the tree nodes
// type Comparer interface {
// 	Compare(other Node) int
// }

// Node are nodes of base BST
type Node struct {
	key   string
	value interface{}
	left  *Node
	right *Node
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
