package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

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

	cleanCommand := flag.NewFlagSet("clean", flag.ExitOnError)
	userFlag := cleanCommand.String("user", "", "My account name, without @.")
	countFlag := cleanCommand.Int("count", 0, "Number of recent retweets to undo retweet.")

	if len(os.Args) == 1 {
		fmt.Printf("usage: %v <command> [options]\n\n", os.Args[0])
		fmt.Printf("The most commonly used git commands are:\n\n")
		fmt.Printf("\tsend\tTo send info as hidden messages.\n")
		fmt.Printf("\trecv\tTo recieve hidden messages.\n")
		fmt.Printf("\tclean\tTo undo a specific number of retweets from my account.\n")
		return
	}

	switch os.Args[1] {
	case "send":
		sendCommand.Parse(os.Args[2:])
	case "recv":
		recvCommand.Parse(os.Args[2:])
	case "clean":
		cleanCommand.Parse(os.Args[2:])
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
	}

	// Get client.
	client, err := resources.GetTwitterClient(log)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Got client!")

	if cleanCommand.Parsed() {
		if *userFlag == "" {
			fmt.Println("Please supply your user using -user option.")
			return
		}

		if *countFlag == 0 {
			fmt.Println("Please supply the number of retweets to undo retweet using -count option.")
			return
		}

		err := resources.UndoRetweet(log, *userFlag, *countFlag, client)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	// Prompts user for a password.
	pwd, err := resources.AskPassword()
	if err != nil {
		log.Fatal(err)
	}

	// Load list of words.
	words, err := resources.LoadWords(log, "./db/words.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Info("List of words loaded successfull!")

	// Seed-Shuffle list of words.
	resources.ShuffleWords(log, pwd, words)

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
