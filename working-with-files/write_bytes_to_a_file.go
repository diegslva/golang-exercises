package main

import (
	"log"
	"os"
)

func main() {
	// Open a new file for writing only
	file, err := os.OpenFile(
		"test.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// Write bytes to a file
	byteSlice := []byte("Bytes!\r\nDone!")
	// file.Write returns how much bytes was wrote
	// and a error if any
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Wrote %d bytes\r\n", bytesWritten)
}
