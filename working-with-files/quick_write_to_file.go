package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// ioutil.WriteFile will create/open, write a slice of byte and close.
	// quick and dirty
	err := ioutil.WriteFile("test.txt", []byte("Hey mamma!\r\nI'm on tv :)"), 0666)
	if err != nil {
		log.Fatalln(err)
	}
	file, err := os.OpenFile("test.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	r, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Name of file: %v\r\n", file.Name())
	fmt.Printf("Your file: %v\r\n", string(r))
}
