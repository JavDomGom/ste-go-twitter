package resources

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/dghubble/go-twitter/twitter"
)

func UndoRetweet(
	log *logrus.Logger,
	user string,
	count int,
	client *twitter.Client,
) error {
	log.Debugf("Looking for retweets in %q user timeline.", user)
	retweets, _, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		ScreenName: user,
		Count:      count,
	})
	if err != nil {
		return err
	}

	if len(retweets) == 0 {
		fmt.Printf("%v's timeline seems empty.\n", user)
		log.Fatalf("%v's timeline seems empty.", user)
	}

	for _, retweet := range ReverseSlice(retweets) {
		log.Infof("Undoing retweet with ID %+v", retweet.ID)

		_, _, err := client.Statuses.Destroy(
			retweet.ID,
			nil,
		)
		if err != nil {
			return err
		}
		fmt.Printf("Retweet with ID %+v successfully undoing.\n", retweet.ID)
	}

	return nil
}
