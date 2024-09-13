package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

const (
	Width   = 800
	Height  = 600
	NumDots = 5000
)

func main() {
	// Membuat gambar baru dengan latar belakang putih
	img := image.NewRGBA(image.Rect(0, 0, Width, Height))

	// Menggambar titik-titik acak dengan warna tertentu
	for i := 0; i < NumDots; i++ {
		x := rand.Intn(Width)
		y := rand.Intn(Height)
		img.Set(x, y, color.RGBA{uint8(rand.Intn(256)), uint8(rand.Intn(256)), uint8(rand.Intn(256)), 255})
	}

	// Menyimpan gambar ke file
	file, err := os.Create("titik-titik.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}
