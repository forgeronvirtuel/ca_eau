package water

import (
	"fmt"
	"strings"
	"unicode"
)

// DisplayRevert Print each string in the array from the last one
// to the first one
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

// IsSubstringOf test if the string s in the given text
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
			fmt.Println(true)
			return true
		}
	}

	return false
}

// UpperOneOnTwo Upper one character over two
func UpperOneOnTwo(txt string) string {
	var cptr int
	var builder strings.Builder
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

// UpperFirstCharEachWord upper every first characters for each word in a string
func UpperFirstCharEachWord(txt string) string {
	var builder strings.Builder
	var nextShouldUpper bool

	// Look each character of the text, if a space has been seen then
	// the current char should be upper, otherwise add the character
	// to the builder.
	for _, r := range txt {
		if nextShouldUpper {
			builder.WriteRune(unicode.ToUpper(r))
			nextShouldUpper = false
			continue
		}
		if r == ' ' {
			nextShouldUpper = true
		}
		builder.WriteRune(r)
	}

	return builder.String()
}
