package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/timkelleher.com/discord-webhooks/pkg/discord"
	"github.com/timkelleher.com/discord-webhooks/pkg/integrations/cli"
)

func getInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Input was: %q\n", line)
		return line
	}
	return ""
}

func main() {
	var title, msg, id, token string
	var idInt int

	title = os.Getenv("discord_title")
	msg = os.Getenv("discord_message")
	idInt, _ = strconv.Atoi(os.Getenv("discord_webhook_id"))
	token = os.Getenv("discord_webhook_token")

	if title == "" {
		fmt.Println("Enter your title: ")
		title = getInput()
	}
	if msg == "" {
		fmt.Println("Enter your message: ")
		msg = getInput()
	}
	if idInt == 0 {
		fmt.Println("Enter ID: ")
		fmt.Scanln(&id)
		idInt, _ = strconv.Atoi(id)
	}
	if token == "" {
		fmt.Println("Enter token: ")
		fmt.Scanln(&token)
	}

	payload := cli.NewPayload(title, msg, idInt, token)
	discord.NewGenericWebHook(payload, payload)
}
