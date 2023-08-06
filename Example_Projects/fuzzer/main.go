package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

// func Reverse(s string) string {
// 	b := []byte(s)
// 	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
// 		b[i], b[j] = b[j], b[i]
// 	}
// 	return string(b)
// }

// // replace above funtion to incoperate runes - Example_Projects/fuzzer/testdata/fuzz/FuzzReverse/01c1047d88ffbd570ab50a537504b15026c6748fcf8ed772efbc56059ae99d9f
// func Reverse(s string) string {
// 	r := []rune(s)
// 	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
// 		r[i], r[j] = r[j], r[i]
// 	}
// 	return string(r)
// }

// // changes made to fix - Example_Projects/fuzzer/testdata/fuzz/FuzzReverse/ac6733e1726c6863a5b63ceaa692e74c47e933c6bbe79346a1d92bd119ec1201
// func Reverse(s string) string {
// 	fmt.Printf("input: %q\n", s)
// 	r := []rune(s)
// 	fmt.Printf("runes: %q\n", r)
// 	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
// 		r[i], r[j] = r[j], r[i]
// 	}
// 	return string(r)
// }

// to fix the error - reverse_test.go:34: Before: "\x82", after: "�"
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

// original main
// func main() {
// 	input := "The quick brown fox jumped over the lazy dog"
// 	rev := Reverse(input)
// 	doubleRev := Reverse(rev)
// 	fmt.Printf("original: %q\n", input)
// 	fmt.Printf("reversed: %q\n", rev)
// 	fmt.Printf("reversed again: %q\n", doubleRev)
// }

// to accomadate additional errors
func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, revErr := Reverse(input)
	doubleRev, doubleRevErr := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
	fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)
}
