package handler

import (
	"YTStreamGoApi/models"
	"YTStreamGoApi/utils"
	"errors"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

func (h *Handler) SignUp(ctx *fiber.Ctx) error {
	var req *signUpRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.NewError(err))
	}

	if err := h.validator.Validate(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.NewError(err))
	}

	if req.Password != req.PasswordConfirm {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.PasswordsNotMatch())
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.NewError(err))
	}

	u := new(models.User)
	u.Name = req.Name
	u.Password = string(hashedPassword)
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	if err := h.userStore.Create(u); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}

	token, err := utils.GenerateJWT(u.ID, h.config)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.NewError(err))
	}

	response := newUserResponse(u)
	response.Token = token

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (h *Handler) SignIn(ctx *fiber.Ctx) error {
	var req *signInRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.NewError(err))
	}

	if err := h.validator.Validate(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.NewError(err))
	}

	u, err := h.userStore.GetByName(req.Name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusForbidden).JSON(utils.AccessForbidden())
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.NewError(err))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(utils.AccessForbidden())
	}

	token, err := utils.GenerateJWT(u.ID, h.config)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.NewError(err))
	}

	response := newUserResponse(u)
	response.Token = token

	return ctx.Status(fiber.StatusOK).JSON(response)
}
