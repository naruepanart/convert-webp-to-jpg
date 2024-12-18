package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"math"
	"os"
	"path/filepath"

	_ "golang.org/x/image/webp"
)

func main() {
	const dir = "./"
	fs, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Failed to read folder:", err)
		return
	}

	for _, f := range fs {
		if f.IsDir() || filepath.Ext(f.Name()) != ".webp" {
			continue
		}

		in := filepath.Join(dir, f.Name())
		out := in[:len(in)-len(filepath.Ext(f.Name()))] + ".jpg"

		if err := convert(in, out); err != nil {
			fmt.Printf("Error converting %s: %v\n", in, err)
			continue
		}

		fmt.Printf("Converted: %s -> %s\n", in, out)
	}
}

func convert(in, out string) error {
	f, err := os.Open(in)
	if err != nil {
		return fmt.Errorf("open file: %v", err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return fmt.Errorf("decode image: %v", err)
	}

	b := img.Bounds()
	w, h := b.Dx(), b.Dy()

	// Resize maintaining aspect ratio
	var resized image.Image
	if h > w {
		resized = resize(img, int(float64(w)*1920/float64(h)), 1920)
	} else {
		resized = resize(img, 1920, int(float64(h)*1920/float64(w)))
	}

	// Adjust and save
	adjusted := adjust(resized)
	file, err := os.Create(out)
	if err != nil {
		return fmt.Errorf("create output: %v", err)
	}
	defer file.Close()

	return jpeg.Encode(file, adjusted, &jpeg.Options{Quality: 100})
}

func resize(src image.Image, w, h int) image.Image {
	srcB := src.Bounds()
	srcW, srcH := srcB.Dx(), srcB.Dy()

	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			srcX := int(float64(x) * float64(srcW) / float64(w))
			srcY := int(float64(y) * float64(srcH) / float64(h))

			if srcX < 0 || srcX >= srcW || srcY < 0 || srcY >= srcH {
				continue
			}

			dst.Set(x, y, src.At(srcX, srcY))
		}
	}

	return dst
}

func adjust(src image.Image) image.Image {
	b := src.Bounds()
	dst := image.NewRGBA(b)
	draw.Draw(dst, b, src, b.Min, draw.Src)

	return contrast(dst, 1.2)
}

func contrast(img *image.RGBA, f float64) *image.RGBA {
	b := img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			c := img.At(x, y)
			r, g, b, a := c.RGBA()

			rf := math.Max(0, math.Min(1, ((float64(r>>8)/255.0-0.5)*f)+0.5))
			gf := math.Max(0, math.Min(1, ((float64(g>>8)/255.0-0.5)*f)+0.5))
			bf := math.Max(0, math.Min(1, ((float64(b>>8)/255.0-0.5)*f)+0.5))

			img.Set(x, y, color.RGBA{
				R: uint8(rf * 255),
				G: uint8(gf * 255),
				B: uint8(bf * 255),
				A: uint8(a >> 8),
			})
		}
	}
	return img
}
