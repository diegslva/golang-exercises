package main

import (
	"flag"
	"log"
	"os"
)

var flagFile string

func init() {
	flag.StringVar(&flagFile, "file", "test.txt", "a file to delete")
}

func main() {
	flag.Parse()
	err := os.Remove(flagFile)
	if err != nil {
		log.Fatalln(err)
	}
}
