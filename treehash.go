package treehash

import (
	"bytes"
	"crypto/sha256"
	"hash"
)

const chunkSize = 1024 * 1024

type digest struct {
	buf *bytes.Buffer
}

// New creates a new hash.Hash computing the Tree Hash checksum.
func New() hash.Hash {
	return &digest{
		new(bytes.Buffer),
	}
}

func (d *digest) Size() int { return sha256.Size }

func (d *digest) BlockSize() int { return sha256.BlockSize }

func (d *digest) Reset() {
	d.buf = new(bytes.Buffer)
}

func chunk(buffer *bytes.Buffer) [][]byte {
	sha := sha256.New()
	chunks := [][]byte{}
	chunk := make([]byte, chunkSize)
	for {
		read, err := buffer.Read(chunk)
		if err != nil {
			break // It never returns an error.
		}
		sha.Reset()
		sha.Write(chunk[:read])
		chunks = append(chunks, sha.Sum(nil))
	}
	return chunks
}

func compute(chunks [][]byte) []byte {
	sha := sha256.New()
	previousLevel := chunks
	for {
		if len(previousLevel) == 1 {
			break
		}

		length := len(previousLevel) / 2
		if len(previousLevel)%2 != 0 {
			length++
		}

		currentLevel := make([][]byte, length)
		for i, j := 0, 0; i < len(previousLevel); i, j = i+2, j+1 {
			if len(previousLevel)-i > 1 {
				sha.Reset()
				sha.Write(previousLevel[i])
				sha.Write(previousLevel[i+1])
				currentLevel[j] = sha.Sum(nil)
			} else {
				currentLevel[j] = previousLevel[i]
			}
		}

		previousLevel = currentLevel
	}
	return previousLevel[0]
}

func (d *digest) Sum(in []byte) []byte {
	chunks := chunk(d.buf)
	return compute(chunks)
}

func (d *digest) Write(p []byte) (n int, err error) {
	return d.buf.Write(p)
}
