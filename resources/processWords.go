package resources

import (
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/dghubble/go-twitter/twitter"
)

func ProcessWords(log *logrus.Logger, tweetText string) []string {
	log.Info("Processing words in tweet.")

	var (
		words    []string
		replacer = strings.NewReplacer(
			"#", "",
			".", "",
			",", "",
			"…", "",
			"?", "",
			"!", "",
			"-", " ",
			"\"", "",
		)
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

func InListOfWords(log *logrus.Logger, word string, words []string) (bool, int) {
	for i, w := range words {
		if w == word {
			log.Infof("Yeah! Word %q in list of words with index %v!", word, i)
			return true, i
		}
	}

	log.Infof("Word %q not in list of words, check next.", word)
	return false, -1
}

func ReverseSlice(s []twitter.Tweet) []twitter.Tweet {
	if len(s) < 2 {
		return s
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}
