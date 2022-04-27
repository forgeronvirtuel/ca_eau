package water

import (
	"fmt"
	"strings"
)

func generateCombination(a, b, c rune, combination []string) []string {
	// If one of the rune is equal to another => left
	if a == b || a == c || b == c {
		return combination
	}

	// Check if a combination already containing those 3 runes exists
	for _, s := range combination {
		if strings.ContainsRune(s, a) && strings.ContainsRune(s, b) && strings.ContainsRune(s, c) {
			return combination
		}
	}

	// Append the new combination
	combination = append(combination, fmt.Sprintf("%c%c%c", a, b, c))

	return combination
}

// GenerateCombination generates all combination of three digit (from 0 to 9)
// with each digit uniq in the combination. For instance: 021 does not appear
// because 012 exists, and 020 does not appear because 012 exists.
func GenerateCombination() []string {
	var combination []string
	for a := '0'; a <= '9'; a++ {
		for b := '0'; b <= '9'; b++ {
			for c := '0'; c <= '9'; c++ {
				combination = generateCombination(a, b, c, combination)
			}
		}
	}

	return combination
}

// GenerateStringTupleCombination Generates strings with tuples from 00 to 99
// and such that the second element starts with the first one + 1. For instance
// there is the following series 97 98, 97 99, 98 99 but there is not 97 01,
// 97 02, ...
func GenerateStringTupleCombination() {
	for a := 0; a <= 99; a++ {
		for b := a + 1; b <= 99; b++ {
			fmt.Printf("%02d %02d\n", a, b)
		}
	}
}
