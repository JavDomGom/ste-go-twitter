package resources

import (
	"strings"
)

func ProcessWords(tweetText string) []string {
	var words []string
	var replacer = strings.NewReplacer(
		"#", "",
		".", "",
		",", "",
		"…", "",
		"?", "",
		"!", "",
		"-", " ",
		"\"", "",
	)

	tweetText = strings.ToLower(replacer.Replace(tweetText))

	for _, word := range strings.Fields(tweetText) {
		if len(word) < 4 || strings.Contains(word, "@") || strings.Contains(word, "/") {
			continue
		}
		words = append(words, word)
	}
	return words
}

func Contains(element string, slice []string) bool {
	for _, s := range slice {
		if s == element {
			return true
		}
	}

	return false
}
