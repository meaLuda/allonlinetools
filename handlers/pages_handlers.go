package handlers

import (
	"github.com/gofiber/fiber/v2"
	"allonlinetools/sessionstore"

)

func HandleGetHome(c *fiber.Ctx) error {
	// Get session from storage
	sess, err := sessionstore.Store.Get(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Server Error")
	}

	// Get user ID from session
	userID := sess.Get("user_id")
	if userID == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	// Render home page with user ID
	return c.Render("home/index", fiber.Map{
		"user_id": userID,
	})
}


func HandleTool_GetImages(c *fiber.Ctx) error {
	// Get session from storage
	sess, err := sessionstore.Store.Get(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Server Error")
	}

	// Get user ID from session
	userID := sess.Get("user_id")
	if userID == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}
	mp := fiber.Map{
		"success": true,
		"message": "I'm receiving success with inline success data",
		"user_id": userID,
	}
	return c.Render("home/pages/image_tool", mp)
}