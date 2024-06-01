package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sujit-baniya/flash"
)

type Error struct {
	err error
}

func (e Error) Error() string {
	return e.err.Error()
}

func AppError(err error) Error {
	return Error{
		err: err,
	}
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	// logger.Log("error", err.Error())
	if _, ok := err.(Error); ok {
		return flash.WithData(c, fiber.Map{"appError": err.Error()}).RedirectBack("/")
	}

	return render500(c)
}

func Render404(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).Render("errors/404", fiber.Map{})
}

func render500(c *fiber.Ctx) error {
	return c.Status(500).Render("errors/500", fiber.Map{})
}
