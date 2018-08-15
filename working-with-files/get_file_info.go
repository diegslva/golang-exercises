package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Stat returns a file info. It will return
	// an error if ther is no file.
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("File name: %v\r\n", fileInfo.Name())
	fmt.Printf("Size in bytes: %v\r\n", fileInfo.Size())
	fmt.Printf("Permissions: %v\r\n", fileInfo.Mode())
	fmt.Printf("Last modified:  %v\r\n", fileInfo.ModTime())
	fmt.Printf("Is Directory: %v\r\n", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\r\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\r\n", fileInfo.Sys())
}
