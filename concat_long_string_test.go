package main

import (
	"strings"
	"testing"
)

func BenchmarkConcatOneLongString(b *testing.B) {
	NRows := 10
	row := "[2020-01-01 12:00:00] GET /get/users/group/1 200 80.52 \n"
	expected := strings.Repeat(row, NRows)

	b.Run("concat", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			actual := ""
			for k := 0; k < NRows; k++ {
				actual += row
			}
			if actual != expected {
				b.Fatalf("actual '%s' != expected '%s'\n", actual, expected)
			}
		}
	})
	b.Run("builder", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			builder := strings.Builder{}
			for k := 0; k < NRows; k++ {
				builder.WriteString(row)
			}
			actual := builder.String()
			if actual != expected {
				b.Fatalf("actual '%s' != expected '%s'\n", actual, expected)
			}
		}
	})
}
