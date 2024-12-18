# WebP to JPEG Converter with HDR Adjustment

This Go program converts `.webp` images in a specified folder into `.jpg` format while applying various image enhancements. It performs the following steps:
- Resizes images based on orientation (portrait or landscape).
- Sharpens the image.
- Adjusts contrast and saturation.
- Converts the image to JPEG with HDR-like adjustments.

## Features
- Automatically processes all `.webp` files in the current directory.
- Resizes images based on their orientation, ensuring that portrait images are resized to a maximum height of 1920 pixels and landscape images are resized to a maximum width of 1920 pixels.
- Applies image adjustments such as sharpening, contrast, and saturation to enhance the image.
- Saves the output as high-quality `.jpg` files with 100% quality.

## Requirements
- Go 1.18+.
- Dependencies:
  - [`disintegration/imaging`](https://github.com/disintegration/imaging): for image manipulation tasks like resizing, sharpening, and adjusting contrast and saturation.
  - `golang.org/x/image/webp`: to support decoding `.webp` images.

## Installation

1. Clone the repository or create a Go module and include the necessary dependencies.
2. Install Go and make sure the following dependencies are installed:

```bash
go get -u github.com/disintegration/imaging
go get -u golang.org/x/image/webp
```

## Usage

### Running the Program

1. Place your `.webp` images in the same directory as the program.
2. Run the program:

```bash
go run main.go
```

The program will read all `.webp` files in the directory, convert them to `.jpg`, apply the HDR adjustments, and save them as new `.jpg` files in the same folder.

### Folder Structure

The program expects the following folder structure:

```
.
├── main.go
├── image1.webp
├── image2.webp
└── ...
```

For every `.webp` file, it will generate a corresponding `.jpg` file in the same folder.

## How It Works

1. **Reading and Looping through Files**: The program reads the contents of the current directory and loops through all files. It skips directories and non-WebP files.
2. **Converting the WebP to JPEG**: 
   - It opens and decodes the `.webp` image file.
   - It resizes the image to 1920px in the larger dimension while maintaining the aspect ratio.
   - It sharpens the image and applies contrast and saturation adjustments.
3. **Saving the JPEG Image**: The image is saved as a high-quality `.jpg` with HDR-like enhancements.

## Code Details

- The `main()` function reads the folder and processes `.webp` files.
- The `convertWebPToJPEGWithHDR()` function performs the image processing, including resizing, sharpening, contrast adjustment, and saturation adjustment.
- It then saves the processed image as `.jpg` using JPEG encoding with a quality of 100.

## Example Output

```
Converted: ./image1.webp -> ./image1.jpg
Converted: ./image2.webp -> ./image2.jpg
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.