package resources

import (
	"bufio"
	"os"

	"github.com/Sirupsen/logrus"
)

type Log struct {
	Logger *logrus.Logger
}

// LoadWords returns a list with all words in file splitted by line.
func (f Log) LoadWords(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	f.Logger.Debug("List of words loaded successfull!")
	return lines, scanner.Err()
}
