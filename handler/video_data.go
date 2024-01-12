package handler

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

func (h *Handler) GetVideoData(ctx *fiber.Ctx) error {
	ids := strings.Split(ctx.Query("videoIds"), " ")
	videos := h.e.GetVideos(ids)
	r := make([]*videoDataResponse, len(videos))

	for i, v := range videos {
		if v != nil {
			r[i] = newVideoDataResponse(v)
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(r)
}
