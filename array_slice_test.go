package main

import (
	"testing"
)

func BenchmarkIndexedArrays(b *testing.B) {
	const arraySize = 1000_000
	b.Run("indexed array", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			items := [arraySize]bool{}
			for n := 0; n < arraySize; n++ {
				items[n] = true
			}
		}
	})
	b.Run("indexed slice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			items := make([]bool, arraySize, arraySize)
			for n := 0; n < arraySize; n++ {
				items[n] = true
			}
		}
	})
	b.Run("indexed map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			items := make(map[int]bool, arraySize)
			for n := 0; n < arraySize; n++ {
				items[n] = true
			}
		}
	})
	b.Run("indexed map no cap", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			items := make(map[int]bool)
			for n := 0; n < arraySize; n++ {
				items[n] = true
			}
		}
	})
}
