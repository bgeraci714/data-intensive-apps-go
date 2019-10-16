package indexer

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"os"
	"regexp"
	"sync"
)

// BuildInvertedIndexWithMultipleWriters builds an inverted index using a multithreaded approach
func BuildInvertedIndexWithMultipleWriters(scanner *bufio.Scanner, threads int) map[string][]int {
	var index sync.Map
	re := regexp.MustCompile(`([a-z]|[A-Z])+`)
	documentIndex := 0
	ch := make(chan Word)
	var wgReaders, wgWriters, wgSorters sync.WaitGroup

	// read in each file one at a time
	for scanner.Scan() {
		file, err := os.Open(scanner.Text())
		if err != nil {
			fmt.Println("ERROR: Could not find file: " + scanner.Text())
			continue
		}

		// read sub file and fire off separate reader goroutines for the file
		document := bufio.NewScanner(file)
		go reader(document, documentIndex, re, ch, &wgReaders)
		wgReaders.Add(1)
		documentIndex++
	}

	// build channels for writers
	chs := make([]chan Word, threads)
	for i := 0; i < threads; i++ {
		chs[i] = make(chan Word)
	}

	// start sorter goroutines
	for i := 0; i < threads; i++ {
		go sorter(ch, chs, &wgSorters)
		wgSorters.Add(1)
	}

	// start writer goroutines
	for i := 0; i < threads; i++ {
		go writer(chs[i], &index, &wgWriters)
		wgWriters.Add(1)
	}

	wgReaders.Wait() // wait for the readers to finish
	close(ch)

	wgSorters.Wait() // wait for the sorters to finish
	for i := 0; i < threads; i++ {
		close(chs[i])
	}

	wgWriters.Wait() // wait for the writers to finish
	return toMap(&index)
}

// makes sure that each channel's keys will be disjoint from each other
func sorter(ch <-chan Word, chs []chan Word, wg *sync.WaitGroup) {
	for {
		if word, more := <-ch; more {
			i := hash(word.word) % len(chs)
			fmt.Printf("Sorting: %s -> %d\n", word.word, i)
			chs[i] <- word // send word to appropriate channel
		} else {
			wg.Done()
			return
		}
	}
}

func hash(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32())
}

// the really important thing here is that the words received over this channel form
// a disjoint set from all other channels. In this way, no locking is necessary to
// maintain consistency between the calls to LoadOrStore and Store.
func writer(ch <-chan Word, index *sync.Map, wg *sync.WaitGroup) {
	for {
		if word, more := <-ch; more {
			fmt.Printf("Writing: %s : %d\n", word.word, word.index)
			seen, loaded := index.LoadOrStore(word.word, []int{word.index})
			if loaded {
				seen = append(seen.([]int), word.index)
				index.Store(word.word, seen)
			}
		} else {
			wg.Done()
			return
		}
	}
}

func toMap(i *sync.Map) map[string][]int {
	m := make(map[string][]int)
	i.Range(func(key, value interface{}) bool {
		m[key.(string)] = value.([]int)
		return true
	})
	return m
}
