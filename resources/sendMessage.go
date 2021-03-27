package resources

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
)

func interact(words []string, code int64, hashtag string, client *twitter.Client) bool {
	target := words[code]

	if len(hashtag) > 0 {
		target += " #" + hashtag
	}

	fmt.Printf("- hashtag: %q, code: %v, target: %q\n", hashtag, code, target)

	searchResult, _, _ := SearchTweets(client, target)

	for i, tweet := range searchResult.Statuses {
		tweetID := tweet.ID

		fmt.Printf("\tTweet [%d] (%v): %+v\n", i, tweetID, tweet.Text)
		for _, word := range ProcessWords(tweet.Text) {
			if !Contains(word, words) {
				fmt.Printf("Word %q not in list of words, check next.\n", word)
				continue
			}

			fmt.Printf("Yeah! Word %q in list of words!\n", word)
			fmt.Printf("Checking if word %q is equal than target.\n", word)

			if word != target {
				fmt.Printf(
					"Word %q is not equal than target %q. Trying with another tweet.\n",
					word, target,
				)
				break
			}

			fmt.Printf("Yeah! Word %q is equal than taget %q.\n", word, target)
			fmt.Printf("Trying to retweet.\n")

			_, _, err := client.Statuses.Retweet(
				tweetID, &twitter.StatusRetweetParams{},
			)
			if err != nil {
				fmt.Printf("Tweet with id %q already retweeted.\n", tweetID)
				fmt.Printf("Error: %v\n", err)
				continue
			}
			fmt.Printf(
				"Tweet with id %+v containing the target %q successfully retweeted!\n",
				tweetID,
				target,
			)
			return true
		}
	}
	return false
}

func SendMessage(encodedMsg []int64, hashtags, words []string) {
	client := GetTwitterClient()

	// Verify credentials
	verifyParams := &twitter.AccountVerifyParams{
		IncludeEmail: twitter.Bool(true),
	}
	user, _, _ := client.Accounts.VerifyCredentials(verifyParams)
	fmt.Printf("Logged as: %+v\n", user.ScreenName)

	for _, code := range encodedMsg {
		for _, hashtag := range hashtags {
			if interact(words, code, hashtag, client) {
				break
			}
		}
	}
}
