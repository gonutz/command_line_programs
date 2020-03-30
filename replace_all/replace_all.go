package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}
	replace := []byte(unescape(os.Args[1]))
	with := []byte(unescape(os.Args[2]))
	data, _ := ioutil.ReadAll(os.Stdin)
	out := bytes.Replace(data, replace, with, -1)
	os.Stdout.Write(out)
}

func unescape(s string) string {
	s = strings.Replace(s, `\r`, "\r", -1)
	s = strings.Replace(s, `\n`, "\n", -1)
	s = strings.Replace(s, `\t`, "\t", -1)
	s = strings.Replace(s, `\\`, "\\", -1)
	return s
}
