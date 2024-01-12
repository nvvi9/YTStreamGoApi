package handler

import "YTStreamGoApi/extractor"

type Handler struct {
	e *extractor.Extractor
}

func NewHandler(e *extractor.Extractor) *Handler {
	return &Handler{e: e}
}
