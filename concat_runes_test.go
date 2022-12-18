package main

import (
	"strings"
	"testing"
)

func BenchmarkConcatRunes(b *testing.B) {
	N := 100
	c := 'A'
	exp := strings.Repeat(string(c), N)

	b.ResetTimer()

	b.Run("by slice", func(b *testing.B) {
		s := make([]byte, 0)
		for i := 0; i < N; i++ {
			s = append(s, byte(c))
		}
		res := string(s)
		if res != exp {
			b.Fatalf("Failed, exp:%s, actual:%s", exp, res)
		}
	})

	b.Run("by builder", func(b *testing.B) {
		s := strings.Builder{}
		for i := 0; i < N; i++ {
			s.WriteRune(c)
		}
		res := s.String()
		if res != exp {
			b.Fatalf("Failed, exp:%s, actual:%s", exp, res)
		}
	})
}
