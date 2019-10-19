package indexers_test

import (
	"bufio"
	"testing"

	"github.com/bgeraci714/indexers"
	test "github.com/bgeraci714/indexers_test"
)

func benchmarkMultipleWriter(threads int, b *testing.B) {
	b.StopTimer()
	f := test.InitializeTests()
	defer f.Close()

	for i := 0; i < b.N; i++ {
		scanner := bufio.NewScanner(f)
		b.StartTimer()
		indexers.BuildInvertedIndexWithMultipleWriters(scanner, threads)
		b.StopTimer()
	}
}

func BenchmarkMultipleWriter1OfEachThread(b *testing.B) { benchmarkMultipleWriter(1, b) }
func BenchmarkMultipleWriter2OfEachThread(b *testing.B) { benchmarkMultipleWriter(2, b) }
func BenchmarkMultipleWriter3OfEachThread(b *testing.B) { benchmarkMultipleWriter(3, b) }
func BenchmarkMultipleWriter4OfEachThread(b *testing.B) { benchmarkMultipleWriter(4, b) }
func BenchmarkMultipleWriter5OfEachThread(b *testing.B) { benchmarkMultipleWriter(5, b) }
func BenchmarkMultipleWriter6OfEach(b *testing.B)       { benchmarkMultipleWriter(6, b) }
func BenchmarkMultipleWriter7OfEach(b *testing.B)       { benchmarkMultipleWriter(7, b) }
func BenchmarkMultipleWriter8OfEach(b *testing.B)       { benchmarkMultipleWriter(8, b) }
func BenchmarkMultipleWriter9OfEach(b *testing.B)       { benchmarkMultipleWriter(9, b) }
func BenchmarkMultipleWriter10OfEach(b *testing.B)      { benchmarkMultipleWriter(10, b) }
