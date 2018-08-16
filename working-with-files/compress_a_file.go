// This example uses gzip but standard library
// supports zlib, bz2, flat and lzw
package main

import (
	"compress/gzip"
	"log"
	"os"
)

func main() {
	// Create a .gz file to write to
	outputFile, err := os.OpenFile(
		"test.txt.gz",
		os.O_RDWR|os.O_CREATE|os.O_TRUNC,
		0666,
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer outputFile.Close()

	// Create a gzip writer on top of file writer
	gzipWriter := gzip.NewWriter(outputFile)
	defer gzipWriter.Close()

	// When we write to the gzip writer
	// it will in turn compress the contents
	// and then write it to the underlying file
	// writer as well
	// We don't have to worry how all the
	// compression works since we just
	// use  it as a simple writer interface
	// that we send bytes to
	_, err = gzipWriter.Write([]byte("Gophers rule!\r\n"))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Compressed data writen to file.")

}
