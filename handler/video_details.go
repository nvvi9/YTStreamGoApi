package handler

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

func (h *Handler) GetVideoDetails(ctx *fiber.Ctx) error {
	ids := strings.Split(ctx.Query("videoIds"), " ")
	videos := h.e.GetVideos(ids)

	return ctx.Status(fiber.StatusOK).JSON(videos)
}
