package main

import (
	"io"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rr *rot13Reader) Read(p []byte) (n int, err error) {
	n2 := 8
	var err2 error = nil
	return n2, err2
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	p = make([]byte, 8)
	r.Read(p)
	//io.Copy(os.Stdout, &r)
}
