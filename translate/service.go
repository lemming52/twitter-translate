package translate

import (
	"context"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	translator TranslatorClient
	cache      Cache
	lgr        *log.Entry
}

func NewService(translator TranslatorClient, cache Cache) *Service {
	return &Service{
		translator: translator,
		cache:      cache,
		lgr: log.WithFields(log.Fields{
			"pkg": "pokemon",
		}),
	}
}

func (s *Service) GetLinks(ctx context.Context, translations []*Translation) ([]string, error) {
	links := make([]string, len(translations))
	for i, t := range translations {
		err := s.getTranslation(ctx, t)
		if err != nil {
			return nil, err
		}
		links[i] = t.getTwitterLink()
	}
	return links, nil
}

func (s *Service) getTranslation(ctx context.Context, t *Translation) error {
	cached, err := s.cache.GetTranslation(ctx, t.Text, t.Language)
	if err == nil {
		t.Translation = cached.Translation
		return nil
	} else {
		s.lgr.WithError(err).Info("failed to retrieve from cache")
	}
	translated, err := s.translator.Translate(ctx, t.Text, t.Language)
	if err != nil {
		return err
	}
	t.Translation = translated
	err = s.cache.AddTranslation(ctx, t)
	if err != nil {
		s.lgr.WithError(err).Warn("failed to update cache")
	}
	return nil
}
