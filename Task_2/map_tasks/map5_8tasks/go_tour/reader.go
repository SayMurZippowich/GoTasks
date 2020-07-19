package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// Add a Read([]byte) (int, error) method to MyReader.
func (mr MyReader) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = 'A'
	}
	return len(p), nil
}

func main() {
	reader.Validate(MyReader{})
}
