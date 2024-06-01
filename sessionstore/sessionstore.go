// sessionstore/sessionstore.go
package sessionstore

import (
    "github.com/gofiber/fiber/v2/middleware/session"
)

// Store is a global session store
var Store = session.New()
