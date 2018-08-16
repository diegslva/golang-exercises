package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// Open a file for reading
	file, err := os.OpenFile("test.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// The file.Read() function will rapilly read a tiny file
	// into a large byte slice, but io.ReadFull() will return
	// an error if the file is smaller than the byte slice.
	byteSlice := make([]byte, 2)
	numBytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Number of bytes read: %d\r\n", numBytesRead)
	log.Printf("Data read: %s\r\n", byteSlice)
}
