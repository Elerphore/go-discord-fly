package main

type WebHookMessage struct {
	Content string `json:"content"`
	Tts     bool   `json:"tts"`
}
