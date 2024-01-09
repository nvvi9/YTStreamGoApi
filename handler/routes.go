package handler

import (
	"YTStreamGoApi/middleware"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Register(r *fiber.App) {
	v1 := r.Group("/api")
	v1.Post("/signup", h.SignUp)
	v1.Post("/signin", h.SignIn)

	details := v1.Group("/details", middleware.Protected(h.config))
	details.Get("/:videoId", h.GetVideoDetails)

	data := v1.Group("/data", middleware.Protected(h.config))
	data.Get("/:videoId", h.GetVideoData)
}
