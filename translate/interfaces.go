package translate

import "context"

type TranslatorClient interface {
	Translate(ctx context.Context, text, language string) (string, error)
}

type Cache interface {
	AddTranslation(ctx context.Context, p *Translation) error
	GetTranslation(ctx context.Context, text, language string) (*Translation, error)
}
