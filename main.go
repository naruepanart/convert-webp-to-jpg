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
	// Folder path to convert files
	const folder = "./"

	// Read all files in the folder
	files, err := os.ReadDir(folder)
	if err != nil {
		fmt.Printf("Failed to read folder: %v\n", err)
		return
	}

	// Process each WebP file
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
	// Open WebP file
	file, err := os.Open(inFile)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Decode image
	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("failed to decode image: %v", err)
	}

	// Apply image adjustments
	img = imaging.Sharpen(img, 1.5)
	img = imaging.AdjustBrightness(img, 10)
	img = imaging.AdjustContrast(img, 20)
	img = imaging.AdjustSaturation(img, 15)

	// Create output file
	out, err := os.Create(outFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer out.Close()

	// Encode as JPEG
	if err := jpeg.Encode(out, img, &jpeg.Options{Quality: 100}); err != nil {
		return fmt.Errorf("failed to encode JPEG: %v", err)
	}

	return nil
}
