package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	lw := logWriter{}

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Print(string(bs))

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("just wrote this many bytes: ", len(bs))
	return len(bs), nil
}
