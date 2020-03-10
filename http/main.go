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
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// // buat byte slice yang berisi 99999 element kosong
	// bs := make([]byte, 99999)
	// // Gunakan interface Body yang memiliki Read function (ini adalah Reader)
	// // Reader interface digunakan untuk mengonversi apapun type data menjadi []byte
	// resp.Body.Read(bs)

	// fmt.Println(string(bs))
	/////////////////////////////////////////////////////

	// io.Copy(os.Stdout, resp.Body)
	////////////////////////////////////

	lw := logWriter{}

	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
