package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	
)


func Protected(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get session
		sess, err := store.Get(c)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Session error")
		}

		// Check if user is logged in
		if sess.Get("user_id") == nil {
			return c.Redirect("/auth/login")
		}

		// Continue to the next handler if logged in
		return c.Next()
	}
}
