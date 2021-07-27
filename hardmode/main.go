package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1]) // go run main.go file.txt
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// data := make([]byte, 99999)
	// file.Read(data)
	// fmt.Println(string(data))

	io.Copy(os.Stdout, file)
}
