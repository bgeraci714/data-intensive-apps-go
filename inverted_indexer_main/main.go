package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/bgeraci714/indexers"
)

func main() {
	if len(os.Args[1:]) != 2 {
		fmt.Println("Wrong number of inputs, need two.")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
		return
	}

	threads, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	var invertedIndex map[string][]int

	switch {
	case threads == 1:
		invertedIndex = indexers.BuildInvertedIndexWithSingleWriter(scanner)
	case threads >= 2:
		invertedIndex = indexers.BuildInvertedIndexWithMultipleWriters(scanner, threads)
	default:
		invertedIndex = indexers.BuildInvertedIndexWithSingleThread(scanner)
	}

	fmt.Println(toString(invertedIndex))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
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
