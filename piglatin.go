package main

import "strings"

const (
	pigLatinSuffix             string = "ay"
	firstLetterExceptions      string = "aeiou"
	firstLetterExceptionSuffix string = "y" + pigLatinSuffix
)

// Encode translates one or more english words into the PigLatin equivalent
func Encode(in string) string {
	var pigLatinWords []string
	englishWords := strings.Split(in, " ")

	for _, word := range englishWords {
		first := word[0:1]
		if strings.Contains(firstLetterExceptions, first) {
			pigLatinWords = append(pigLatinWords, word+firstLetterExceptionSuffix)
		} else {
			var index int
			for pos, char := range word {
				if strings.ContainsRune(firstLetterExceptions, char) {
					index = pos
					break
				}
			}
			pigLatinWords = append(pigLatinWords, word[index:]+word[0:index]+pigLatinSuffix)
		}
	}
	return strings.Join(pigLatinWords, " ")
}