package cache

import (
	"github.com/lemming52/twitter-translate/translate"
)

type translationDTO struct {
	Key          string
	Text         string
	Translation  string
	Language     string
	Alternatives []string
}

// toTranslationDTO converts to the stored record format
func toTranslationDTO(t *translate.Translation, key string) *translationDTO {
	return &translationDTO{
		Key:          key,
		Text:         t.Text,
		Translation:  t.Translation,
		Language:     t.Language,
		Alternatives: t.Alternatives,
	}
}

// fromTranslationDTO converts from the stored record format
func fromTranslationDTO(t *translationDTO) *translate.Translation {
	return &translate.Translation{
		Text:         t.Text,
		Translation:  t.Translation,
		Language:     t.Language,
		Alternatives: t.Alternatives,
	}
}
