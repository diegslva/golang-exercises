package main

import (
        "os"
        "log"
        "io"
)

func main() {
        file, err := os.OpenFile("test.txt", os.O_RDONLY, 0666)
        if err != nil {
                log.Fatalln(err)
        }
        defer file.Close()

        byteSlice := make([]byte, 512)
        minBytes := 8
        // io.ReadAtLeast() will return a error if it cannot
        // find at least minBytes to read. It will read as
        // many bytes as byteSlice can hold
        numBytesRead, err := io.ReadAtLeast(file, byteSlice, minBytes)
        if err != nil {
                log.Fatalln(err)
        }
        log.Printf("Number of bytes read: %d\r\n", numBytesRead)
        log.Printf("Data read: %s\r\n", byteSlice)
        log.Printf("Minimum bytes to read: %d\r\n", minBytes)
}
