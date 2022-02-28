package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	var comps [1]ComponentField

	comps[0] = ComponentField{
		Type:     2,
		Label:    "Click me!",
		Style:    1,
		CustomId: "click_one",
	}

	componentOne :=
		Component{
			Type:       1,
			Components: comps,
		}

	var compss [1]Component

	compss[0] = componentOne

	msg := WebHookMessage{
		Content:    "Lol, sorry.",
		Tts:        false,
		Username:   "Some random name",
		AvatarUrl:  "https://i.pinimg.com/originals/cd/bf/5f/cdbf5f93123170b215d1d81980d8a1ca.jpg",
		Components: compss,
	}

	client := &http.Client{}

	requestBody, err := json.Marshal(msg)

	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", webhookToken, bytes.NewBuffer(requestBody))
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		log.Fatal(err)
	}

	client.Do(req)
}
