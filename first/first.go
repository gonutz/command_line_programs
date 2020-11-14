package main

import (
	"io"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		io.Copy(os.Stdout, os.Stdin)
		return
	}
	io.CopyN(os.Stdout, os.Stdin, int64(n))
}
