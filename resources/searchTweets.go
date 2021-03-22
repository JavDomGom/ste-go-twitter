package resources

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
)

// SearchTweets look for tweets with some params and print the result.
func SearchTweets(client *twitter.Client) {
	searchParams := &twitter.SearchTweetParams{
		Query:           "#cyber",
		Geocode:         "",
		Lang:            "en",
		Locale:          "",
		ResultType:      "recent",
		Count:           5,
		SinceID:         0,
		MaxID:           0,
		Until:           "",
		Since:           "",
		Filter:          "",
		IncludeEntities: new(bool),
		TweetMode:       "",
	}

	searchResult, _, _ := client.Search.Tweets(searchParams)

	for _, tweet := range searchResult.Statuses {
		fmt.Printf("Tweet: %+v\n", tweet.Text)
	}
}
