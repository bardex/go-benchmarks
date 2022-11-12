package main

import (
	"strings"
	"testing"
)

type Author struct {
	id        int
	lastname  string
	firstname string
}

type Book struct {
	index     []int
	maps      map[string]string
	authors   []Author
	comments1 string
	comments2 string
	comments3 string
	comments4 string
}

func (b *Book) GetContentFast() string {
	return b.comments1
}

func (b Book) GetContentSlow() string {
	return b.comments1
}

func newBook() Book {
	book := Book{
		index:     make([]int, 1000, 1000),
		authors:   make([]Author, 1000, 1000),
		maps:      make(map[string]string, 1000),
		comments1: strings.Repeat("Wow!", 5000000),
		comments2: strings.Repeat("Wow!", 5000000),
		comments3: strings.Repeat("Wow!", 5000000),
		comments4: strings.Repeat("Wow!", 5000000),
	}
	return book
}

func BenchmarkReceiver(b *testing.B) {
	// make large book
	book := newBook()
	b.ResetTimer()
	b.Run("pointer receiver", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			book.GetContentFast()
		}
	})
	b.Run("regular receiver", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			book.GetContentSlow()
		}
	})
}
