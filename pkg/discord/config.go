package discord

type DiscordConfig interface {
	WebHookID() int
	WebHookToken() string
}
