package handlers

import (
	"log"
	"path/filepath"
	"github.com/gofiber/fiber/v2"
	tool "allonlinetools/tools"

)

// Get image uploaded by users
func Handler_UploadImage(c *fiber.Ctx) error {
	// Get the uploaded file from the form
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("Error in uploading image:", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}
	
	// Define the destination path
	destination := filepath.Join("static", "uploads", file.Filename)
	log.Println("Destination:", destination)

	// Save the file to the destination path
	if err := c.SaveFile(file, destination); err != nil {
		log.Println("Error saving uploaded file:", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	context := fiber.Map{
		"filename": file.Filename,
	}	
	return c.Render("partials/tools/image_conv", context)
}

func Handler_rotateImage(c *fiber.Ctx) error {
	src := c.Query("image_name")
	// Define the destination path
	srcPath := filepath.Join("static", "uploads", src)
	log.Println("Image Src:", srcPath)

	// Load the image
	img, err := tool.LoadImage(srcPath)
	if err != nil {
		log.Println("Error loading image:", err)
		context := fiber.Map{
			"error": err,
		}
		return c.Render("partials/error_badge", context)
	}

	// rotate image
	rotated_image := tool.RotateImage(img,45.3)
	// save new image
	err = tool.SaveImage(rotated_image,filepath.Join("static", "uploads"))
	if err != nil {
		log.Println("Error saving image:", err)
		context := fiber.Map{
			"error": "Unable to save image ",
		}
		return c.Render("partials/error_badge", context)
	}
	context := fiber.Map{
		"filename": src,
	}
	return c.Render("partials/tools/updated_image", context)
}
