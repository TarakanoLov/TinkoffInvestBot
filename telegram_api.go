package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func my_send(text string) {
	request_url := "https://api.telegram.org/bot" + os.Getenv("telegram_token") + "/sendMessage"

	client := &http.Client{}

	values := map[string]string{"text": text, "chat_id": os.Getenv("chat_id")}
	json_paramaters, _ := json.Marshal(values)

	req, _ := http.NewRequest("POST", request_url, bytes.NewBuffer(json_paramaters))
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	} else {
		body, _ := io.ReadAll(res.Body)
		fmt.Println(res.Status, string(body), 9, os.Getenv("chat_id"), 9, os.Getenv("telegram_token"))
		defer res.Body.Close()
	}
}
