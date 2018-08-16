package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// open file for reading
	file, err := os.OpenFile("test.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// os.File.Read(), io.ReadFull(), and io.ReadAtLeast
	// all work with a fixed byte slice that you make
	// before you read

	// ioutil.ReadAll read every byte
	// from the reader ( in this case a file )
	// and return a slice of unknown slice
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Data as hex: %x\n", data)
	fmt.Printf("Data as string: %s\n", data)
	fmt.Printf("Number of bytes read: %d", len(data))
}
