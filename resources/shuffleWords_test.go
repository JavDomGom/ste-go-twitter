package resources

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestShuffleWords(t *testing.T) {
	log := ConfigLogger()
	log.SetOutput(ioutil.Discard)

	type test struct {
		pwd      string
		words_ok []string
	}

	tests := []test{
		{"abcdefgh", []string{"word3", "word1", "word0", "word4", "word2"}},
		{"abcd1234", []string{"word4", "word0", "word1", "word2", "word3"}},
		{"1234abcd", []string{"word3", "word0", "word2", "word4", "word1"}},
	}

	for _, x := range tests {
		// Load list of words.
		words, err := LoadWords(log, "../db/words_test.txt")
		if err != nil {
			log.Fatal(err)
		}

		pwd := x.pwd
		ShuffleWords(log, pwd, words)

		if !reflect.DeepEqual(words, x.words_ok) {
			t.Errorf("Password: %q Expected: %v Got: %v", pwd, x.words_ok, words)
		}
	}
}
