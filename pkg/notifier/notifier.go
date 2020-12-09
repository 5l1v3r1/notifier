package notifier

import (
	"github.com/drsigned/gos"
	"github.com/valyala/fasthttp"
)

// Notify handles the notification engine
type Notify struct {
	options *Options
	client  *fasthttp.Client
	slack   *Slack
}

// New notify instance
func New(options *Options) (*Notify, error) {
	client := &fasthttp.Client{}
	return &Notify{
		options: options,
		client:  client,
		slack: &Slack{
			client:     client,
			webHookURL: options.YAMLConfig.Platforms.Slack.WebHookURL,
		},
	}, nil
}

// SendNotification to registered webhooks
func (n *Notify) SendNotification(message string) error {
	message = gos.StripANSI(message)

	if n.options.YAMLConfig.Platforms.Slack.Use {
		err := n.slack.send(message)
		if err != nil {
			return err
		}
	}

	return nil
}
