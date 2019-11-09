package indexers

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// BuildInvertedIndexWithSingleThread builds single threaded inverted index builder
func BuildInvertedIndexWithSingleThread(scanner *bufio.Scanner) map[string][]int {

	index := make(map[string][]int)
	re := regexp.MustCompile(`([a-z]|[A-Z])+`)
	documentIndex := 0
	for scanner.Scan() {
		file, err := os.Open(scanner.Text())
		if err != nil {
			fmt.Println("ERROR: Could not find file: " + scanner.Text())
			continue
		}
		// read sub file
		document := bufio.NewScanner(file)
		AddDocumentToIndex(document, documentIndex, &index, re)

		documentIndex++
	}

	return index
}

// AddDocumentToIndex adds a document to an index given a single thread
func AddDocumentToIndex(document *bufio.Scanner, documentIndex int, index *map[string][]int, re *regexp.Regexp) {
	for document.Scan() {
		line := document.Text()
		words := re.FindAllString(string(line), -1)
		for _, word := range words {
			seen, ok := (*index)[word]
			if !ok {
				(*index)[word] = []int{documentIndex}
			} else if seen[len(seen)-1] != documentIndex {
				(*index)[word] = append(seen, documentIndex)
			}
		}
	}
}
