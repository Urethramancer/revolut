package main

import "github.com/Urethramancer/slog"

// WebhookCmd adds a webhook for callback triggering.
type WebhookCmd struct {
	Args struct {
		URL string `required:"true" positional-arg-name:"URL" description:"URL of webhook to add."`
	} `positional-args:"true"`
}

func (cmd *WebhookCmd) Execute(args []string) error {
	c, err := newClient()
	if err != nil {
		return err
	}

	slog.Msg("Webhook added.")
	return c.AddWebhook(cmd.Args.URL)
}
