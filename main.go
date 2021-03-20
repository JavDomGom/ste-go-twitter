package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"os"

	"github.com/JavDomGom/ste-go-twitter/config"
	"github.com/JavDomGom/ste-go-twitter/resources"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
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
		fmt.Printf("Your plain text message are: %q\n", *messageFlag)
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
	fmt.Printf("Logged as: %+v\n", user.ScreenName)

	// Prompts user for a password.
	pwd, err := resources.AskPassword()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pwdSHA256 := sha256.Sum256([]byte(pwd))
	pwdSHA256String := hex.EncodeToString(pwdSHA256[:])

	fmt.Printf("SHA256 [32]byte: %v\n", pwdSHA256)
	fmt.Printf("SHA256 string:\t %v\n", pwdSHA256String)

	// Load words database.
	words, err := resources.LoadWords("./db/words.txt")

	if err != nil {
		log.Fatalf("LoadWords: %s", err)
	}

	for i := 0; i < 64; i += 8 {
		pwdSHA256BigInt := new(big.Int)
		pwdSHA256BigInt.SetString(pwdSHA256String[i:i+8], 16)
		pwdSHA256Int64 := pwdSHA256BigInt.Int64()
		fmt.Printf(
			"%T[%v]:\t (%v) %v\n",
			pwdSHA256Int64,
			i,
			pwdSHA256String[i:i+8],
			pwdSHA256Int64,
		)
		rand.Seed(pwdSHA256Int64)
		rand.Shuffle(len(words), func(i, j int) {
			words[i], words[j] = words[j], words[i]
		})
	}

	fmt.Println(words)

	// Search and print some tweets.
	resources.SearchTweets(client)
}
