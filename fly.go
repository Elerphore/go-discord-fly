// https://discord.com/developers/docs/resources/webhook#create-webhook
// POST/channels/{channel.id}/webhooks

// https://discord.com/developers/docs/resources/webhook#get-channel-webhooks
// GET/channels/{channel.id}/webhooks

// https://discord.com/developers/docs/resources/webhook#get-guild-webhooks
// GET/guilds/{guild.id}/webhooks

// https://discord.com/developers/docs/resources/webhook#get-webhook
// GET/webhooks/{webhook.id}

// https://discord.com/developers/docs/resources/webhook#get-webhook-with-token
// GET/webhooks/{webhook.id}/{webhook.token}

// https://discord.com/developers/docs/resources/webhook#modify-webhook
// PATCH/webhooks/{webhook.id}

// https://discord.com/developers/docs/resources/webhook#modify-webhook-with-token
// PATCH/webhooks/{webhook.id}/{webhook.token}

// https://discord.com/developers/docs/resources/webhook#delete-webhook
// DELETE/webhooks/{webhook.id}

// https://discord.com/developers/docs/resources/webhook#delete-webhook-with-token
// DELETE/webhooks/{webhook.id}/{webhook.token}

// https://discord.com/developers/docs/resources/webhook#execute-webhook
// POST/webhooks/{webhook.id}/{webhook.token}

// https://discord.com/developers/docs/resources/webhook#get-webhook-message
// GET/webhooks/{webhook.id}/{webhook.token}/messages/{message.id}

// https://discord.com/developers/docs/resources/webhook#edit-webhook-message
// PATCH/webhooks/{webhook.id}/{webhook.token}/messages/{message.id}

// https://discord.com/developers/docs/resources/webhook#delete-webhook-message
// DELETE/webhooks/{webhook.id}/{webhook.token}/messages/{message.id}

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
		Content:    "Message for  the cuttiest girl ever..",
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

	const webhookToken = "https://discord.com/api/webhooks/947040904034779156/zSAT0yoEUsgK02tjNHAuaPKtsf_UEoXSr3B1Mhv7qINlSKCQqevCpqXVsBOuNMT30wYE"
	req, err := http.NewRequest("POST", webhookToken, bytes.NewBuffer(requestBody))
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		log.Fatal(err)
	}

	client.Do(req)
}
