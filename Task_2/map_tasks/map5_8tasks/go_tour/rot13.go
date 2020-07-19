package main

import (
	"io"
	"os"
	"strings"
)

func rot13(b byte) byte {
	switch {
	case b < 'A' || b > 'z':
		return b
	// "M" - последняя буква первой половины алфавита
	case (b > 'M' && b <= 'Z') || (b > 'm' && b <= 'z'):
		return b - 13
	default:
		return b + 13
	}
	return b
}

type rot13Reader struct {
	ir io.Reader
}

func (rtR rot13Reader) Read(p []byte) (int, error) {
	n, err := rtR.ir.Read(p)
	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
