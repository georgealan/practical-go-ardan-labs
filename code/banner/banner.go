package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	banner("Go", 6)
	banner("G☺", 6)
	fmt.Println()

	// Test string length with and without emoji
	s1 := "Go"
	s2 := "G☺"
	fmt.Println("string Go length:", len(s1))
	fmt.Println("string G☺ length in byte (uint8):", len(s2))
	fmt.Println("string G☺ length in rune (int32):", utf8.RuneCountInString(s2))
	fmt.Println()
	// code point is = rune ~=  unicode character

	/*
		strings in Go have duality, they have sequences of bytes and collections of rune and characters
		Have two types:
		byte (uint8)
		rune (int32)
	*/
	b := s2[0]
	fmt.Printf("%c of type %T\n", b, b) // byte (uint8)
	fmt.Println()

	for i, r := range s2 {
		fmt.Println(i, r)
		if i == 0 {
			fmt.Printf("%c of type %T\n", r, r) // rune (int32)
		}
	}

	fmt.Println()
	x, y := 1, "1"
	fmt.Printf("x=%v, y=%v\n", x, y)
	fmt.Printf("x=%#v, y=%#v\n", x, y) // Use #v in debug/log

	fmt.Println("g", isPalindrome("g"))
	fmt.Println("go", isPalindrome("go"))
	fmt.Println("gog", isPalindrome("gog"))
	fmt.Println("g☺g", isPalindrome("g☺g"))
	fmt.Println("gogo", isPalindrome("gogo"))
}

// isPalindrome("g") -> true
// isPalindrome("go") -> false
// isPalindrome("gog") -> true
// isPalindrome("gogo") -> false
func isPalindrome(s string) bool {
	rs := []rune(s) // Get slice of runes out of s, for get unicode
	for i := 0; i < len(rs)/2; i++ {
		if rs[i] != rs[len(rs)-i-1] {
			return false
		}
	}
	return true
}

func banner(text string, width int) {
	padding := (width - utf8.RuneCountInString(text)) / 2
	// padding := (width - len(text)) / 2 // BUG: len is in bytes
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}
