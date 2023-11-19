package main

import (
	"fmt"
	"io"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (n int, err error) {
	lenFromFirstReader, errFromFirstReader := r.r.Read(b)
	max := byte('a' + 25)
	if errFromFirstReader != io.EOF {
		for i := 0; i < lenFromFirstReader; i++ {
			if b[i] != 32 {
				if b[i]+byte(13) > max {
					b[i] = 'a' + b[i] + byte(13) - max
				} else {
					b[i] += 13
				}
			}
		}
		return lenFromFirstReader, nil
	}
	return 0, errFromFirstReader
}

func main() {
	s := strings.NewReader(" ")
	r := rot13Reader{s}
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	//bytesT := make([]byte, 8)
	//r.Read(bytesT)
	//io.Copy(os.Stdout, &r)

}
