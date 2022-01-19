package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (reader MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	b[1] = 'A'
	b[2] = 'A'
	b[3] = 'A'
	return 2, nil
}

func main() {
	reader.Validate(MyReader{})
}
