package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Open file and create a buffered reader on top
	file, err := os.OpenFile("test.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	bufferedReader := bufio.NewReader(file)

	// Get bytes without advancing pointer
	byteSlice := make([]byte, 5)
	byteSlice, err = bufferedReader.Peek(5)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Peeked at 5 bytes: %s\n", byteSlice)

	// Read and advance pointer
	numBytesRead, err := bufferedReader.Read(byteSlice)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Read %d bytes: %s\r\n", numBytesRead, byteSlice)

	// Read 1 byte, error if no byte to read
	mByte, err := bufferedReader.ReadByte()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Read 1 byte: %c\r\n", mByte)

	// Read up to and including delimiter
	// Returns byte slice
	dataBytes, err := bufferedReader.ReadBytes('\n')
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Read bytes: %s\r\n", dataBytes)

	// Read up to and including delimiter
	// Returns string
	dataString, err := bufferedReader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Read string: %s\r\n", dataString)

}
