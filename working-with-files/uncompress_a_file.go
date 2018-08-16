package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Open the gzip file that we want to
	// uncompress. The file is a reader but
	// we could use any data source.
	// It is common to web servers to return
	// gzipped contents to save bandwith
	// and in that case the data is not in a file
	// on the file system but in a memory buffer
	gzipFile, err := os.OpenFile("test.txt.gz", os.O_RDONLY, 0)
	if err != nil {
		log.Fatalln(err)
	}
	defer gzipFile.Close()

	// Create a gzip reader on top of
	// the file reader. Again could be
	// any  type reader.
	gzipReader, err := gzip.NewReader(gzipFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer gzipReader.Close()

	// Uncompress to a writer. We'll use a file writer.
	outfileWriter, err := os.OpenFile(
		"unzipped.txt",
		os.O_RDWR|os.O_CREATE|os.O_TRUNC,
		0666,
	)
	defer outfileWriter.Close()

	// Copy contents of gzipped file to output file
	writtenBytes, err := io.Copy(outfileWriter, gzipReader)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Bytes extracted: %v", writtenBytes)
}
