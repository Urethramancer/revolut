package main

// WebhookCmd adds a webhook for callback triggering.
type WebhookCmd struct {
	Args struct {
		ID string `required:"true" positional-arg-name:"URL" description:"URL of webhook to add."`
	} `positional-args:"true"`
}

func (cmd *WebhookCmd) Execute(args []string) error {
	return nil
}
