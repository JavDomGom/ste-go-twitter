package resources

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/dghubble/go-twitter/twitter"
)

func interact(
	log *logrus.Logger,
	words []string,
	code int64,
	hashtag string,
	tweetCount int,
	client *twitter.Client,
) bool {
	target := words[code]

	if len(hashtag) > 0 {
		target += " #" + hashtag
	}

	searchResult, _, err := SearchTweets(log, client, target)
	if err != nil {
		log.Fatal(err)
	}

	for i, tweet := range searchResult.Statuses {
		tweetID := tweet.ID

		log.Debugf("Tweet [%d] (%v): %+v", i, tweetID, tweet.Text)

		for _, word := range ProcessWords(log, tweet.Text) {

			log.Infof("Checking if word %q is in list of words.", word)
			isInList, _ := InListOfWords(log, word, words)
			if !isInList {
				log.Infof("Word %q is not in list of words. Continue and try with the next word.", word)
				continue
			}

			log.Infof("Checking if word %q is equal than word %q from list.", word, words[code])
			if word != words[code] {
				log.Infof(
					"Word %q is not equal than %q. Break and try with another tweet.",
					word, words[code],
				)
				break
			}
			log.Infof("Yeah! Word %q is equal than %q.", word, words[code])

			log.Info("Trying to retweet.")
			_, _, err := client.Statuses.Retweet(
				tweetID, &twitter.StatusRetweetParams{},
			)
			if err != nil {
				log.Infof(
					"Tweet with ID %+v already retweeted. Continue an try with another tweet.",
					tweetID,
				)
				log.Info(err)
				continue
			}

			fmt.Printf(
				"Tweet %03d with ID %+v containing the target %q successfully retweeted!\n",
				tweetCount,
				tweetID,
				target,
			)
			return true
		}
	}
	return false
}

func SendMessage(
	log *logrus.Logger,
	words,
	hashtags []string,
	encodedMsg []int64,
) {
	// Get client.
	client, err := GetTwitterClient(log)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Got client!")

	tweetCount := 1

	for _, code := range encodedMsg {
		for _, hashtag := range hashtags {
			log.Debugf("hashtag: %q, code %v.", hashtag, code)

			if interact(log, words, code, hashtag, tweetCount, client) {
				tweetCount++
				break
			}
		}
	}
}
