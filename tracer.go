package main

import (
	"image"
	"image/color"
)

func trace(r Ray, scene Scene) color.RGBA {
	impact := raycast(r, scene)

	if impact == nil {
		return color.RGBA{0, 0, 0, 0}
	}
	return color.RGBA{255, 255, 255, 255}
}

func raycast(r Ray, scene Scene) *Impact {
	var res *Impact

	for i := range scene.spheres {
		r := RaycastSphere(r, scene.spheres[i])
		if r != nil && (res == nil || r.dist < res.dist) {
			res = r
		}
	}

	return res
}

// Runs raytracing algorithm to render the scene into the image
func renderScene(img *image.RGBA, scene Scene, cam Camera) {
	right := cam.up.Cross(cam.fw)

	aspect_ratio := float64(img.Rect.Dx()) / float64(img.Rect.Dy())

	topleft := cam.origin.
		Add(cam.fw).
		Add(cam.up.Scaled(aspect_ratio * 0.5)).
		Add(right.Scaled(-0.5))

	dx := right.Scaled(1 / float64(img.Rect.Dx()))
	dy := cam.up.Scaled(-aspect_ratio / float64(img.Rect.Dy()))

	for x := 0; x < img.Rect.Dx(); x++ {
		for y := 0; y < img.Rect.Dy(); y++ {
			go func() {
				target := topleft.
					Add(dx.Scaled(float64(x))).
					Add(dy.Scaled(float64(y)))
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
