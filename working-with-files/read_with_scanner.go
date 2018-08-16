package main

// [Scanner] is part of bufio package. It is useful to step through files
// at specific delimiters. Commonly, the new line character is used
// as a delimiter to break up a file by lines.
// In a CSV file, commas would be a delimiter. The [os.File] can be
// wrapped in a [bufio.Scanner] just like a buffered reader.
// We call Scan() to read up to the next delimiter, and then use
// Text() or Bytes() to get the data that was read.
//
// The delimiter is not a simple byte or character. There is actually
// a special function you have to implement that will determine
// where the next delimiter is, how far foward to advanced the
// pointer and what data to return. If no custom [SplitFunc] is
// provided, it defaults to [ScanLines] which will split
// at every newline character.
// Other split functions included in bufio is [ScanRunes] and [ScanWords]

// To define your own split function, match this fingerprint
//
// type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)
//
// Returning (0, nil, nil) will tell the scanner
// to scan again but with a bigger buffer because
// it wasn`t enough data to reach the delimiter
import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Open file and create scanner on top of it
	file, err := os.OpenFile("test.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(file)

	// Default scanner is bufio.ScanLines. Lets use ScanWords
	// Could also use a custom function of SplitFunc type
	scanner.Split(bufio.ScanWords)

	// Scan for the next token.
	success := scanner.Scan()
	if !success {
		// False on error or EOF. Check error
		err = scanner.Err()
		if err != nil {
			log.Println("Scan completed and reached EOF.")
		} else {
			log.Fatalln(err)
		}
	}

	// Get data from scan with Bytes() or Text()
	fmt.Println("First world found:", scanner.Text())

	// Call scanner.Scan() again to find next token
}
