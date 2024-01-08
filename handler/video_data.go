package handler

import (
	"YTStreamGoApi/utils"

	"github.com/gofiber/fiber/v2"
	ytstream "github.com/nvvi9/YTStreamGo"
)

func (h *Handler) GetVideoData(ctx *fiber.Ctx) error {
	videoId := ctx.Params("videoId")

	videoData, err := ytstream.ExtractVideoData(videoId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.NewError(err))
	}

	return ctx.Status(fiber.StatusOK).JSON(newVideoDataResponse(videoData))
}
