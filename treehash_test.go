package treehash

import (
	"testing"
)

func BenchmarkTreeHash(b *testing.B) {
	b.StopTimer()
	data := make([]uint8, ChunkSize*2)
	for i := 0; i < ChunkSize*2; i++ {
		data[i] = uint8(i)
	}
	th := New()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		th.Reset()
		th.Write(data)
		th.Sum(nil)
	}
}
