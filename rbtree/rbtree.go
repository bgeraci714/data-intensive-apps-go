package rbtree

import (
	"fmt"
)

// RBTree implementation
type RBTree struct {
	Root    *Node
	Compare func(a, b interface{}) int
}

// Get returns the value found for a given key, using an a boolean indicator if found
func (t *RBTree) Get(key string) (interface{}, bool) {
	node, found := t.GetNode(key)
	if found {
		return node.value, true
	}
	return nil, false
}

// GetNode returns the node found for a given key, using a boolean indicator if found
func (t *RBTree) GetNode(key string) (*Node, bool) {
	return getNodeRec(t.Root, key, t.Compare)
}

func getNodeRec(n *Node, key string, compare func(a, b interface{}) int) (*Node, bool) {
	// node not found
	if n == nil {
		return nil, false
	}

	cmp := compare(key, n.key)
	switch {
	case cmp > 0: // key > n.key
		return getNodeRec(n.right, key, compare)
	case cmp < 0: // key < n.key
		return getNodeRec(n.left, key, compare)
	default: // node was found
		return n, true
	}
}

func recolor(nodes []*Node) {
	for _, n := range nodes {
		n.color = !n.color
	}
}

// Translated from https://www.cs.auckland.ac.nz/software/AlgAnim/red_black.html
func leftRotate(t *RBTree, x *Node) {
	y := x.right

	// turn y's left sub-tree into x's right sub-tree
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}

	// y's new parent was x's parent
	y.parent = x.parent

	// set the parent to point to y instead of x
	// need to check first if we're at the root
	if x.parent == nil {
		t.Root = y
	} else if x == x.parent.left {
		// x is on left of its parent
		x.parent.left = y
	} else { // otherwise x must have been on the right
		x.parent.right = y
	}

	// put x on y's left
	y.left = x
	x.parent = y
}

func rightRotate(t *RBTree, x *Node) {
	y := x.left

	// turn y's left sub-tree into x's left sub-tree
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}

	// y's new parent was x's parent
	y.parent = x.parent

	// set the parent to point to y instead of x
	// need to check first if we're at the root
	if x.parent == nil {
		t.Root = y
	} else if x == x.parent.right {
		// x is on right of its parent
		x.parent.right = y
	} else { // otherwise x must have been on the left
		x.parent.left = y
	}

	// put x on y's right
	y.right = x
	x.parent = y
}

// Insert inserts with balanced algorithm involved
// Translated from https://www.cs.auckland.ac.nz/software/AlgAnim/red_black.html
func (t *RBTree) Insert(key string, val interface{}) {
	// Perform tree insert for tree T and node n
	t.insert(key, val)
	n, _ := t.GetNode(key) // will be optimized later

	n.color = Red
	for n != t.Root && n.parent.color == Red {
		if n.parent == n.LeftUncle() {
			// if n's parent is a left, the uncle is x's right uncle
			uncle := n.RightUncle()
			if uncle != nil && uncle.color == Red {
				// Do case 1: change the colors
				n.parent.color = Black
				uncle.color = Black
				n.Grandparent().color = Red

				// move n up the tree
				n = n.Grandparent()
			} else {
				// uncle is a black node
				if n == n.parent.right {
					// and n is to the right
					// case 2 - move x up and rotate
					n = n.parent
					leftRotate(t, n)
				}
				// case 3
				n.parent.color = Black
				n.Grandparent().color = Red
				rightRotate(t, n.Grandparent())
			}
		} else {
			// if n's parent is a right, the uncle is x's left uncle
			uncle := n.LeftUncle()
			if uncle != nil && uncle.color == Red {
				// Do case 1: change the colors
				n.parent.color = Black
				uncle.color = Black
				n.Grandparent().color = Red

				// move n up the tree
				n = n.Grandparent()
			} else {
				// uncle is a black node
				if n == n.parent.left {
					// and n is to the left
					// case 2 - move x up and rotate
					n = n.parent
					rightRotate(t, n)
				}
				// case 3
				n.parent.color = Black
				n.Grandparent().color = Red
				leftRotate(t, n.Grandparent())
			}
		}
	}
	t.Root.color = Black
}

// Insert adds a new key value pair to the tree
func (t *RBTree) insert(key string, val interface{}) {
	t.Root = insertRec(t.Root, key, val, t.Compare, nil)
}

// Size returns the size of the tree
func (t RBTree) Size() int {
	return sizeRec(t.Root)
}

// Height returns the height of the tree
func (t RBTree) Height() int {
	return height(t.Root)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func height(n *Node) int {
	if n == nil {
		return -1
	}
	return 1 + max(height(n.left), height(n.right))
}

func sizeRec(n *Node) int {
	if n == nil {
		return 0
	}
	return 1 + sizeRec(n.left) + sizeRec(n.right)
}

func insertRec(n *Node, key string, val interface{}, compare func(a, b interface{}) int, parent *Node) *Node {
	if n == nil {
		return &Node{key, val, nil, nil, false, parent} // might be an issue if this is allocated on function's stack, be mindful of this
	}

	cmp := compare(key, n.key)
	switch {
	case cmp > 0: // key > n.key
		n.right = insertRec(n.right, key, val, compare, n)
	case cmp < 0: // key < n.key
		n.left = insertRec(n.left, key, val, compare, n)
	default: // key == n.key
		n.value = val // overwrite old value if the keys match
	}
	return n
}

// PrintInorder prints out all the nodes of a subtree inorder
func PrintInorder(n *Node) {
	if n != nil {
		PrintInorder(n.left)
		fmt.Print(n.key)
		PrintInorder(n.right)
	}
}

// Delete deletes item with the matching key
func (t *RBTree) Delete(key string) {
	t.Root = delete(t.Root, key, t.Compare)
}

func delete(n *Node, key string, compare func(a, b interface{}) int) *Node {
	if n == nil {
		return nil
	}

	if cmp := compare(key, n.key); cmp > 0 { // key > n.key
		n.right = delete(n.right, key, compare)
	} else if cmp < 0 { // key < n.key
		n.left = delete(n.left, key, compare)
	} else { // key == n.key
		if n.right == nil { // if no right child
			return n.left
		} else if n.left == nil { // if there's a right but no left
			return n.right
		}

		// both right and left child
		tmp := n                       // copy over node n
		n = min(tmp.right)             // swap n for its min on the right
		n.right = deleteMin(tmp.right) // replace the min that was just copied
		n.left = tmp.left              // copy over original left node
	}
	return n
}

func min(n *Node) *Node {
	if n.left == nil {
		return n
	}
	return min(n.left)
}

func deleteMin(n *Node) *Node {
	if n.left == nil {
		return n.right
	}
	n.left = deleteMin(n.left)
	return n
}

// ToString prints out a dash spaced version of the tree
func (t RBTree) ToString() string {
	return printSubtree(t.Root, 0)
}

func printSubtree(n *Node, h int) string {
	if n == nil {
		return ""
	}
	s := ""
	for i := 0; i < h; i++ {
		s += "-"
	}
	s += n.key + "\n"
	s += printSubtree(n.left, h+1)
	s += printSubtree(n.right, h+1)
	return s
}
