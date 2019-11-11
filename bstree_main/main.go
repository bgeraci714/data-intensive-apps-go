package main

import (
	"fmt"
	"strings"

	"github.com/bgeraci714/bstree"
)

func main() {
	tree := bstree.BSTree{
		Root: nil,
		Compare: func(a, b interface{}) int {
			return strings.Compare(a.(string), b.(string))
		}}

	tree.Insert("f", "1")
	tree.Insert("c", "2")
	tree.Insert("d", "3")
	tree.Insert("b", "4")
	tree.Insert("a", "5")
	tree.Insert("e", "6")
	tree.Insert("g", "7")
	tree.Insert("g", "8")
	tree.Insert("j", "11")
	tree.Insert("h", "9")
	tree.Insert("i", "10")
	tree.Insert("k", "12")
	tree.Insert("l", "13")
	tree.Insert("m", "14")

	tree.Delete("j")
	s := tree.ToString()
	fmt.Println(s)
}
