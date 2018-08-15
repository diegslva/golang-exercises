package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Create a hard link
	// You will have two files that point to the same contents
	// Changing the contents of one will change the other
	// Deleting/renaming one will not affect the other
	err := os.Link("original.txt", "original_also.txt")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("creating sym")
	// Create a symlink
	err = os.Symlink("original.txt", "original_sym.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// Lstat will return a file info, but if it is actually
	// a symlink, it will return info about the symlink
	// It will not follow the link and give real information
	// about the real file
	// Symlink do not work in Windows
	fileInfo, err := os.Lstat("original_sym.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Link info: %+v\r\n", fileInfo)

	// Change the ownership of a symlink only
	// and not the file it points to
	err = os.Lchown("original_sym.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Fatalln(err)
	}
}
