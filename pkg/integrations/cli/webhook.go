package cli

func NewPayload(title, msg string, id int, token string) *CLIPayload {
	return &CLIPayload{
		title:        title,
		content:      msg,
		PayloadID:    id,
		PayloadToken: token,
	}
}

type CLIPayload struct {
	title        string
	content      string
	PayloadID    int
	PayloadToken string
}

func (cp CLIPayload) Integration() string {
	return "CLI"
}

func (cp CLIPayload) Title() string {
	return cp.title
}

func (cp CLIPayload) Content() string {
	return cp.content
}

func (cp CLIPayload) WebHookID() int {
	return cp.PayloadID
}

func (cp CLIPayload) WebHookToken() string {
	return cp.PayloadToken
}
