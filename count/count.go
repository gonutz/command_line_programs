package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	n, err := count(os.Args[1:], os.Stdin)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
	} else {
		fmt.Fprint(os.Stdout, n)
	}
}

func usage() error {
	return fmt.Errorf(`usage: %s bytes/lines/letters

  Counts the number of specified items in standard input and prints the number
  to standard output. If an error occurs, nothing is printed to standard output
  and the error is printed to standard error.

  These are the options, specify exactly one of these:

  bytes        - counts all bytes
  lines        - counts all lines, delimited by "\r\n" or "\n"
  letters      - counts all UTF-8 letters, skips BOM
`, os.Args[0])
}

func count(args []string, r io.Reader) (uint64, error) {
	isArg := func(arg string) bool {
		return len(args) == 1 && args[0] == arg
	}

	if isArg("bytes") {
		var w lineCounter
		_, err := io.Copy(&w, r)
		return w.totalBytes, err
	} else if isArg("lines") {
		return new(lineCounter).countAll(r)
	} else if isArg("letters") {
		return new(utf8counter).countAll(r)
	} else {
		return 0, usage()
	}
}

type lineCounter struct {
	totalBytes, lines uint64
}

func (w *lineCounter) countAll(r io.Reader) (uint64, error) {
	w.lines = 1
	_, err := io.Copy(w, r)
	if err != nil {
		return 0, err
	}
	if w.totalBytes == 0 {
		return 0, nil
	}
	return w.lines, nil
}

func (w *lineCounter) Write(b []byte) (int, error) {
	w.totalBytes += uint64(len(b))
	for _, b := range b {
		if b == '\n' {
			w.lines++
		}
	}
	return len(b), nil
}

type utf8counter struct {
	count uint64
	skip  int
	err   error
}

func (w *utf8counter) countAll(r io.Reader) (uint64, error) {
	// Read the first 3 bytes separately, they might contain the byte order mark
	// (BOM) for UTF-8. If it is the BOM, we do not count it as a letter.
	// See below, after reading everything we subtract the first letter if it
	// was the BOM.
	var bom [3]byte
	n, err := io.ReadFull(r, bom[:])
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		return 0, err
	}

	// We still write the first 3 bytes along with the rest, in case it is not
	// the BOM. We subtract it later from the count if it was the BOM.
	_, err = w.Write(bom[:n])
	if err != nil {
		return 0, err
	}
	if w.err != nil {
		return 0, w.err
	}

	// Copy all the bytes, either the reader can fail or our write can fail.
	// Keep either error.
	_, err = io.Copy(w, r)
	if err != nil {
		return 0, err
	}
	if w.err != nil {
		return 0, w.err
	}

	// The Write function checks for multi-byte UTF-8 encodings. It knows by the
	// first byte how many bytes are to come in a code point. At the end of a
	// valid UTF-8 character, w.skip reaches 0. If we are not at 0 at the end of
	// all bytes, then the last letter is incomplete and we do not have valid
	// UTF-8.
	if w.skip != 0 {
		return 0, errors.New("invalid UTF-8 code point")
	}

	// Now, after reading everything, we check if it started with the BOM and
	// subtract one in that case. The BOM has a valid UTF-8 structure and is
	// thus counted by the Write function as one letter.
	// We subtract after the fact because w.count is unsigned and we do not want
	// to run into any peculiarities with decrementing an unsigned 0 value.
	if bom == [3]byte{0xEF, 0xBB, 0xBF} {
		w.count--
	}

	return w.count, nil
}

func (w *utf8counter) Write(b []byte) (int, error) {
	for _, b := range b {
		if w.skip == 0 {
			w.count++
			if b&0x80 == 0 {
			} else if b&0xE0 == 0xC0 {
				w.skip = 1
			} else if b&0xF0 == 0xE0 {
				w.skip = 2
			} else if b&0xF8 == 0xF0 {
				w.skip = 3
			} else {
				w.err = errors.New("invalid UTF-8 code point")
				return 0, w.err
			}
		} else {
			if b&0xC0 != 0x80 {
				w.err = errors.New("invalid UTF-8 code point")
				return 0, w.err
			}
			w.skip--
		}
	}
	return len(b), nil
}
