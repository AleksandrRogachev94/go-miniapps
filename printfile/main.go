package main

import (
	"io"
	"os"
)

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, f)

	f.Close()
}
