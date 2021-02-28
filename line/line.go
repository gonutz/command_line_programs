package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		panic("expecting one argument: the line index (starting at 0)")
	}
	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("expecting one argument: the line index (starting at 0)")
	}
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("unable to read stdin: " + err.Error())
	}
	lines := bytes.Split(data, []byte{'\n'})
	if 0 <= i && i < len(lines) {
		line := lines[i]
		if bytes.HasSuffix(line, []byte{'\r'}) {
			line = line[:len(line)-1]
		}
		fmt.Printf("%s", line)
	}
}
