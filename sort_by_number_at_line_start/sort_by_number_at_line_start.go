package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var lines []string
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	sort.Stable(byFirstNumber(lines))
	for _, line := range lines {
		fmt.Println(line)
	}
}

type byFirstNumber []string

func (s byFirstNumber) Len() int      { return len(s) }
func (s byFirstNumber) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s byFirstNumber) Less(i, j int) bool {
	extractNum := func(s string) int {
		sign := 1
		for _, r := range s {
			if r == '-' {
				sign = sign * -1
			} else {
				break
			}
		}

		n := 0
		for _, r := range s {
			if '0' <= r && r <= '9' {
				n = n*10 + int(r-'0')
			} else {
				break
			}
		}

		return sign * n
	}
	return extractNum(s[i]) < extractNum(s[j])
}
