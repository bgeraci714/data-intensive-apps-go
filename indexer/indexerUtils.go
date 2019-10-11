package indexer

import (
	"bufio"
	"fmt"
	"regexp"
	"sync"
)

// Word is the canonical definition for a given word and the document index it was found in
type Word struct {
	word  string
	index int
}

// AddDocumentToIndex adds a document to an index given a single thread
func AddDocumentToIndex(document *bufio.Scanner, documentIndex int, index *map[string][]int, re *regexp.Regexp) {
	for document.Scan() {
		line := document.Text()
		words := re.FindAllString(string(line), -1)
		for _, word := range words {
			seen, ok := (*index)[word]
			if ok && seen[len(seen)-1] != documentIndex {
				(*index)[word] = append(seen, documentIndex)
			} else {
				(*index)[word] = []int{documentIndex}
			}
		}
	}
}

func reader(document *bufio.Scanner, documentIndex int, re *regexp.Regexp, ch chan<- Word, wg *sync.WaitGroup) {
	for document.Scan() {
		line := document.Text()
		words := re.FindAllString(string(line), -1)
		for _, word := range words {
			// send word in channel
			fmt.Printf("Sending: %s : %d\n", word, documentIndex)
			ch <- Word{word, documentIndex}
		}
	}
	wg.Done()
}

func contains(list []int, target int) bool {
	for _, n := range list {
		if target == n {
			return true
		}
	}
	return false
}
