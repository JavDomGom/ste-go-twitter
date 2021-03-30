package resources

import (
	"fmt"
	"strconv"

	"github.com/JavDomGom/ste-go-twitter/config"
)

func CodeToString(code int) string {
	var bits_str = fmt.Sprintf("%012b", code)

	char1, err := strconv.ParseInt(bits_str[:6], 2, 32)
	if err != nil {
		fmt.Println(err)
	}
	char2, err := strconv.ParseInt(bits_str[6:], 2, 32)
	if err != nil {
		fmt.Println(err)
	}

	return string(config.CharSet[char1]) + string(config.CharSet[char2])
}
