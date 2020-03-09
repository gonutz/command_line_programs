package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	n := 0
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		n++
	}
	fmt.Println(n)
}
