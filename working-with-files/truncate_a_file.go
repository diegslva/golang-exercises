package main

import (
	"log"
	"os"
)

func main() {
	// trucate a file to 100 bytes. if file is
	// less than 100 bytes the original content will remain
	// at beginning, and the rest of the space is filled
	// with null bytes. If it is over 100 bytes,
	// everything past 100 bytes will be lost. Either way
	// we will end up with exactly 100 bytes.
	// Pass in 0 to truncate to a completely  empty file

	err := os.Truncate("test.txt", 100)
	if err != nil {
		log.Fatalln(err)
	}

}
