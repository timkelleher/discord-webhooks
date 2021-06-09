package main

import (
	"fmt"
	"strconv"

	"github.com/timkelleher.com/discord-webhooks/pkg/discord"
	"github.com/timkelleher.com/discord-webhooks/pkg/integrations/cli"
)

func main() {
	var msg, id, token string

	fmt.Println("Enter your message: ")
	fmt.Scanln(&msg)
	fmt.Println("Enter ID: ")
	fmt.Scanln(&id)
	fmt.Println("Enter token: ")
	fmt.Scanln(&token)

	id2, _ := strconv.Atoi(id)
	payload := cli.NewPayload(msg, id2, token)
	payload.Succeeded = false
	discord.NewWebHook(payload, payload)
}
