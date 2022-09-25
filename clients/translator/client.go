package translator

import (
	"context"
	"errors"

	"cloud.google.com/go/translate"
	log "github.com/sirupsen/logrus"
	"golang.org/x/text/language"
)

type Client struct {
	client *translate.Client
	lgr    *log.Entry
}

func NewTranslator() *Client {
	client, err := translate.NewClient(context.Background())
	if err != nil {
		log.WithError(err).Fatal("couldn't start the translator")
	}
	return &Client{
		lgr: log.WithFields(log.Fields{
			"pkg": "translator",
		}),
		client: client,
	}
}

func (c *Client) Translate(ctx context.Context, text, target string) (string, error) {
	targetLanguage, err := language.Parse(target)
	if err != nil {
		return "", err
	}
	resp, err := c.client.Translate(ctx, []string{text}, targetLanguage, nil)
	if err != nil {
		return "", err
	}
	if len(resp) == 0 {
		return "", errors.New("borked")
	}
	log.Info(resp)
	return resp[0].Text, nil
}
