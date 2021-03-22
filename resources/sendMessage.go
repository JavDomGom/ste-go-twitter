package resources

import "fmt"

func SendMessage(encodedMsg []int64, hashtags []string) {
	fmt.Println("hashtags:", hashtags)

	for _, code := range encodedMsg {
		for _, hashtag := range hashtags {
			fmt.Printf("hashtags: %q, code: %v\n", hashtag, code)
		}
	}
}
