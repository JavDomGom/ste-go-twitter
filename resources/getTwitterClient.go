package resources

import (
	"github.com/JavDomGom/ste-go-twitter/config"
	"github.com/Sirupsen/logrus"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func GetTwitterClient(log *logrus.Logger) *twitter.Client {
	log.Info("Getting HTTP twitter client.")
	consumerKey, consumerSecret, accessToken, accessSecret := config.GetCredentials()

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	return twitter.NewClient(httpClient)
}
