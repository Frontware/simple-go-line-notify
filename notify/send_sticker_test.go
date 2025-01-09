package notify_test

import (
	"os"
	"testing"

	"github.com/frontware/simple-go-line-notify/notify"
)

func ExampleSendSticker() {
	accessToken := "29jWoO****p70eK3AKA********ooHfusvDP6***ZmR"
	message := "hello, world!"
	stickerPackageId := 1
	stickerId := 113

	if err := notify.SendSticker(accessToken, message, stickerPackageId, stickerId); err != nil {
		panic(err)
	}
}

func TestSendStickerSuccess(t *testing.T) {
	accessToken, err := os.ReadFile("test_token.txt")
	if err != nil {
		return
	}
	message := "test"

	if err := notify.SendSticker(string(accessToken), message, 1, 113); err != nil {
		t.Log(err)
		t.Fail()
	}
}
