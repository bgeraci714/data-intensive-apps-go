package indexer

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sync"
)

func putIfAbsent(index *map[string][]int, word Word, lock *sync.Mutex) []int {
	// grab global put lock
	lock.Lock()
	defer lock.Unlock()

	// check if object is already in map
	if seen, ok := (*index)[word.word]; ok {
		return seen
	}

	(*index)[word.word] = []int{word.index}
	// if so, release lock and return false
	return nil
}

func writer(ch <-chan Word, index *map[string][]int, wg *sync.WaitGroup) {
	// for {
	// 	word, more := <-ch
	// 	if more {
	// 		fmt.Printf("Received: %s : %d\n", word.word, word.index)

	// 		if seen := putIfAbsent(index, word); seen != nil {
	// 			// lock individual map element

	// 			// append

	// 			// unlock
	// 		}
	// 		// seen, ok := (*index)[word.word]
	// 		// if ok && !contains(seen, word.index) {
	// 		// 	(*index)[word.word] = append(seen, word.index)
	// 		// } else {
	// 		// 	(*index)[word.word] = []int{word.index}
	// 		// }
	// 	} else {
	// 		wg.Done()
	// 		return
	// 	}
	// }
}

// BuildInvertedIndexWithMultipleWriters builds an inverted index using a multithreaded approach
func BuildInvertedIndexWithMultipleWriters(scanner *bufio.Scanner, threads int) map[string][]int {
	// initialization
	index := make(map[string][]int)
	re := regexp.MustCompile(`([a-z]|[A-Z])+`)
	documentIndex := 0
	ch := make(chan Word)
	var wgReaders, wgWriters sync.WaitGroup

	for scanner.Scan() {
		file, err := os.Open(scanner.Text())
		if err != nil {
			fmt.Println("ERROR: Could not find file: " + scanner.Text())
			continue
		}

		// read sub file and fire off separate reader goroutine
		document := bufio.NewScanner(file)
		go reader(document, documentIndex, re, ch, &wgReaders)
		wgReaders.Add(1)
		documentIndex++
	}
	// writer
	// done := make(chan bool, 1)
	for i := 0; i < threads; i++ {
		go writer(ch, &index, &wgWriters)
		wgWriters.Add(1)
	}

	wgReaders.Wait() // wait for the readers to finish
	close(ch)        // send a message to the channel that it's done receiving new Words
	wgWriters.Wait() // wait for the writers to finish
	return index
}
