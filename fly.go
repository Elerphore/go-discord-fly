package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	msg := WebHookMessage{
		Content: "There's some structs I've used here.",
		Tts:     true,
	}

	client := &http.Client{}

	requestBody, err := json.Marshal(msg)

	log.Println(requestBody)

	if err != nil {
		log.Fatal("error")
	}

	req, err := http.NewRequest("POST", webhookToken, bytes.NewBuffer(requestBody))
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		log.Fatal("error")
	}

	client.Do(req)
}
