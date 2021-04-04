package resources

import (
	"crypto/sha256"
	"encoding/hex"
	"math/big"
	"math/rand"

	"github.com/Sirupsen/logrus"
)

func ShuffleWords(log *logrus.Logger, pwd string, words []string) {
	pwdSHA256 := sha256.Sum256([]byte(pwd))
	pwdSHA256String := hex.EncodeToString(pwdSHA256[:])

	log.Debugf("Password SHA256 string: %v", pwdSHA256String)
	log.Debug(
		"Cutting password SHA256 string in chunks of 4 bytes and seed-shuffle list of words.",
	)

	c := 1
	for i := 0; i < 64; i += 8 {
		pwdSHA256BigInt := new(big.Int)
		pwdSHA256BigInt.SetString(pwdSHA256String[i:i+8], 16)
		pwdSHA256Int64 := pwdSHA256BigInt.Int64()
		log.Debugf(
			"Seeding with chunk %d: %v => %v as %T.",
			c,
			pwdSHA256String[i:i+8],
			pwdSHA256Int64,
			pwdSHA256Int64,
		)
		rand.Seed(pwdSHA256Int64)

		log.Debug("Shuffling list of words.")
		rand.Shuffle(len(words), func(i, j int) {
			words[i], words[j] = words[j], words[i]
		})
		c++
	}
}
