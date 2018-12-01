package revolut

// WebhookRequest only holds the URL to add.
type WebhookRequest struct {
	// URL must be secure.
	URL string `json:"url"`
}

// AddWebhook adds URLs to post events to when transactions are created or updated.
func (c *Client) AddWebhook(url string) error {
	hook := WebhookRequest{
		URL: url,
	}
	contents, code, err := c.PostJSON(epWebhook, hook)
	if err != nil {
		return err
	}

	if code != 204 {
		return jsonError(contents)
	}

	return nil
}
