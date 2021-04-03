package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/dghubble/go-twitter/twitter"

	"github.com/JavDomGom/ste-go-twitter/config"
	"github.com/JavDomGom/ste-go-twitter/resources"
)

func main() {
	file := resources.GetLogFile()
	defer file.Close()

	log := resources.ConfigLogger()
	log.SetOutput(file)

	log.Info("Starting program.")

	// Flags, commands and params config.
	sendCommand := flag.NewFlagSet("send", flag.ExitOnError)
	messageFlag := sendCommand.String("message", "", "Secret message to hide.")
	hashtagsFlag := sendCommand.String("hashtags", "", "List of hashtags to consider. (Optional)")

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

	// Prompts user for a password.
	pwd, err := resources.AskPassword()
	if err != nil {
		log.Fatal(err)
	}

	// Get client.
	client := resources.GetTwitterClient(log)

	// Verify credentials
	verifyParams := &twitter.AccountVerifyParams{
		IncludeEmail: twitter.Bool(true),
	}
	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Logged as: %+v\n", user.ScreenName)
	log.Infof("Logged as: %+v", user.ScreenName)

	words, err := resources.LoadWords(log, "./db/words.txt")
	if err != nil {
		log.Fatalf("LoadWords: %s", err)
	}
	log.Info("List of words loaded successfull!")

	pwdSHA256 := sha256.Sum256([]byte(pwd))
	pwdSHA256String := hex.EncodeToString(pwdSHA256[:])

	words = resources.GetShuffledWords(log, pwdSHA256String, words)

	if sendCommand.Parsed() {
		var hashtags []string

		if *messageFlag == "" {
			fmt.Println("Please supply the message to hide using -message option.")
			return
		}

		if *hashtagsFlag != "" {
			hashtags = strings.Split(*hashtagsFlag, ",")
		}

		encodedMsg := resources.GetEncodedMsg(
			log, strings.ToLower(*messageFlag), config.MsgLenChunk,
		)

		hashtags = append(hashtags, "")

		resources.SendMessage(log, client, encodedMsg, hashtags, words)
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

		resources.ReadMessage(log, words, *senderFlag, *retweetsFlag, client)
	}
}
