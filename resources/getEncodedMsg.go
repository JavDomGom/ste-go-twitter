package resources

import (
	"unicode/utf8"

	"github.com/Sirupsen/logrus"
)

func GetEncodedMsg(log *logrus.Logger, msg string, length int) []int64 {
	log.Debug("Getting encoded message.")

	var l, r int

	chunks := []string{}
	seqList := []int64{}

	for l, r = 0, length; r < len(msg); l, r = r, r+length {
		for !utf8.RuneStart(msg[r]) {
			r--
		}
		chunks = append(chunks, msg[l:r])
	}
	chunks = append(chunks, msg[l:])

	for _, chunk := range chunks {
		seqList = append(seqList, StringToCode(chunk))
	}

	log.Debugf("Encoded message is: %v", seqList)

	return seqList
}
