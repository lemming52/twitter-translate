package web

import (
	"context"
	"github.com/lemming52/twitter-translate/translate"
)

//go:generate mockgen -source=interfaces.go -destination=interfaces_mock.go -package=web

type TranslationService interface {
	GetLinks(ctx context.Context, req []*translate.Translation) ([]string, error)
}
