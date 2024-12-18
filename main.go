package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
	_ "golang.org/x/image/webp"
)

func main() {
	const folder = "./"
	files, err := os.ReadDir(folder)
	if err != nil {
		fmt.Println("Failed to read folder:", err)
		return
	}

	for _, f := range files {
		if f.IsDir() || filepath.Ext(f.Name()) != ".webp" {
			continue
		}
		inFile := filepath.Join(folder, f.Name())
		outFile := inFile[:len(inFile)-len(filepath.Ext(f.Name()))] + ".jpg"
		if err := convertWebPToJPEGWithHDR(inFile, outFile); err != nil {
			fmt.Printf("Error converting %s: %v\n", inFile, err)
		} else {
			fmt.Printf("Converted: %s -> %s\n", inFile, outFile)
		}
	}
}

func convertWebPToJPEGWithHDR(inFile, outFile string) error {
	file, err := os.Open(inFile)
	if err != nil {
		return fmt.Errorf("open file: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("decode image: %v", err)
	}

	// Get image dimensions
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	// Resize based on orientation
	if height > width {
		// Portrait: Resize height to 1920 and maintain aspect ratio
		img = imaging.Resize(img, int(float64(width)*1920/float64(height)), 1920, imaging.Lanczos)
	} else {
		// Landscape: Resize width to 1920 and maintain aspect ratio
		img = imaging.Resize(img, 1920, int(float64(height)*1920/float64(width)), imaging.Lanczos)
	}

	// Apply other adjustments
	img = imaging.Sharpen(img, 1.5)
	img = imaging.AdjustBrightness(img, 5)
	img = imaging.AdjustContrast(img, 10)
	img = imaging.AdjustSaturation(img, 5)

	out, err := os.Create(outFile)
	if err != nil {
		return fmt.Errorf("create output file: %v", err)
	}
	defer out.Close()

	return jpeg.Encode(out, img, &jpeg.Options{Quality: 100})
}
