package main

import (
	"reflect"
	"strconv"
	"testing"
	"time"
)

type User struct {
	ID           int
	FirstName    string
	LastName     string
	Age          int
	RegisteredAt time.Time `dump:"02.01.2006"`
}

func (u User) DumpDirect() string {
	return "ID:" + strconv.Itoa(u.ID) + "\n" +
		"FirstName:" + u.FirstName + "\n" +
		"LastName:" + u.LastName + "\n" +
		"Age:" + strconv.Itoa(u.Age) + "\n" +
		"RegisteredAt:" + u.RegisteredAt.Format("02.01.2006") + "\n"
}

func (u User) DumpReflect() string {
	dump := ""
	r := reflect.ValueOf(u)
	for i := 0; i < r.NumField(); i++ {
		if !r.Field(i).CanInterface() {
			continue
		}
		field := r.Type().Field(i).Name
		val := ""

		switch v := r.Field(i).Interface().(type) {
		case string:
			val = v
		case int:
			val = strconv.Itoa(v)
		case time.Time:
			val = v.Format(r.Type().Field(i).Tag.Get("dump"))
		}
		dump += field + ":" + val + "\n"
	}
	return dump
}

func BenchmarkReflect(b *testing.B) {
	u := User{
		ID:           1,
		FirstName:    "Alice",
		LastName:     "Bobs",
		Age:          18,
		RegisteredAt: time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
	}
	expected := "ID:1\nFirstName:Alice\nLastName:Bobs\nAge:18\nRegisteredAt:31.12.2022\n"

	b.Run("direct", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			actual := u.DumpDirect()
			if actual != expected {
				b.Fatalf("actual '%s' != expected '%s'\n", actual, expected)
			}
		}
	})
	b.Run("reflect", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			actual := u.DumpReflect()
			if actual != expected {
				b.Fatalf("actual '%s' != expected '%s'\n", actual, expected)
			}
		}
	})
}
