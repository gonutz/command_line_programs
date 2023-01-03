package main

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {
	start := time.Now()
	print := make(chan bool)
	go func() {
		for {
			select {
			case <-print:
				start = time.Now()
			default:
				s := fmt.Sprint(time.Now().Sub(start)) + "          "
				s += strings.Repeat("\b", utf8.RuneCountInString(s))
				fmt.Print(s)
				time.Sleep(time.Millisecond)
			}
		}
	}()
	for {
		fmt.Scanln()
		print <- true
	}
}
