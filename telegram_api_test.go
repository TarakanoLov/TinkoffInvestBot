package main

import (
	"testing"
)

func TestHelloName(t *testing.T) {
	t.Setenv("chat_id", "some_chat_id")
	t.Setenv("telegram_token", "some_telegram_token")
	my_send("some")
	// todo
}
