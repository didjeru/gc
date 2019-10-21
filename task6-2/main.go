package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	white := color.RGBA{255, 255, 255, 255}
	red := color.RGBA{255, 0, 0, 255}

	// Прямоугольник
	img := image.NewRGBA(image.Rect(0, 0, 200, 200))
	draw.Draw(img, img.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)

	// Линии
	for x := 0; x < 200; x++ {
		img.Set(x, x, red)
		img.Set(x, 200-x, red)
		img.Set(100, x, red)
		img.Set(x, 100, red)
	}

	file, err := os.Create("rectangle.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}

	defer file.Close()

	png.Encode(file, img)
}
