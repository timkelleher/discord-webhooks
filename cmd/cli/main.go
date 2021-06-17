package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/timkelleher.com/discord-webhooks/pkg/discord"
	"github.com/timkelleher.com/discord-webhooks/pkg/integrations/cli"
)

func main() {
	var msg, id, token string
	var idInt int

	msg = os.Getenv("discord_message")
	idInt, _ = strconv.Atoi(os.Getenv("discord_webhook_id"))
	token = os.Getenv("discord_webhook_token")

	if msg == "" {
		fmt.Println("Enter your message: ")
		fmt.Scanln(&msg)
		fmt.Println("Enter ID: ")
		fmt.Scanln(&id)
		fmt.Println("Enter token: ")
		fmt.Scanln(&token)
		idInt, _ = strconv.Atoi(id)
	}

	payload := cli.NewPayload(msg, idInt, token)
	payload.Succeeded = false
	discord.NewWebHook(payload, payload)
}
