package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func send(text string, bot string, chat_id string) {
	request_url := "https://api.telegram.org/" + bot + "/sendMessage"

	client := &http.Client{}

	values := map[string]string{"text": text, "chat_id": chat_id}
	json_paramaters, _ := json.Marshal(values)

	req, _ := http.NewRequest("POST", request_url, bytes.NewBuffer(json_paramaters))
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.Status)
		defer res.Body.Close()
	}
}
