package indexer

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sync"
)

func singleWriter(ch <-chan Word, index *map[string][]int, done chan<- bool) {
	for {
		word, more := <-ch
		if more {
			seen, ok := (*index)[word.word]
			fmt.Printf("Received: %s : %d\n", word.word, word.index)
			if ok && !contains(seen, word.index) {
				(*index)[word.word] = append(seen, word.index)
			} else {
				(*index)[word.word] = []int{word.index}
			}
		} else {
			done <- true
			return
		}
	}
}

// BuildInvertedIndexWithSingleWriter builds an inverted index using a multithreaded approach
func BuildInvertedIndexWithSingleWriter(scanner *bufio.Scanner) map[string][]int {
	// initialization
	index := make(map[string][]int)
	re := regexp.MustCompile(`([a-z]|[A-Z])+`)
	documentIndex := 0
	ch := make(chan Word)
	var wg sync.WaitGroup

	// reader

	for scanner.Scan() {
		file, err := os.Open(scanner.Text())
		if err != nil {
			fmt.Println("ERROR: Could not find file: " + scanner.Text())
			continue
		}

		// read sub file
		document := bufio.NewScanner(file)
		go reader(document, documentIndex, re, ch, &wg)
		wg.Add(1)
		documentIndex++
	}
	// writer
	done := make(chan bool, 1)
	go singleWriter(ch, &index, done)

	wg.Wait()
	close(ch)
	<-done
	return index
}
