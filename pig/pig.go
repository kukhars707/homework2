package pig

import (
	"strings"

	"github.com/pkg/errors"
)

var vowels = [5]string{"a", "e", "i", "o", "u"}
var consonants = [21]string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z"}

func PigLatin(word string) (string, error) {
	low := strings.ToLower(word)

	if low == "" || !isVowel(string(low[0])) && !isConsonant(string(low[0])) {
		return "", errors.New("Must be letter")
	}

	if isVowel(string(low[0])) {
		return word + "ay", nil
	} else if isConsonant(string(low[0])) {
		return updateConsonantString(low), nil
	}

	return low, nil

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
	vowelLookupTable := make(map[string]bool)
	for _, v := range vowels {
		vowelLookupTable[v] = true
	}

	return vowelLookupTable[x]
}

func isConsonant(x string) bool {
	consonantLookupTable := make(map[string]bool)
	for _, v := range consonants {
		consonantLookupTable[v] = true
	}

	return consonantLookupTable[x]

}
