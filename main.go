package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"strings"

	"github.com/Sirupsen/logrus"

	"github.com/JavDomGom/ste-go-twitter/config"
	"github.com/JavDomGom/ste-go-twitter/resources"
)

func main() {
	var logger = logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{DisableHTMLEscape: true})
	logger.SetLevel(logrus.DebugLevel)

	if _, err := os.Stat(config.LogPath); os.IsNotExist(err) {
		os.MkdirAll(config.LogPath, 0744)
	}
	file, err := os.OpenFile(
		config.LogPath+"/ste-go-twitter.log",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)
	if err != nil {
		logger.Fatal(err)
	}
	defer file.Close()

	logger.SetOutput(file)
	log := resources.Log{
		Logger: logger,
	}

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

	var hashtags []string

	if sendCommand.Parsed() {
		if *messageFlag == "" {
			fmt.Println("Please supply the message to hide using -message option.")
			return
		}

		if *hashtagsFlag != "" {
			hashtags = strings.Split(*hashtagsFlag, ",")
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

	encodedMsg := resources.GetEncodedMsg(strings.ToLower(*messageFlag), config.MsgLenChunk)
	logger.Debugf("encodedMsg is %v", encodedMsg)

	// Prompts user for a password.
	pwd, err := resources.AskPassword()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pwdSHA256 := sha256.Sum256([]byte(pwd))
	logger.Debugf("SHA256 [32]byte: %v", pwdSHA256)

	pwdSHA256String := hex.EncodeToString(pwdSHA256[:])
	logger.Debugf("SHA256 string: %v", pwdSHA256String)

	logger.Debug("Loading words from file.")
	words, err := log.LoadWords("./db/words.txt")
	if err != nil {
		logger.Errorf("LoadWords: %s", err)
	}

	logger.Debug("Cutting password SHA256 string in chunks of 4 bytes and use it to seed-shuffle list of words.")
	c := 1
	for i := 0; i < 64; i += 8 {
		pwdSHA256BigInt := new(big.Int)
		pwdSHA256BigInt.SetString(pwdSHA256String[i:i+8], 16)
		pwdSHA256Int64 := pwdSHA256BigInt.Int64()
		logger.Debugf(
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

	resources.SendMessage(encodedMsg, append(hashtags, ""), words)
}
