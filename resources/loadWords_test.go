package resources

import (
	"io/ioutil"
	"testing"
)

func TestLoadWords(t *testing.T) {
	log := ConfigLogger()
	log.SetOutput(ioutil.Discard)

	words, err := LoadWords(log, "../db/words_test.txt")
	if err != nil {
		t.Errorf("Expected: \"File opened successfully\", Got: %q", err)
	}

	lenWords := len(words)
	if lenWords != 5 {
		t.Error("Expected:", 5, "Got:", lenWords)
	}
}
