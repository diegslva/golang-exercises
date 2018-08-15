package main

import (
	"flag"
	"log"
	"os"
)

var flagFile string

func init() {
	flag.StringVar(&flagFile, "file", "test.txt", "a new file to create")
}

func main() {
	flag.Parse()
	newFile, err := os.Create(flagFile)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(newFile)
	newFile.Close()
}
