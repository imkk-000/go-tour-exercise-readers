package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (myReader MyReader) Read(buffer []byte) (int, error) {
	bufferLength := len(buffer)
	for index := 0; index < bufferLength; index++ {
		buffer[index] = 'A'
	}
	return bufferLength, nil
}

func main() {
	reader.Validate(MyReader{})
}
