package cache

import (
	"context"
	"fmt"

	"github.com/lemming52/twitter-translate/translate"
)

// Cache is a basic in memory cache
// With the interfaces, swapping this for a full implementation requires no changes to the service code.
type Cache struct {
	store map[string]*translationDTO
}

// NewCache instantiate the in memory map
func NewCache() *Cache {
	return &Cache{
		store: map[string]*translationDTO{},
	}
}

func (c *Cache) AddTranslation(ctx context.Context, t *translate.Translation) error {
	key := translationRecordKey(t.Text, t.Language)
	c.store[key] = toTranslationDTO(t, key)
	return nil
}

func (c *Cache) GetTranslation(ctx context.Context, text, language string) (*translate.Translation, error) {
	key := translationRecordKey(text, language)
	t, ok := c.store[key]
	if !ok {
		return nil, fmt.Errorf("borked")
	}
	return fromTranslationDTO(t), nil
}

func translationRecordKey(text, language string) string {
	return fmt.Sprintf("translation: %s %s", text, language)
}
