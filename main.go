package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"path/filepath"

	"golang.org/x/image/draw"
	"golang.org/x/image/webp"
)

func main() {
	const dir = "./"
	fs, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Failed to read folder:", err)
		return
	}

	for _, f := range fs {
		// Only process .webp files (skip directories and non-webp files)
		if f.IsDir() || filepath.Ext(f.Name()) != ".webp" {
			continue
		}

		in := filepath.Join(dir, f.Name())
		out := in[:len(in)-len(filepath.Ext(f.Name()))] + ".jpg"

		if err := convert(in, out); err != nil {
			fmt.Printf("Error converting %s: %v\n", in, err)
		} else {
			fmt.Printf("Converted: %s -> %s\n", in, out)
		}
	}
}

func convert(in, out string) error {
	// Open input WebP file
	f, err := os.Open(in)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %v", in, err)
	}
	defer f.Close()

	// Decode WebP image
	img, err := webp.Decode(f)
	if err != nil {
		return fmt.Errorf("failed to decode image %s: %v", in, err)
	}

	// Resize image while maintaining aspect ratio
	bounds := img.Bounds()
	var resized image.Image
	if bounds.Dy() > bounds.Dx() {
		// Landscape (height > width): Resize based on width
		resized = resize(img, int(float64(bounds.Dx())*1920/float64(bounds.Dy())), 1920)
	} else {
		// Portrait (width > height): Resize based on height
		resized = resize(img, 1920, int(float64(bounds.Dy())*1920/float64(bounds.Dx())))
	}

	// Adjust contrast of the resized image
	adjusted := adjust(resized)

	// Create output file
	file, err := os.Create(out)
	if err != nil {
		return fmt.Errorf("failed to create output file %s: %v", out, err)
	}
	defer file.Close()

	// Encode the adjusted image as JPEG
	return jpeg.Encode(file, adjusted, &jpeg.Options{Quality: 100})
}

func resize(src image.Image, w, h int) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	// Use x/image/draw for efficient resizing
	draw.ApproxBiLinear.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Over, nil)
	return dst
}

func adjust(src image.Image) image.Image {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)

	// Apply contrast adjustment
	applyContrast(dst, src, 1.2)
	return dst
}

func applyContrast(dst *image.RGBA, src image.Image, factor float64) {
	bounds := dst.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := src.At(x, y)
			r, g, b, a := c.RGBA()

			// Apply contrast adjustment
			rf := clamp((float64(r>>8)/255.0-0.5)*factor + 0.5)
			gf := clamp((float64(g>>8)/255.0-0.5)*factor + 0.5)
			bf := clamp((float64(b>>8)/255.0-0.5)*factor + 0.5)

			// Set the adjusted color value
			dst.Set(x, y, color.RGBA{
				R: uint8(rf * 255),
				G: uint8(gf * 255),
				B: uint8(bf * 255),
				A: uint8(a >> 8),
			})
		}
	}
}

// Helper function to clamp values between 0 and 1
func clamp(val float64) float64 {
	if val < 0 {
		return 0
	} else if val > 1 {
		return 1
	}
	return val
}
