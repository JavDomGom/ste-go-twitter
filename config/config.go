package config

var (
	LogPath     = "log"
	MsgLenChunk = 2
	CharSet     = "abcdefghijklmnopqrstuvwxyz0123456789 '.,:?!/-+=<>$_*()@#|%&[]{}^"

	SearchTweetParamLang            = "en"
	SearchTweetParamIncludeEntities = false
	SearchTweetParamCount           = 100
	SearchTweetParamResultType      = "mixed"
	// See also: https://developer.twitter.com/en/docs/twitter-api/v1/tweets/search/api-reference/get-search-tweets
)

// GetCredentials returns Twitter API credentials.
func GetCredentials() (string, string, string, string) {
	consumerKey := "XXX"
	consumerSecret := "XXX"
	accessToken := "XXX"
	accessSecret := "XXX"

	return consumerKey, consumerSecret, accessToken, accessSecret
}
