package main

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkConcatManyShortStrings(b *testing.B) {
	s := []string{
		"2020-01-01 12:00:00",
		"GET",
		"/get/users/group/1",
		"200",
		"80.52",
	}
	expected := "2020-01-01 12:00:00 GET /get/users/group/1 200 80.52"

	b.Run("concat", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			actual := s[0] + " " + s[1] + " " + s[2] + " " + s[3] + " " + s[4]
			if actual != expected {
				b.Fatalf("actual '%s' != expected '%s'\n", actual, expected)
			}
		}
	})
	b.Run("builder", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			builder := strings.Builder{}
			builder.WriteString(s[0])
			builder.WriteString(" ")
			builder.WriteString(s[1])
			builder.WriteString(" ")
			builder.WriteString(s[2])
			builder.WriteString(" ")
			builder.WriteString(s[3])
			builder.WriteString(" ")
			builder.WriteString(s[4])
			actual := builder.String()
			if actual != expected {
				b.Fatalf("actual '%s' != expected '%s'\n", actual, expected)
			}
		}
	})
	b.Run("join", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			actual := strings.Join(s, " ")
			if actual != expected {
				b.Fatalf("actual '%s' != expected '%s'\n", actual, expected)
			}
		}
	})
	b.Run("fmt", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			actual := fmt.Sprintf("%s %s %s %s %s", s[0], s[1], s[2], s[3], s[4])
			if actual != expected {
				b.Fatalf("actual '%s' != expected '%s'\n", actual, expected)
			}
		}
	})
}
