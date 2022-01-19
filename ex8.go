package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot13.r.Read(b)
	for i := 0; i < len(b); i++ {
		if (b[i] > 'A' || b[i] > 'a') && (b[i] < 'N' || b[i] < 'n') {
			b[i] += 13
		} else if (b[i] > 'M' || b[i] > 'm') && (b[i] < 'Z' || b[i] < 'z') {
			b[i] -= 13
		}

	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
