package resources

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

// AskPassword prompts user for an input that is not echo-ed on terminal.
func AskPassword() (string, error) {
	fmt.Printf("Password: ")

	raw, err := terminal.MakeRaw(0)
	if err != nil {
		return "", err
	}
	defer terminal.Restore(0, raw)

	var (
		prompt string
		answer string
	)

	term := terminal.NewTerminal(os.Stdin, prompt)
	for {
		char, err := term.ReadPassword(prompt)
		if err != nil {
			return "", err
		}
		answer += char

		if char == "" || char == answer {
			return answer, nil
		}
	}
}
