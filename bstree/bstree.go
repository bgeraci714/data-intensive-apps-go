package bstree

import (
	"fmt"
)

// BSTree implementation
type BSTree struct {
	Root    *Node
	Compare func(a, b interface{}) int
}

// Insert adds a new key value pair to the tree
func (t *BSTree) Insert(key string, val interface{}) {
	t.Root = insertRec(t.Root, key, val, t.Compare)
}

// Size returns the size of the tree
func (t BSTree) Size() int {
	return sizeRec(t.Root)
}

// Height returns the height of the tree
func (t BSTree) Height() int {
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

func insertRec(n *Node, key string, val interface{}, compare func(a, b interface{}) int) *Node {
	if n == nil {
		return &Node{key, val, nil, nil} // might be an issue if this is allocated on function's stack, be mindful of this
	}

	cmp := compare(key, n.key)
	switch {
	case cmp > 0: // key > n.key
		n.right = insertRec(n.right, key, val, compare)
	case cmp < 0: // key < n.key
		n.left = insertRec(n.left, key, val, compare)
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
func (t *BSTree) Delete(key string) {
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
func (t BSTree) ToString() string {
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
