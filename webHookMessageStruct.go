package main

type Component struct {
	Type       int               `json:"type"`
	Components [1]ComponentField `json:"components"`
}

type ComponentField struct {
	Type     int    `json:"type"`
	Label    string `json:"label"`
	Style    int    `json:"style"`
	CustomId string `json:"custom_id"`
}

type WebHookMessage struct {
	Content    string       `json:"content"`
	Tts        bool         `json:"tts"`
	Username   string       `json:"username"`
	AvatarUrl  string       `json:"avatar_url"`
	Components [1]Component `json:"components"`
}
