package notify_test

import (
	"github.com/juunini/simple-go-line-notify/notify"
	"io/ioutil"
	"testing"
)

func ExampleSendImage() {
	accessToken := "29jWoO****p70eK3AKA********ooHfusvDP6***ZmR"
	message := "hello, world!"
	imageURL := "https://d.line-scdn.net/n/line_lp/img/ogimage.png"

	if err := notify.SendImage(accessToken, message, imageURL); err != nil {
		panic(err)
	}
}

func TestSendImage(t *testing.T) {
	accessToken, err := ioutil.ReadFile("test_token.txt")
	if err != nil {
		return
	}
	imageURL := "https://d.line-scdn.net/n/line_lp/img/ogimage.png"

	if err := notify.SendImage(string(accessToken), "a", imageURL); err != nil {
		t.Error(err)
		t.Fail()
	}
}
