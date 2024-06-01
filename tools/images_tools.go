package tools

import (
	"image"
	"image/color"
	"log"
	"os"
    "strings"
	"github.com/disintegration/imaging"
    _ "image/png"
    _ "image/jpeg"
    "path/filepath"
)

// LoadImage loads an image from the specified file path
func LoadImage(src string) (image.Image, error) {
	log.Println("--------Image Loading ----------")
	log.Printf("Image Path: %v",src)
	file, err := os.Open(src)
	if err != nil {
		log.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Println("Error decoding image:", err)
		return nil, err
	}
	return img, nil
}

// SaveImage saves the image to the specified file path
func SaveImage(img image.Image, path string) error {
	ext := strings.ToLower(filepath.Ext(path))
	var err error
	switch ext {
	case ".jpg", ".jpeg":
		err = imaging.Save(img, path, imaging.JPEGQuality(95))
	case ".png":
		err = imaging.Save(img, path)
	default:
		err = imaging.Save(img, path) // Default to PNG if no extension or unsupported format
	}

	if err != nil {
		log.Println("Error saving image:", err)
	}
	return err
}


// Resize image to desired width and height
func ResizeImage(srcImage image.Image, width, height int) image.Image {
    // Create a new image with the desired size
    dstImage := imaging.Resize(srcImage, width, height, imaging.Lanczos)
    return dstImage
}

// Rotate image by a given angle
func RotateImage(srcImage image.Image, angle float64) image.Image {
    // Rotate the image
    dstImage := imaging.Rotate(srcImage, angle, color.Transparent)

    return dstImage
}

// Blur image using a Gaussian kernel
func BlurImage(srcImage image.Image, sigma float64) image.Image {
    // Apply Gaussian blur
    dstImage := imaging.Blur(srcImage, sigma)

    return dstImage
}

// Crop image to a given rectangle
func CropImage(srcImage image.Image, rect image.Rectangle) image.Image {
    // Crop the image
    dstImage := imaging.Crop(srcImage, rect)

    return dstImage
}