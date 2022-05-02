package ca_eau

import (
	"fmt"
	"strings"
	"unicode"
)

// DisplayRevert print each string in the array from the last one
// to the first one.
func DisplayRevert(ar []string) {
	for i := len(ar) - 1; i > 0; i-- {
		fmt.Println(ar[i])
	}
}

func isSubstring(s, text string, idx int) bool {
	// Loop with two indices, one for the string and one for the text
	// Stop when one of the two indices will
	for i, j := idx, 0; i <= len(text) && j < len(s); i, j = i+1, j+1 {
		if i >= len(text) {
			return false
		}
		if text[i] != s[j] {
			return false
		}
	}
	return true
}

// IsSubstringOf test if the string s in the given text.
func IsSubstringOf(txt, s string) bool {
	if s == "" {
		return true
	}
	fistChar := rune(s[0])
	for idx, c := range txt {
		if c != fistChar {
			continue
		}
		if isSubstring(s, txt, idx) {
			return true
		}
	}

	return false
}

// UpperOneOnTwo upper one character over two
func UpperOneOnTwo(txt string) string {
	var cptr int
	var builder strings.Builder

	// Check each char if it should be upper or not.
	// A char should be upper only if it is an alphabet char and
	// if its position is pair.
	for i := 0; i <= len(txt)-1; i++ {
		r := rune(txt[i])
		isAlphabet := 'a' <= r && r <= 'z' || 'A' <= r && r <= 'Z'
		if !isAlphabet {
			builder.WriteRune(r)
			continue
		}

		if cptr%2 == 0 && isAlphabet {
			builder.WriteRune(unicode.ToUpper(r))
		} else {
			builder.WriteRune(r)
		}
		cptr++
	}
	return builder.String()
}

// UpperFirstCharEachWord upper every first characters for each word in a string.
func UpperFirstCharEachWord(txt string) string {
	var builder strings.Builder
	var nextShouldUpper bool

	// Look each character of the text, if a space has been seen then
	// the current char should be upper, otherwise add the character
	// to the builder.
	var first bool = true
	for _, r := range txt {
		if nextShouldUpper || first {
			builder.WriteRune(unicode.ToUpper(r))
			nextShouldUpper = false
			first = false
			continue
		}
		if r == ' ' {
			nextShouldUpper = true
		}
		builder.WriteRune(r)
	}

	return builder.String()
}

// IsOnlyDigit run on each character and check if it is a number.
func IsOnlyDigit(s string) bool {
	for _, r := range s {
		if r < '0' || '9' < r {
			return false
		}
	}
	return true
}

// IsOnlyAlphabet check if a string is only composed of letters.
func IsOnlyAlphabet(s string) bool {
	for _, c := range s {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return true
}

// FindIndexOfString finds the index of a string.
func FindIndexOfString(list []string, toSearch string) int {
	for idx, s := range list {
		if s == toSearch {
			return idx
		}
	}
	return -1
}

func isBefore(s1, s2 string) bool {
	maxlen := len(s1)
	if maxlen < len(s2) {
		maxlen = len(s2)
	}

	// Check character by character
	for idx := 0; idx < maxlen; idx++ {
		r1, r2 := s1[idx], s2[idx]
		if r1 == r2 {
			continue
		}

		if r1 < r2 {
			return true
		}
		return false
	}

	// If every character are the same, use the length. For instance: "Bon" is before "Bonjour".
	return len(s1) < len(s2)
}

// SortString sorts in alphabetical order the given list of string
func SortString(toSort []string) []string {
	for cptr := len(toSort) - 1; cptr > 0; cptr-- {
		for i := 0; i < cptr; i++ {
			if isBefore(toSort[i], toSort[i+1]) {
				continue
			}
			toSort[i], toSort[i+1] = toSort[i+1], toSort[i]
		}
	}
	return toSort
}
