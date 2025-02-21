package notify_test

import (
	"os"
	"testing"

	"github.com/frontware/simple-go-line-notify/notify"
)

func ExampleSendText() {
	accessToken := "29jWoO****p70eK3AKA********ooHfusvDP6***ZmR"
	message := "hello, world!"

	if err := notify.SendText(accessToken, message); err != nil {
		panic(err)
	}
}

func TestSendTextInvalidAccessToken(t *testing.T) {
	accessToken := "29jWoO****p70eK3AKA********ooHfusvDP6***ZmR"
	message := "hello, world!"

	err := notify.SendText(accessToken, message)
	if !(err != nil && err.Error() == "401: Invalid access token") {
		t.Log(err)
		t.Fail()
	}
}

func TestSendTextEmptyMessage(t *testing.T) {
	accessToken, err := os.ReadFile("test_token.txt")
	if err != nil {
		return
	}
	message := ""

	if err := notify.SendText(string(accessToken), message); !(err != nil && err.Error() == "400: message: must not be empty") {
		t.Log(err)
		t.Fail()
	}
}

func TestSendTextSuccess(t *testing.T) {
	accessToken, err := os.ReadFile("test_token.txt")
	if err != nil {
		return
	}
	message := "test"

	if err := notify.SendText(string(accessToken), message); err != nil {
		t.Log(err)
		t.Fail()
	}
}
