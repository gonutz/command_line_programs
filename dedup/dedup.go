package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	seen := make(map[string]bool)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if !seen[s.Text()] {
			fmt.Println(s.Text())
			seen[s.Text()] = true
		}
	}
}
