package resources

import (
	"unicode/utf8"
)

func GetEncodedMsg(msg string, length int) []int64 {
	chunks := []string{}
	seqList := []int64{}

	var l, r int
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
	return seqList
}
