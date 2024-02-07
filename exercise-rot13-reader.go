package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	capital := b >= 'A' && b <= 'Z'
	if !capital && (b < 'a' || b > 'z') {
		return b
	}
	b += 13
	if capital && b > 'Z' {
		b -= 26
	}
	if !capital && b > 'z' {
		b -= 26
	}
	return b
}

func (r13 *rot13Reader) Read(b []byte) (n int, e error) {
	n, err := r13.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
