package gosaic

import (
	"image"
	_ "image/jpeg"
	"image/png"
	"math"
	"os"
)

// Scale scales the src image to the given rectangle using Nearest Neighbor
func Scale(srcImg image.Image, r image.Rectangle) image.Image {
	dstImg := image.NewRGBA(r)

	sw := srcImg.Bounds().Dx()
	sh := srcImg.Bounds().Dy()

	dw := dstImg.Bounds().Dx()
	dh := dstImg.Bounds().Dy()

	xAspect := float64(sw) / float64(dw)
	yAspect := float64(sh) / float64(dh)

	for y := 0; y < dh; y++ {
		for x := 0; x < dw; x++ {
			srcX := int(math.Floor(float64(x) * xAspect))
			srcY := int(math.Floor(float64(y) * yAspect))
			pix := srcImg.At(srcX, srcY)
			dstImg.Set(x, y, pix)
		}
	}

	return dstImg
}

func ScaleToFile(srcPath, dstPath string, r image.Rectangle) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	srcImg, _, err := image.Decode(srcFile)
	if err != nil {
		return err
	}
	dstImg := Scale(srcImg, r)

	return SavePng(dstImg, dstPath)
}

func SavePng(img image.Image, outfile string) error {
	toFile, err := os.Create(outfile)
	if err != nil {
		return err
	}
	defer toFile.Close()

	return png.Encode(toFile, img)
}
