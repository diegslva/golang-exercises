package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// Open original file
	originalFile, err := os.Open("test.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer originalFile.Close()

	// Create a new file
	newFile, err := os.Create("test_copy.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer newFile.Close()

	// Copy the bytes to destination from source
	bytesWriten, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Copied %d bytes.", bytesWriten)

	// Commit the file contents
	// Flushes memory to disk
	err = newFile.Sync()
	if err != nil {
		log.Fatalln(err)
	}
}
