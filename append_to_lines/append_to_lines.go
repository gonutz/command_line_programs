package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		panic("One argument expected: the string to append")
	}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fmt.Println(s.Text() + os.Args[1])
	}
}
