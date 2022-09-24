package web

import (
	"encoding/json"
	"net/http"
)

type LinkHandler struct {
	translate TranslationService
}

func NewLinkHandler(t TranslationService) *LinkHandler {
	return &LinkHandler{
		translate: t,
	}
}

func (l *LinkHandler) GetLinks(r *http.Request) (int, interface{}, error) {
	req := &LinkRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	res, err := l.translate.GetLinks(r.Context(), toTranslations(req))
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, toLinkResponse(res), nil
}
