package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	client := &http.Client{}

	requestBody, err := json.Marshal(map[string]string{
		"content": "Again.",
	})

	if err != nil {
		log.Fatal("error")
	}

	const webhookToken = ""
	req, err := http.NewRequest("POST", webhookToken, bytes.NewBuffer(requestBody))
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		log.Fatal("error")
	}

	client.Do(req)
}
