package tools_test

import (
    "image"
    "testing"
	"math"
    "github.com/disintegration/imaging"
    "github.com/stretchr/testify/assert"

    "allonlinetools/tools"
)

func TestImageOperations(t *testing.T) {
    // Load the source image
    srcImage, err := imaging.Open("test_image.png")
    if err != nil {
        // Handle error
        t.Fatalf("Failed to open test image: %v", err)
    }

    // Test ResizeImage
    dstImage := tools.ResizeImage(srcImage, 128, 128)
    assert.Equal(t, 128, dstImage.Bounds().Dx())
    assert.Equal(t, 128, dstImage.Bounds().Dy())

    // Test RotateImage
    dstImage = tools.RotateImage(dstImage, 45.0*math.Pi/180.0)
    assert.NotEqual(t, srcImage, dstImage)

    // Test BlurImage
    dstImage = tools.BlurImage(dstImage, 5.0)
    assert.NotEqual(t, srcImage, dstImage)

    // Test CropImage
    rect := image.Rect(10, 10, 118, 118)
    dstImage = tools.CropImage(dstImage, rect)
    assert.Equal(t, 108, dstImage.Bounds().Dx())
    assert.Equal(t, 108, dstImage.Bounds().Dy())
}