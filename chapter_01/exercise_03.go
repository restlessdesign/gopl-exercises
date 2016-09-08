package main

import (
	"strings"
	"testing"
)

// @todo
// Get this running properly

var words = []string{"foo", "bar", "baz", "qux"}

func StringConcat(words []string) string {
	s, sep := "", ""

	for _, arg := range words {
		s += sep + arg
		sep = " "
	}

	return s
}

func StringJoin(words []string) string {
	s := strings.Join(words, " ")

	return s
}

func BenchmarkStringConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringConcat(words)
	}
}

func BenchmarkStringJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringJoin(words)
	}
}
