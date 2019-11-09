package indexers

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

func (w Word) String() string {
	return fmt.Sprintf("[%s, %d]", w.word, w.index)
}

func reader(document *bufio.Scanner, documentIndex int, re *regexp.Regexp, ch chan<- Word, wg *sync.WaitGroup) {
	for document.Scan() {
		line := document.Text()
		words := re.FindAllString(string(line), -1)
		for _, word := range words {
			// send word in channel
			w := Word{word, documentIndex}
			ch <- w
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
