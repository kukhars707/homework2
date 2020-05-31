package pig

import (
	"strings"
)

func PigLatin(word string) string {
	low, err := strings.ToLower(word), "Must be letter"

	if low == "" || !isVowel(string(low[0])) && !isConsonant(string(low[0])) {
		return err
	}

	if isVowel(string(low[0])) {
		return word + "ay"
	} else if isConsonant(string(low[0])) {
		return updateConsonantString(low)
	}

	return low

}

func updateConsonantString(w string) string {
	consunant := ""
	for _, l := range w {
		if isVowel(string(l)) {
			break
		}
		consunant += string(l)
	}

	t := len(consunant)

	return w[t:] + consunant + "ay"
}

func isVowel(x string) bool {
	vowels := [5]string{"a", "e", "i", "o", "u"}
	vowelLookupTable := make(map[string]bool)
	for _, v := range vowels {
		vowelLookupTable[v] = true
	}

	return vowelLookupTable[x]
}

func isConsonant(x string) bool {
	consonants := [21]string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z"}
	consonantLookupTable := make(map[string]bool)
	for _, v := range consonants {
		consonantLookupTable[v] = true
	}

	return consonantLookupTable[x]

}
