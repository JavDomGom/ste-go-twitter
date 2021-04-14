<p align="center"><img src="https://github.com/JavDomGom/ste-go-twitter/blob/main/img/ste-go-twitter_logo.png"></p>

## Status

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-brightgreen.svg)](https://www.gnu.org/licenses/gpl-3.0)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/JavDomGom/ste-go-twitter)
![Contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg)

## Basic overview

Golang steganography software to hide data in Twitter using retweets.

## Dependencies

```bash
~$ go get github.com/Sirupsen/logrus
~$ go get github.com/dghubble/go-twitter/twitter
~$ go get github.com/dghubble/oauth1
~$ go get golang.org/x/crypto/ssh/terminal
```

## Compiling Binary

```bash
~$ go build main.go
```

## To send hidden message

```bash
./main send -message "Hello world!" -hashtags "music,food,travel"
```

Output:

```bash
Password: <you must type here your password to hide the message>
Logged as: SenderTwitterUser
Tweet 001 with ID 1381619419126845442 containing the target "program #music" successfully retweeted!
Tweet 002 with ID 1382102318339006467 containing the target "waves #music" successfully retweeted!
Tweet 003 with ID 1381676033737285635 containing the target "nationally" successfully retweeted!
Tweet 004 with ID 1380543128747139075 containing the target "psychology #food" successfully retweeted!
Tweet 005 with ID 1381928667958423554 containing the target "rick #music" successfully retweeted!
Tweet 006 with ID 1382229146974220290 containing the target "exclude" successfully retweeted!
```

or

```bash
~$ go run main.go send -message "Hello world" -hashtags "music,food,travel"
```

## To recieve message

```bash
~$ ./main recv -sender "SenderTwitterUser" -retweets "6"
```

or

```bash
~$ go run main.go recv -sender "SenderTwitterUser" -retweets "6"
```

Output:

```bash
Password: <you must type here the same password used to hide the message>
Logged as: YourTwitterUser
Secret message: "hello world!"
```

## Acknowledgements

* [PhD Daniel Lerch](https://daniellerch.me/), for the original idea in Python: https://github.com/daniellerch/stego-retweet
* [PhD Alfonso Muñoz](https://www.linkedin.com/in/alfonsomu%C3%B1oz/), for his patience and dedication to the dissemination of Cryptography and Steganography.
