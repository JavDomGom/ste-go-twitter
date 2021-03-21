package config

var LogPath = "log"

// GetCredentials returns Twitter API credentials.
func GetCredentials() (string, string, string, string) {
	consumerKey := "XXX"
	consumerSecret := "XXX"
	accessToken := "XXX"
	accessSecret := "XXX"

	return consumerKey, consumerSecret, accessToken, accessSecret
}
