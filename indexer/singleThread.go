package indexer

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
