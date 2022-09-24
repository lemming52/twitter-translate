package web

import "github.com/lemming52/twitter-translate/translate"

type LinkRequest struct {
	Text      string
	Languages []string
}

type LinkResponse struct {
	Links []string
}

func toTranslations(req *LinkRequest) []*translate.Translation {
	translations := make([]*translate.Translation, len(req.Languages))
	for i, l := range req.Languages {
		translations[i] = toTranslation(req, l)
	}
	return translations
}

func toTranslation(req *LinkRequest, language string) *translate.Translation {
	return &translate.Translation{
		Text:     req.Text,
		Language: language,
	}
}

func toLinkResponse(links []string) *LinkResponse {
	return &LinkResponse{
		Links: links,
	}
}
