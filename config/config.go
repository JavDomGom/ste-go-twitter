package config

var LogPath = "log"
var MsgLenChunk = 2
var CharSet = "abcdefghijklmnopqrstuvwxyz0123456789 '.,:?!/-+=<>$_*()@#|%&[]{}^"
var TweetsLang = "en"

// GetCredentials returns Twitter API credentials.
func GetCredentials() (string, string, string, string) {
	consumerKey := "XXX"
	consumerSecret := "XXX"
	accessToken := "XXX"
	accessSecret := "XXX"

	return consumerKey, consumerSecret, accessToken, accessSecret
}
