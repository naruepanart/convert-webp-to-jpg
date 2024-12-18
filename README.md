# WebP to JPEG Converter

This Go program converts all `.webp` images in the current directory to `.jpg` format. It resizes the images to a width or height of 1920 pixels (depending on the aspect ratio) while maintaining the original aspect ratio. The images are also adjusted with a slight contrast enhancement before being saved as JPEGs.

## Features

- Converts all `.webp` files in the directory to `.jpg`.
- Maintains aspect ratio when resizing.
- Adjusts contrast of the image before saving.
- Saves output with 100% quality.

## Requirements

- Go 1.18 or higher.
- The `golang.org/x/image/webp` package for WebP support.

## Installation

1. Install Go (if not installed yet) from [the official Go website](https://golang.org/dl/).
2. Download or clone this repository.
3. Install the necessary Go package:

```bash
go get golang.org/x/image/webp
```

## Usage

1. Place the `.webp` files you want to convert in the same directory as the Go program.
2. Run the program:

```bash
go run main.go
```

3. The program will convert all `.webp` files in the current directory to `.jpg` format, applying resizing and contrast adjustment.
   
4. The output files will be saved in the same directory with the `.jpg` extension.

## Example

If the current directory contains a file `image.webp`, after running the program, it will be converted and saved as `image.jpg`.

## Code Explanation

- **resize**: Resizes the image while maintaining its aspect ratio.
- **adjust**: Applies contrast enhancement to the image.
- **convert**: Main function that opens a `.webp` file, resizes and adjusts it, and saves it as `.jpg`.