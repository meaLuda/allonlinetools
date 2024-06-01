package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/gofiber/fiber/v2"

    "allonlinetools/sessionstore" // global session store
)

const(
	csrfTokenHeader = "X-CSRF-Token"
	csrfTokenLocal = "csrf_token"
)

// generateToken creates a random CSRF token
func generateToken() (string, error) {
    token := make([]byte, 32)
    _, err := rand.Read(token)
    if err != nil {
        return "", err
    }
    return base64.StdEncoding.EncodeToString(token), nil
}

// CSRFProtection is the middleware function
func CSRFProtection() fiber.Handler {
    return func(c *fiber.Ctx) error {
        // Skip if the request method is safe (GET, HEAD, OPTIONS, TRACE)
        if c.Method() == fiber.MethodGet || c.Method() == fiber.MethodHead || c.Method() == fiber.MethodOptions || c.Method() == fiber.MethodTrace {
            return c.Next()
        }


        // Get CSRF token from header or form
        csrfToken := c.Get(csrfTokenHeader)
        if csrfToken == "" {
            csrfToken = c.FormValue("csrf_token")
            fmt.Println("----------- Middleware csrfToken -------------------")
            fmt.Println(csrfToken)
        }
        
        // Get session from storage
        sess, err := sessionstore.Store.Get(c)
        if err != nil {
            panic(err)
        }
        // fmt.Println("----------- Middleware SESSION KEYS -------------------")
        // Get all Keys
        // keys := sess.Keys()
        // fmt.Println(keys)
        local_token := sess.Get(csrfTokenLocal)
        fmt.Println(local_token)
        if csrfToken == local_token {
            return c.Next()
        }

        return c.Status(fiber.StatusForbidden).SendString("Invalid CSRF Token")
    }
}


// CSRFTOKENMiddleware sets a CSRF token in the locals for safe methods
func CSRFTOKENMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        // Only set token for safe methods
        if c.Method() == fiber.MethodGet || c.Method() == fiber.MethodHead {
            token, err := generateToken()
            if err != nil {
                return c.Status(fiber.StatusInternalServerError).SendString("Error generating CSRF Token")
            }
            // Store token in locals
            c.Locals(csrfTokenLocal, token)
        }
        return c.Next()
    }
}