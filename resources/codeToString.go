package resources

import (
	"fmt"
	"strconv"

	"github.com/JavDomGom/ste-go-twitter/config"
	"github.com/Sirupsen/logrus"
)

func CodeToString(log *logrus.Logger, code int) string {
	var bits_str = fmt.Sprintf("%012b", code)

	code1, err := strconv.ParseInt(bits_str[:6], 2, 32)
	if err != nil {
		log.Fatal(err)
	}
	code2, err := strconv.ParseInt(bits_str[6:], 2, 32)
	if err != nil {
		log.Fatal(err)
	}

	if code1 > int64(len(config.CharSet)) || code2 > int64(len(config.CharSet)) {
		fmt.Println("ERROR: Password incorrect, please try again.")
		log.Fatalf(
			"Password incorrect: Index can't be greater than CharSet length (%d)",
			len(config.CharSet),
		)
	}

	return string(config.CharSet[code1]) + string(config.CharSet[code2])
}
