package resources

import (
	"fmt"

	"github.com/JavDomGom/ste-go-twitter/config"
	"github.com/Sirupsen/logrus"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func GetTwitterClient(log *logrus.Logger) (*twitter.Client, error) {
	log.Info("Getting HTTP twitter client.")
	consumerKey, consumerSecret, accessToken, accessSecret := config.GetCredentials()

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	// HTTP client will automatically authorize Requests.
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client.
	client := twitter.NewClient(httpClient)

	// Verify credentials
	verifyParams := &twitter.AccountVerifyParams{
		IncludeEmail: twitter.Bool(true),
	}
	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return client, err
	}

	fmt.Printf("Logged as: %+v\n", user.ScreenName)
	log.Infof("Logged as: %+v", user.ScreenName)

	return client, nil
}
