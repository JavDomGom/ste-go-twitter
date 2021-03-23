package resources

import (
	"strings"
)

func ExtractWords(tweetText string) []string {
	var words []string
	for _, word := range strings.Fields(tweetText) {
		if len(word) < 4 || strings.Contains(word, "@") || strings.Contains(word, "/") {
			continue
		}
		words = append(words, word)
	}
	return words
}
