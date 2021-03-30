package resources

import (
	"net/http"

	"github.com/JavDomGom/ste-go-twitter/config"
	"github.com/Sirupsen/logrus"
	"github.com/dghubble/go-twitter/twitter"
)

// SearchTweets look for tweets with some params and print the result.
func SearchTweets(log *logrus.Logger, client *twitter.Client, target string) (*twitter.Search, *http.Response, error) {
	var IncludeEntities = false

	log.Infof(
		"Looking for recent tweets with target %q in language %q.",
		target,
		config.TweetsLang,
	)

	searchParams := &twitter.SearchTweetParams{
		Query:           target,
		Lang:            config.TweetsLang,
		IncludeEntities: &IncludeEntities,
		Count:           100,
	}

	// See also: https://developer.twitter.com/en/docs/twitter-api/v1/tweets/search/api-reference/get-search-tweets

	return client.Search.Tweets(searchParams)
}
