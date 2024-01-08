package handler

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Register(r *fiber.App) {
	v1 := r.Group("/api")
	v1.Post("/signup", h.SignUp)
	v1.Post("/signin", h.SignIn)

	jwtMiddleware := jwtware.New(
		jwtware.Config{
			SigningKey: jwtware.SigningKey{Key: []byte(h.config.JwtSecret)},
		})

	details := v1.Group("/details", jwtMiddleware)
	details.Get("/:videoId", h.GetVideoDetails)

	data := v1.Group("/data", jwtMiddleware)
	data.Get("/:videoId", h.GetVideoData)
}
