package main

import "testing"

func TestUrlJoin(t *testing.T) {
	cases := []struct {
		a string
		b string
		e string
	}{
		{a: "http://asd.com", b: "file.json", e: "http://asd.com/file.json"},
		{a: "http://asd.com/", b: "file.json", e: "http://asd.com/file.json"},
		{a: "http://asd.com/", b: "/file.json", e: "http://asd.com/file.json"},
		{a: "http://asd.com", b: "/file.json", e: "http://asd.com/file.json"},
		{a: "http://asd.com", b: "\\file.json", e: "http://asd.com/file.json"},
	}
	for _, c := range cases {
		got := urlJoin(c.a, c.b)
		if got != c.e {
			t.Errorf("Expected '%v', got '%v'", c.e, got)
		}
	}
}
