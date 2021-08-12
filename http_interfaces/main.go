package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	//fmt.Printf("%+v", resp)

	/*
		Read metho spews out the output to the byte slice that is provided as an input.
		If the size of the byte slice is less than the output, read method truncates the output.
		To prevent this, we are creating a byte slice with a specified size.
		We dont do this in general. There are lot of helper functions available to pass inputs to read method.
	*/
	/*bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))*/

	lw := logWriter{}
	/* Another way to print the output of a http request*/
	//io.Copy(os.Stdout, resp.Body)
	io.Copy(lw, resp.Body)
}
