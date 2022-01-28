package main

import (
	"bytes"
	"testing"

	"github.com/gonutz/check"
)

func TestCountBytes(t *testing.T) {
	checkBytes(t, "", 0)
	checkBytes(t, "a", 1)
	checkBytes(t, "abc", 3)
	checkBytes(t, "ü", 2)
}

func TestCountLines(t *testing.T) {
	checkLines(t, "", 0)
	checkLines(t, "a", 1)
	checkLines(t, "a\nb", 2)
	checkLines(t, "a\r\nb", 2)
	checkLines(t, "1\n2\n3", 3)
}

func TestUTF8Letters(t *testing.T) {
	checkUTF8LettersByteForByte(t, "", 0)
	checkUTF8LettersByteForByte(t, "a", 1)
	checkUTF8LettersByteForByte(t, "abc", 3)
	checkUTF8LettersByteForByte(t, "äöü", 3)
	checkUTF8LettersByteForByte(t, "嗎", 1)
	checkUTF8LettersByteForByte(t, "aä嗎嗎äa", 6)
}

func TestUTF8LettersWithBOM(t *testing.T) {
	// BOM at the start is not counted as a letter.
	checkUTF8Letters(t, "\xEF\xBB\xBF", 0)
	checkUTF8Letters(t, "\xEF\xBB\xBFabc", 3)
	// BOM after the start is counted as a letter.
	checkUTF8Letters(t, "\xEF\xBB\xBFabc\xEF\xBB\xBF", 4)
}

func TestInvalidUTF8(t *testing.T) {
	failUTF8(t, 0xFF)
	failUTF8(t, 'a', 0xFF)
	// Incomplete BOM:
	failUTF8(t, 0xEF)
	failUTF8(t, 0xEF, 0xBB)
}

func failUTF8(t *testing.T, input ...byte) {
	t.Helper()
	n, err := count([]string{"letters"}, bytes.NewReader(input))
	check.Eq(t, n, 0)
	check.Neq(t, err, nil)
}

func checkCount(t *testing.T, kind, input string, wantCount int) {
	t.Helper()
	n, err := count([]string{kind}, bytes.NewReader([]byte(input)))
	check.Eq(t, err, nil)
	check.Eq(t, n, wantCount)
}

func checkBytes(t *testing.T, input string, wantCount int) {
	t.Helper()
	checkCount(t, "bytes", input, wantCount)
}

func checkLines(t *testing.T, input string, wantCount int) {
	t.Helper()
	checkCount(t, "lines", input, wantCount)
}

func checkUTF8Letters(t *testing.T, input string, wantCount int) {
	t.Helper()
	checkCount(t, "letters", input, wantCount)
}

func checkUTF8LettersByteForByte(t *testing.T, input string, wantCount int) {
	t.Helper()
	checkCount(t, "letters", input, wantCount)

	b := []byte(input)
	var w utf8counter
	for _, b := range b {
		n, err := w.Write([]byte{b})
		check.Eq(t, n, 1)
		check.Eq(t, err, nil)
	}
	check.Eq(t, w.count, uint64(wantCount))
}
