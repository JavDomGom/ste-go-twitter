package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"math/big"
	"math/rand"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	log "github.com/sirupsen/logrus"

	"github.com/JavDomGom/ste-go-twitter/config"
	"github.com/JavDomGom/ste-go-twitter/resources"
)

func main() {
	if _, err := os.Stat(config.LogPath); os.IsNotExist(err) {
		os.Mkdir(config.LogPath, 0744)
	}
	file, err := os.OpenFile(
		config.LogPath+"/ste-go-twitter.log",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.SetFormatter(&log.JSONFormatter{DisableHTMLEscape: true})
	log.SetLevel(log.DebugLevel)

	sendCommand := flag.NewFlagSet("send", flag.ExitOnError)
	messageFlag := sendCommand.String("message", "", "Secret message to hide.")

	recvCommand := flag.NewFlagSet("recv", flag.ExitOnError)
	senderFlag := recvCommand.String("sender", "", "Sender twitter account name, without @.")
	retweetsFlag := recvCommand.Int("retweets", 0, "Number of recent retweets to search hidden information.")

	if len(os.Args) == 1 {
		fmt.Printf("usage: %v <command> [options]\n\n", os.Args[0])
		fmt.Printf("The most commonly used git commands are:\n\n")
		fmt.Printf("\tsend\tTo send info as hidden messages.\n")
		fmt.Printf("\trecv\tTo recieve hidden messages.\n")
		return
	}

	switch os.Args[1] {
	case "send":
		sendCommand.Parse(os.Args[2:])
	case "recv":
		recvCommand.Parse(os.Args[2:])
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
	}

	if sendCommand.Parsed() {
		if *messageFlag == "" {
			fmt.Println("Please supply the message to hide using -message option.")
			return
		}
	}

	if recvCommand.Parsed() {
		if *senderFlag == "" {
			fmt.Println("Please supply the sender user using -sender option.")
			return
		}

		if *retweetsFlag == 0 {
			fmt.Println("Please supply the number of retweets using -retweets option.")
			return
		}

		fmt.Printf("senderFlag: %q\n", *senderFlag)
		fmt.Printf("retweetsFlag: %d\n", *retweetsFlag)
	}

	consumerKey, consumerSecret, accessToken, accessSecret := config.GetCredentials()

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Verify credentials
	verifyParams := &twitter.AccountVerifyParams{
		IncludeEmail: twitter.Bool(true),
	}
	user, _, _ := client.Accounts.VerifyCredentials(verifyParams)
	log.Printf("Logged as: %+v", user.ScreenName)

	// Prompts user for a password.
	pwd, err := resources.AskPassword()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pwdSHA256 := sha256.Sum256([]byte(pwd))
	log.Debugf("SHA256 [32]byte: %v", pwdSHA256)

	pwdSHA256String := hex.EncodeToString(pwdSHA256[:])
	log.Debugf("SHA256 string: %v", pwdSHA256String)

	log.Debug("Loading words from file.")
	words, err := resources.LoadWords("./db/words.txt")

	if err != nil {
		log.Errorf("LoadWords: %s", err)
	}

	log.Debug("Cutting SHA256 string in 8 chunks of 4 bytes and use it to seed-shuffle list of words.")
	c := 1
	for i := 0; i < 64; i += 8 {
		pwdSHA256BigInt := new(big.Int)
		pwdSHA256BigInt.SetString(pwdSHA256String[i:i+8], 16)
		pwdSHA256Int64 := pwdSHA256BigInt.Int64()
		log.Debugf(
			"Chunk %d: %v => %v (%T)",
			c,
			pwdSHA256String[i:i+8],
			pwdSHA256Int64,
			pwdSHA256Int64,
		)
		rand.Seed(pwdSHA256Int64)
		rand.Shuffle(len(words), func(i, j int) {
			words[i], words[j] = words[j], words[i]
		})
		c++
	}

	// Search and print some tweets.
	resources.SearchTweets(client)
}
