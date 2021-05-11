package main

import (
	"encoding/base64"
	"io"
	"os"
)

func main() {
	w := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	io.Copy(w, os.Stdin)
	w.Close()
}
