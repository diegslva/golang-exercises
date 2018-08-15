package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	// Offset is how many bytes will move
	// Offset can be positive or negative
	var offset int64 = 5

	// Whence is the point of reference for offset
	// 0 - Begining of the file
	// 1 - Current position
	// 2 - End of file
	var whence int = 0

	newPosition, err := file.Seek(offset, whence)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("just moved to 5:", newPosition)

	// Go back 2 bytes from current position
	offset = -2
	// Current position
	whence = 1
	newPosition, err = file.Seek(offset, whence)
	if err != nil {
		fmt.Println("just moved back 2 bytes: ", newPosition)
	}

	// Find the current position byte getting the
	// return value from Seek after moving 0 bytes
	offset = 0
	whence = 1
	currentPosition, err := file.Seek(offset, whence)
	fmt.Println("Current position: ", currentPosition)

	// Go to begining of file
	newPosition, err = file.Seek(0, 0)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Position after seeking 0,0: ", newPosition)
}
