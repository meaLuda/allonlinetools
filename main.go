package main

import (
	"context"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"go.mongodb.org/mongo-driver/mongo"

	// ------- local packages
	"allonlinetools/db"
	"allonlinetools/handlers"
	"allonlinetools/middleware"
	"allonlinetools/sessionstore"

	// ------- fiber packages
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/django/v3"
	"github.com/joho/godotenv"
)

var (
	userController *handlers.UserController
	
)

func initApp() (error) {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	var userCollection *mongo.Collection = db.OpenCollection("user")
	ctx := context.Background()

	userController = handlers.NewUserController(userCollection, ctx)
	
	// Switch states according to the environment we are in.
	weAreIn := os.Getenv("LOCAL_DEVELOPMENT")
	log.Println(weAreIn)

	
	return nil
}

func createEngine() *django.Engine {
	engine := django.New("./views", ".html")
	engine.Reload(true)
	engine.AddFunc("css", func(name string) (res template.HTML) {
		filepath.Walk("static/assets", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Name() == name {
				res = template.HTML("<link rel=\"stylesheet\" href=\"/" + path + "\">")
			}
			return nil
		})
		return
	})

	return engine
}


// route test
func ToolsRoutes(app *fiber.App) {
	//// ------------------- Tools routes
	tools := app.Group("/tools")
	
	//// ---------- tools in the shed
	images := tools.Group("/images")
	// images.Get("",middleware.Protected(store),handlers.HandleTool_GetImages)
	images.Get("",handlers.HandleTool_GetImages)
	//// functionalities for images
	images.Post("/upload_image",handlers.Handler_UploadImage)
	images.Get("/rotate-image", handlers.Handler_rotateImage)
}



func main() {
	//// Initialize database and app
	// db.Init()
	err := initApp()
	if err != nil {
		log.Fatal(err)
	}
	app := fiber.New(fiber.Config{
		ErrorHandler:          handlers.ErrorHandler,
		DisableStartupMessage: false,
		PassLocalsToViews:     true,
		Views:                 createEngine(),
	})
	app.Use(favicon.New(favicon.ConfigDefault))
	app.Use(recover.New())
	app.Use(logger.New())
	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	
	// Use CSRF token middleware
    app.Use(middleware.CSRFTOKENMiddleware())
	// Apply CSRF protection middleware for form submissions
	app.Use(middleware.CSRFProtection())
	// Initialize session middleware
	

	

	app.Static("/static", "./static", fiber.Static{
		CacheDuration: 0,
	})
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
		c.Set("Pragma", "no-cache")
		c.Set("Expires", "0")
		c.Set("Surrogate-Control", "no-store")
		return c.Next()
	})


	//// pages
	app.Get("/",middleware.Protected(sessionstore.Store), handlers.HandleGetHome)

	// User Auth
	app.Get("/auth/signup",userController.Handler_RenderSignUpPage)
	app.Post("/auth/signup", userController.CreateUser)
	app.Get("/auth/login",userController.Handler_RenderLoginPage)
	app.Post("/auth/login", userController.LoginUser)

	// making code cleaner.
	ToolsRoutes(app)
	
	// app.Get("/resize-image", resizeImage)
	// app.Get("/flip-image", flipImage)
	// app.Get("/convert-jpeg", convertToJPEG)
	// app.Get("/convert-png", convertToPNG)
	// app.Get("/convert-gif", convertToGIF)
	// app.Get("/adjust-brightness", adjustBrightness)
	// app.Get("/adjust-contrast", adjustContrast)
	// app.Get("/adjust-saturation", adjustSaturation)



	err = app.Listen(":3000")
	if err != nil {
		log.Fatal("Error in running the server")
		return
	}
	log.Println("Server is running")
}



