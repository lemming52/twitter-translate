package translator

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"
)

type Client struct {
	lgr *log.Entry
}

func NewTranslator() *Client {
	return &Client{
		lgr: log.WithFields(log.Fields{
			"pkg": "translator",
		}),
	}
}

func (c *Client) Translate(ctx context.Context, text, language string) (string, error) {
	if text != "kiev" {
		return "", errors.New("can't handle")
	}
	return "Київ", nil
}
