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
		panic("expecting one argument: the field index (starting at 0)")
	}
	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("expecting one argument: the field index (starting at 0)")
	}
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("unable to read stdin: " + err.Error())
	}
	fields := bytes.Fields(data)
	if 0 <= i && i < len(fields) {
		fmt.Printf("%s", fields[i])
	}
}
