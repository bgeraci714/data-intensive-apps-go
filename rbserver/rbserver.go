package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/bgeraci714/rbtree"
)

// Entry represents the expected shape of the request body for put calls
type Entry struct {
	Key   string
	Value string
}

// Key represents the expected shape of the reuqest body for get calls
type Key struct {
	Key string
}

// little test function to verify server is working correctly
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, world.\n")
}

// put takes a pointer to a red black tree and returns
// a callback that adds the key value pair to the tree
func put(tree *rbtree.RBTree) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		var e Entry

		// try to decode the body, could probably be refactored
		err := json.NewDecoder(req.Body).Decode(&e)
		if err != nil || e.Key == "" {
			log.Printf("Error decoding body: %v", err)
			http.Error(w, "Error occurred when trying to read body.", http.StatusBadRequest)
			return
		}

		tree.Insert(e.Key, e.Value)

		fmt.Fprintf(w, "{%s: %s}\n", e.Key, e.Value)

		// return results and print state of tree (purely for testing purposes initially)
		fmt.Printf("tree state:\n" + tree.ToString())
	}
}

// get takes a pointer to a red black tree and returns
// a callback retrieves a key if found in the tree
func get(tree *rbtree.RBTree) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		var k Key

		// try to decode the body
		err := json.NewDecoder(req.Body).Decode(&k)
		if err != nil || k.Key == "" {
			log.Printf("Error decoding body: %v", err)
			http.Error(w, "can't decode body", http.StatusBadRequest)
			return
		}

		// try to find the value in the tree, indicate if found in response
		if val, found := tree.Get(k.Key); found {
			fmt.Fprintf(w, "{found: true, value: %s}", val)
			return
		}

		fmt.Fprintf(w, "{found: false}")
	}

}

func main() {

	store := rbtree.RBTree{
		Root: nil,
		Compare: func(a, b interface{}) int {
			return strings.Compare(a.(string), b.(string))
		}}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/put", put(&store))
	http.HandleFunc("/get", get(&store))

	http.ListenAndServe(":8080", nil)
}
