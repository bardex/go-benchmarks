package main

import (
	"fmt"
	"testing"
)

func getStructMap(in map[int]struct{}) {
	for k, v := range in {
		if v != struct{}{} {
			panic(fmt.Sprintf("invaild value in %d", k))
		}
	}
}

func getBoolMap(in map[int]bool) {
	for k, v := range in {
		if v != true {
			panic(fmt.Sprintf("invaild value in %d", k))
		}
	}
}

func BenchmarkMaps(b *testing.B) {
	const arraySize = 1000_000
	b.Run("bool map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			items := make(map[int]bool)
			for n := 0; n < arraySize; n++ {
				items[n] = true
			}
			getBoolMap(items)
		}
	})
	b.Run("struct map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			items := make(map[int]struct{})
			for n := 0; n < arraySize; n++ {
				items[n] = struct{}{}
			}
			getStructMap(items)
		}
	})
}
