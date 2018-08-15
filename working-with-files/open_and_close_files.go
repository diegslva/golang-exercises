package main

import (
	"log"
	"os"
)

func main() {
	// Simply read only open. We will cover actually reading
	// and writing in examples further down the page
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatalln(err)
	}
	file.Close()

	// Open file with more options. Last param is the permission
	// mode. Second param is the attributes when opening
	file, err = os.OpenFile("test.txt", os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	file.Close()

	// Use theses attributes individually or combined
	// with an OR for second arg of OpenFile()
	// e.g. os.CREATE|os.O_APPEND
	// or os.O_CREATE|os.O_TRUNC|os.O_WRONLY

	// os.O_RDONLY // Read only
	// os.O_WRONLY // Write only
	// os.O_RDWR // Read and Write
	// os.O_APPEND // Append to end of file
	// os.O_CREATE // Create if none exist
	// os.O_TRUNC // Truncate file when opening

}
