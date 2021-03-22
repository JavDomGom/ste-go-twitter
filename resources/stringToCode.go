package resources

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/JavDomGom/ste-go-twitter/config"
)

func StringToCode(s string) int64 {
	for {
		if len(s)%config.MsgLenChunk != 0 {
			s += " "
		} else {
			break
		}
	}

	var binString string

	for _, c := range s {
		binString += fmt.Sprintf("%06b", strings.Index(config.CharSet, string(c)))
	}

	code, err := strconv.ParseInt(binString, 2, 32)
	if err != nil {
		fmt.Println(err)
	}

	return code
}
