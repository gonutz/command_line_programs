package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		panic("expecting one argument: the line to look for")
	}
	refLine := []byte(os.Args[1])
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("unable to read stdin: " + err.Error())
	}
	lines := bytes.Split(data, []byte{'\n'})
	for i := range lines {
		line := lines[i]
		if bytes.HasSuffix(line, []byte{'\r'}) {
			line = line[:len(line)-1]
		}
		if bytes.Equal(line, refLine) {
			fmt.Print(i)
			return
		}
	}
	fmt.Print("-1")
}
