package main

import (
	"fmt"
	"strings"

	"github.com/bgeraci714/rbtree"
)

func main() {
	tree := rbtree.RBTree{ // bstree.BSTree{
		Root: nil,
		Compare: func(a, b interface{}) int {
			return strings.Compare(a.(string), b.(string))
		}}

	tree.Insert("f", "1")
	tree.Insert("g", "2")
	tree.Insert("h", "3")
	tree.Insert("i", "4")
	tree.Insert("j", "5")
	tree.Insert("k", "6")
	tree.Insert("l", "6")
	tree.Insert("m", "6")
	tree.Insert("n", "6")

	// tree.Delete("c")
	s := tree.ToString()
	fmt.Println(s)
}
