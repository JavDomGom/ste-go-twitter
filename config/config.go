package config

var (
	LogPath                         = "log"
	MsgLenChunk                     = 2
	CharSet                         = "abcdefghijklmnopqrstuvwxyz0123456789 '.,:?!/-+=<>$_*()@#|%&[]{}^"
	SearchTweetParamLang            = "en"
	SearchTweetParamIncludeEntities = false
	SearchTweetParamCount           = 200
	SearchTweetParamResultType      = "mixed"
)

// GetCredentials returns Twitter API credentials.
func GetCredentials() (string, string, string, string) {
	consumerKey := "XXX"
	consumerSecret := "XXX"
	accessToken := "XXX"
	accessSecret := "XXX"

	return consumerKey, consumerSecret, accessToken, accessSecret
}
