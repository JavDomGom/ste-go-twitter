package resources

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
)

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
			target := words[code]

			if len(hashtag) > 0 {
				target += " #" + hashtag
			}

			fmt.Printf("- hashtag: %q, code: %v, target: %q\n", hashtag, code, target)

			searchResult, _, _ := SearchTweets(client, target)

			for i, tweet := range searchResult.Statuses {
				fmt.Printf("\tTweet [%d] (%v): %+v\n", i, tweet.IDStr, tweet.Text)
				for c, word := range ExtractWords(tweet.Text) {
					fmt.Printf("word[%d]: %v\n", c, word)
				}
			}
		}
	}
}
