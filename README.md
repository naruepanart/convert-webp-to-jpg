# WebP to JPEG Converter with HDR Adjustments

This Go program converts all `.webp` images in a specified folder to `.jpg` format while applying High Dynamic Range (HDR) adjustments. It uses the `disintegration/imaging` package to enhance the images with sharpening, brightness, contrast, and saturation adjustments.

## Features

- Converts `.webp` files to `.jpg` format.
- Applies HDR-like adjustments:
  - Sharpens the image.
  - Adjusts brightness, contrast, and saturation.
- Processes all `.webp` files in the current folder.
- Efficient and minimal error handling.

## Requirements

- Go 1.16+ (for `os.ReadDir`)
- Go modules enabled
- Dependencies:
  - `github.com/disintegration/imaging` for image manipulation.
  - `golang.org/x/image/webp` for WebP image decoding.

## Installation

1. Clone or download this repository.
2. Install Go and set up your Go environment if you haven't already.

```bash
go mod tidy
```

3. Install dependencies:

```bash
go get github.com/disintegration/imaging
go get golang.org/x/image/webp
```

## Usage

1. Place the Go source file in your desired project folder.
2. Place all the `.webp` files you want to convert into the same folder (or adjust the `folder` variable in the code to point to your folder of choice).
3. Run the program:

```bash
go run main.go
```

The program will:

- Read all files in the folder.
- Convert each `.webp` file to `.jpg` with the same name, applying the following adjustments:
  - Sharpening: 1.5x strength
  - Brightness: +2
  - Contrast: +20
  - Saturation: +15
- The converted files will be saved as `.jpg` in the same folder.

### Example

If the folder contains a file named `image.webp`, the output file will be `image.jpg`.

## Error Handling

- If the program fails to read the folder or process any files, it will output a corresponding error message.

## License

This program is open-source software released under the MIT License.