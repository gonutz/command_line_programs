package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		io.Copy(os.Stdout, os.Stdin)
	} else if len(args) == 1 && args[0] == "byte" {
		nBytes(1)
	} else if len(args) == 2 && args[1] == "bytes" {
		n, _ := strconv.Atoi(args[0])
		nBytes(n)
	} else if len(args) == 1 && args[0] == "line" {
		nLines(1)
	} else if len(args) == 2 && args[1] == "lines" {
		n, _ := strconv.Atoi(args[0])
		nLines(n)
	} else if len(args) == 1 {
		n, _ := strconv.Atoi(args[0])
		nBytes(n)
	}
}

func nBytes(n int) {
	if n <= 0 {
		return
	}

	io.CopyN(os.Stdout, os.Stdin, int64(n))
}

func nLines(n int) {
	if n <= 0 {
		return
	}

	s := bufio.NewScanner(os.Stdin)
	for n > 0 && s.Scan() {
		os.Stdout.Write(s.Bytes())
		os.Stdout.Write([]byte{'\n'})
		n--
	}
}
