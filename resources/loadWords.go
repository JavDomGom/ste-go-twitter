package resources

import (
	"bufio"
	"os"

	"github.com/Sirupsen/logrus"
)

// LoadWords returns a list with all words in file splitted by line.
func LoadWords(log *logrus.Logger, path string) ([]string, error) {
	log.Debug("Loading words from file.")

	var lines []string

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
