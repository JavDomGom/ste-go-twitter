package resources

import (
	"math/big"
	"math/rand"

	"github.com/Sirupsen/logrus"
)

func GetShuffledWords(log *logrus.Logger, pwdSHA256String string, words []string) []string {
	log.Debug(
		"Cutting password SHA256 string in chunks of 4 bytes and seed-shuffle list of words.",
	)
	log.Debugf("SHA256 string: %v", pwdSHA256String)

	c := 1
	for i := 0; i < 64; i += 8 {
		pwdSHA256BigInt := new(big.Int)
		pwdSHA256BigInt.SetString(pwdSHA256String[i:i+8], 16)
		pwdSHA256Int64 := pwdSHA256BigInt.Int64()
		log.Debugf(
			"Chunk %d: %v => %v as %T.",
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

	return words
}
