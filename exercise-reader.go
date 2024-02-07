package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
//
//	Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
func (r MyReader) Read(p []byte) (n int, e error) {
	for i := range p {
		p[i] = 'A'
	}

	return len(p), nil
}

func main() {
	reader.Validate(MyReader{})
}
