package resources

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/dghubble/go-twitter/twitter"
)

func ReadMessage(
	log *logrus.Logger,
	words []string,
	senderTwitterUser string,
	count int,
	client *twitter.Client,
) {
	log.Debugf("Looking for retweets in %q user timeline.", senderTwitterUser)
	retweets, _, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		ScreenName: senderTwitterUser,
		Count:      count,
	})
	if err != nil {
		log.Fatal(err)
	}

	for i, retweet := range ReverseSlice(retweets) {
		fmt.Printf("Retweet [%d] (%+v): %+v\n", i, retweet.ID, retweet.Text)
		for _, word := range ProcessWords(log, retweet.Text) {
			log.Infof("Checking if word %q is in list of words.", word)
			if InListOfWords(log, word, words) {
				break
			}
		}
	}
}
