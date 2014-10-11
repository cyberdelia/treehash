package treehash

import (
	"fmt"
	"io"
	"testing"
)

func TestTreeHash(t *testing.T) {
	th := New()
	io.WriteString(th, "Even if I could be Shakespeare, I think I should still choose to be Faraday. - A. Huxley")
	s := fmt.Sprintf("%x", th.Sum(nil))
	if s != "1fb2eb3688093c4a3f80cd87a5547e2ce940a4f923243a79a2a1e242220693ac" {
		t.Fatal("wrong checksum")
	}
}

func BenchmarkTreeHash(b *testing.B) {
	b.StopTimer()
	data := make([]uint8, chunkSize*2)
	for i := 0; i < chunkSize*2; i++ {
		data[i] = uint8(i)
	}
	th := New()
	b.SetBytes(int64(len(data)))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		th.Reset()
		th.Write(data)
		th.Sum(nil)
	}
}
