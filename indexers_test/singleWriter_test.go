package indexers_test

import (
	"bufio"
	"testing"

	"github.com/bgeraci714/indexers"
	test "github.com/bgeraci714/indexers_test"
)

func BenchmarkSingleWriter(b *testing.B) {
	b.StopTimer()
	f := test.InitializeTests()
	defer f.Close()

	for i := 0; i < b.N; i++ {
		scanner := bufio.NewScanner(f)
		b.StartTimer()
		indexers.BuildInvertedIndexWithSingleWriter(scanner)
		b.StopTimer()
	}

}
