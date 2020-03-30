package main

import (
	"github.com/gonutz/check"
	"testing"
)

func TestUnescapeString(t *testing.T) {
	check.Eq(t, unescape(`a\r\n\\\t\x`), "a\r\n\\\t\\x")
}
