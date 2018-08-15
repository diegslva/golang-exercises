package main

import (
	"log"
	"os"
)

func main() {
	// Test read and write perrmisions. It is possible the file
	// does not exist and that will return a different
	// error that can be checked os.IsNotExist(err)
	file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Fatalln("Error: Write permission denied.")
		}
	}
	file.Close()

	// Test read permission
	file, err = os.OpenFile("test.txt", os.O_RDONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Fatalln("Error: Read permission denied.")
		}
	}

}
