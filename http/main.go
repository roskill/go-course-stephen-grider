package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{} // custom type implementing Writer interface

func main() {
	res, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}
	io.Copy(lw, res.Body)

	// io.Copy(os.Stdout, res.Body)

	// pass a byte slice to Read which fills it up with res
	// data := make([]byte, 99999)
	// res.Body.Read(data)
	// fmt.Println(string(data))
}

func (logWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	fmt.Println("Just wrote this many bytes:", len(p))
	return len(p), nil
}
