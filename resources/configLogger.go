package resources

import (
	"log"
	"os"

	"github.com/JavDomGom/ste-go-twitter/config"
	"github.com/Sirupsen/logrus"
)

func ConfigLogger() *logrus.Logger {
	var log = logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{DisableHTMLEscape: true})
	log.SetLevel(logrus.DebugLevel)

	return log
}

func GetLogFile() *os.File {
	if _, err := os.Stat(config.LogPath); os.IsNotExist(err) {
		os.MkdirAll(config.LogPath, 0744)
	}
	file, err := os.OpenFile(
		config.LogPath+"/ste-go-twitter.log",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
