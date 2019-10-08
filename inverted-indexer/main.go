package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Println("Wrong number of inputs, need two.")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	invertedIndex := buildInvertedIndex(scanner)
	fmt.Println(toString(invertedIndex))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func buildInvertedIndex(scanner *bufio.Scanner) map[string][]int {

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
		for document.Scan() {
			line := document.Text()
			words := re.FindAllString(string(line), -1)
			for _, word := range words {
				seen, ok := index[word]
				if ok && seen[len(seen)-1] != documentIndex {
					index[word] = append(seen, documentIndex)
				} else {
					index[word] = []int{documentIndex}
				}
			}
		}

		documentIndex++
	}

	return index
}

func toString(m map[string][]int) string {
	b := new(bytes.Buffer)
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, key := range keys {
		str := ""
		indices := m[key]
		for _, index := range indices {
			str += strconv.Itoa(index) + " "
		}
		str = strings.Trim(str, " ")
		fmt.Fprintf(b, "%s: %s\n", key, str)
	}
	return strings.Trim(b.String(), "\n")
}
