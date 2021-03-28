package resources

import (
	"net/http"

	"github.com/JavDomGom/ste-go-twitter/config"
	"github.com/Sirupsen/logrus"
	"github.com/dghubble/go-twitter/twitter"
)

// SearchTweets look for tweets with some params and print the result.
func SearchTweets(log *logrus.Logger, client *twitter.Client, target string) (*twitter.Search, *http.Response, error) {
	log.Infof(
		"Looking for recent tweets with target %q in language %q.",
		target,
		config.TweetsLang,
	)
	searchParams := &twitter.SearchTweetParams{
		Query:           target,
		Geocode:         "",
		Lang:            config.TweetsLang,
		Locale:          "",
		ResultType:      "recent",
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
