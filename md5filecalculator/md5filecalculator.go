package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"

	"github.com/go-playground/log"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	h := md5.New()
	io.Copy(h, file)
	fmt.Printf("%x\n", h.Sum(nil))
}
