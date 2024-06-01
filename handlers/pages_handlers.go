package handlers

import (
	"github.com/gofiber/fiber/v2"
	"allonlinetools/sessionstore"
	"fmt"
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

	// Perform any checks or operations with the user ID
	fmt.Println("User ID:", userID)

	// Render home page with user ID
	return c.Render("home/index", fiber.Map{
		"user_id": userID,
	})
}


func HandleTool_GetImages(c *fiber.Ctx) error {
	mp := fiber.Map{
		"success": true,
		"message": "I'm receiving success with inline success data",
	}
	return c.Render("home/pages/image_tool", mp)
}