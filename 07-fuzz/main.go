package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	// To handle possible error return.
	rev, revErr := Reverse(input)
	doubleRev, doubleRevErr := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
	fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)
}

// Reverse func with bug: size of byte.
// func Reverse(s string) string {
// 	b := []byte(s)
// 	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
// 		b[i], b[j] = b[j], b[i]
// 	}
// 	return string(b)
// }

// Fixed the first bug.
// But was included another one: wrong conversion UTF-8. Go encodes the
// byte slice to UTF-8, and replaces the byte with the UTF-8 character �
// func Reverse(s string) string {
// 	// This will help us understand what is going wrong when converting
// 	// the string to a slice of runes.
// 	fmt.Printf("input: %q\n", s)
// 	r := []rune(s)
// 	fmt.Printf("runes: %q\n", r)
// 	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
// 		r[i], r[j] = r[j], r[i]
// 	}
// 	return string(r)
// }

// Fixed the second bug.
func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}
