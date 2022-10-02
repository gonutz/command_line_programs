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
	} else if len(args) == 1 && args[0] == "line" {
		nLines(1)
	} else if len(args) == 2 && args[1] == "lines" {
		n, _ := strconv.Atoi(args[0])
		nLines(n)
	}
}

func nLines(n int) {
	if n <= 0 {
		return
	}

	buffer := make([]string, n)
	next := 0
	count := 0
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		buffer[next] = s.Text()
		next = (next + 1) % len(buffer)
		count++
	}

	if count < len(buffer) {
		buffer = buffer[:count]
		next = 0
	}

	for i := 0; i < n; i++ {
		os.Stdout.WriteString(buffer[(next+i)%len(buffer)])
		if i != n-1 {
			os.Stdout.Write([]byte{'\n'})
		}
	}
}
