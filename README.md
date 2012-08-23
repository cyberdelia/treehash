# treehash

treehash implements SHA256 Tree Hash algorithm, notably used by Amazon Glacier.

## Installation

Download and install :

```
$ go get github.com/cyberdelia/treehash
```

Add it to your code :

```go
import "github.com/cyberdelia/treehash"
```

## Use

```go
file, _ := os.Open("archive.tar.gz")
th := treehash.New()
io.Copy(th, file)
checksum := fmt.Sprintf("%x", th.Sum(nil))
```
