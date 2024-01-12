package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) Register(r *fiber.App) {
	v1 := r.Group("/api")

	details := v1.Group("/details")
	details.Get("/", h.GetVideoDetails)

	data := v1.Group("/data")
	data.Get("/", h.GetVideoData)
}
