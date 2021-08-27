package exhibit

import (
	"math/bits"
	"testing"
)

var values = []struct {
	in   uint64
	want int
}{
	{0xffff, 16},
	{0, 0},
	{1, 1},
	{2, 1},
	{7, 3},
	{9, 2},
	{0xffffff00, 24},
}

func BenchmarkCountOnes(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, value := range values {
			got := countOnes(value.in)
			if got != value.want {
				b.Fatalf("%d:%x:: got=%d want=%d", value.in, value.in, got, value.want)
			}
		}
	}
}

func BenchmarkCountOnesFast(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, value := range values {
			got := bits.OnesCount64(value.in)
			if got != value.want {
				b.Fatalf("%d:%x:: got=%d want=%d", value.in, value.in, got, value.want)
			}
		}
	}
}
