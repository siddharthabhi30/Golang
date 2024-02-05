package main

import (
	"fmt"
	"io"
	"strings"
)

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (m MyReader) Read(b []byte) (i int, e error) {
	for x := range b {
		b[x] = 'A'
	}
	fmt.Println(cap(b))
	return len(b), nil
}

func main() {
	r := strings.NewReader("Hello,Read")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	
}

// Read populates the given byte slice ///
//with data and returns the number of bytes populated and an error value. It returns an io.EOF error when the stream ends.

// The example code creates a strings.Reader and consumes
// its output 8 bytes at a time.
