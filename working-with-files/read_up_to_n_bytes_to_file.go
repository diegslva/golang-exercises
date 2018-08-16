package main

import (
	"log"
	"os"
)

func main() {
	// Open file for reading
	file, err := os.OpenFile("test.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// Read up to len(b) bytes from the file
	// Zero bytes written means end of file
	// End of file returns error type io.EOF
	byteSlice := make([]byte, 16)
	bytesRead, err := file.Read(byteSlice)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Number of bytes read: %d", bytesRead)
	log.Printf("Data read: %v", byteSlice)

}
