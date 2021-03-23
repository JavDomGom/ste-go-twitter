package resources

import (
	"net/http"

	"github.com/JavDomGom/ste-go-twitter/config"
	"github.com/dghubble/go-twitter/twitter"
)

// SearchTweets look for tweets with some params and print the result.
func SearchTweets(client *twitter.Client, target string) (*twitter.Search, *http.Response, error) {
	searchParams := &twitter.SearchTweetParams{
		Query:           target,
		Geocode:         "",
		Lang:            config.TweetsLang,
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

	// See also: https://developer.twitter.com/en/docs/twitter-api/v1/tweets/search/api-reference/get-search-tweets

	return client.Search.Tweets(searchParams)
}
