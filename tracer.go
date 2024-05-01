package main

import (
	"image"
	"image/color"
	"math/rand"
)

// Runs raytracing algorithm to render the scene into the image
// TODO: take scene and camera as param
func trace(img *image.RGBA) {
	for x := 0; x < img.Rect.Dx(); x++ {
		for y := 0; y < img.Rect.Dy(); y++ {
			img.Set(x, y, color.RGBA{
				R: uint8(rand.Uint32() % 256),
				G: uint8(rand.Uint32() % 256),
				B: uint8(rand.Uint32() % 256),
				A: 255,
			})
		}
	}
}
