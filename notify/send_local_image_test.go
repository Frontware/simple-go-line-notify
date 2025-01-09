package notify_test

import (
	"os"
	"testing"

	"github.com/frontware/simple-go-line-notify/notify"
)

func ExampleSendLocalImage() {
	accessToken := "29jWoO****p70eK3AKA********ooHfusvDP6***ZmR"
	message := "hello, world!"
	imagePath := "./sample.png"

	if err := notify.SendLocalImage(accessToken, message, imagePath); err != nil {
		panic(err)
	}
}

func TestSendLocalImage(t *testing.T) {
	accessToken, err := os.ReadFile("test_token.txt")
	if err != nil {
		return
	}
	message := "local image test"
	imagePath := "./sample.jpg"

	if err := notify.SendLocalImage(string(accessToken), message, imagePath); err != nil {
		t.Error(err)
		t.Fail()
	}
}
