package main

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, World", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev := Reverse(orig)
		doubleRev := Reverse(rev)
		if orig != doubleRev {
			t.Error("Before: %q, After: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Error("Reverse produced invalid  UTF-8 string %q", rev)
		}
	})
}

// Look at this example to run through plenty of test cases
// func TestReverse(t *testing.T) {
// 	testCases := []struct {
// 		in, want string
// 	}{
// 		{"Hello, world", "dlrow ,olleH"},
// 		{" ", " "},
// 		{"!12345", "54321!"},
// 	}
// 	for _, tc := range testCases {
// 		rev := Reverse(tc.in)
// 		if rev != tc.want {
// 			t.Errorf("Reverse: %q, want%q", rev, tc.want)
// 		}
// 	}
// }
