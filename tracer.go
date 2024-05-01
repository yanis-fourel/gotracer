package main

import (
	"image"
	"image/color"
)

func trace(r Ray, scene Scene) color.RGBA {
	return color.RGBA{0, 127, 127, 255}
}

// Runs raytracing algorithm to render the scene into the image
func renderScene(img *image.RGBA, scene Scene, cam Camera) {
	right := cam.up.Cross(cam.fw)

	aspect_ratio := float32(img.Rect.Dx()) / float32(img.Rect.Dy())

	topleft := cam.origin.
		Add(cam.fw).
		Add(cam.up.Scaled(aspect_ratio * 0.5)).
		Add(right.Scaled(-0.5))

	dx := right.Scaled(1 / float32(img.Rect.Dx()))
	dy := cam.up.Scaled(-aspect_ratio / float32(img.Rect.Dy()))

	for x := 0; x < img.Rect.Dx(); x++ {
		for y := 0; y < img.Rect.Dy(); y++ {
			go func() {
				target := topleft.
					Add(dx.Scaled(float32(x))).
					Add(dy.Scaled(float32(y)))
				ray := Ray{
					ori: cam.origin,
					dir: target.Sub(cam.origin).Normalized(),
				}
				color := trace(ray, scene)
				img.Set(x, y, color)
			}()
		}
	}
}
