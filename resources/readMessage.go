package resources

import (
	"fmt"
	"strings"

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
	var (
		interactions  []int
		secretMessage string
	)

	log.Debugf("Looking for retweets in %q user timeline.", senderTwitterUser)
	retweets, _, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		ScreenName: senderTwitterUser,
		Count:      count,
	})
	if err != nil {
		log.Fatal(err)
	}

	if len(retweets) == 0 {
		fmt.Printf("%v's timeline seems empty. Contact the user.\n", senderTwitterUser)
		log.Fatalf("%v's timeline seems empty. Contact the user.", senderTwitterUser)
	}

	for _, retweet := range ReverseSlice(retweets) {
		log.Infof("Processing retweet with ID %+v\n", retweet.ID)

		for _, word := range ProcessWords(log, retweet.Text) {
			log.Infof("Checking if word %q is in list of words.", word)

			isInList, index := InListOfWords(log, word, words)

			if isInList {
				log.Debugf(
					"Appending seq %v from word %q to list of interactions", index, word,
				)

				interactions = append(interactions, index)

				break
			}
		}
	}

	log.Debugf("interactions: %v", interactions)

	log.Info("Unhiding secret message.")
	for _, code := range interactions {
		if code == -1 {
			continue
		}

		secretMessage += CodeToString(log, code)
	}

	fmt.Printf("Secret message: %q\n", strings.TrimRight(secretMessage, "\t \n"))
}
