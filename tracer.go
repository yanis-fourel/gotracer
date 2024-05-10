package main

import (
	"image/color"
)

// Returns the perceived color by a ray shot in the scene
func trace(r Ray, scene Scene) color.RGBA {
	impact := raycast(r, scene)

	if impact == nil {
		return color.RGBA{0, 127, 127, 255}
	}
	return color.RGBA{
		uint8((impact.n.x + 1) * 100),
		uint8((impact.n.y + 1) * 100),
		uint8((impact.n.z + 1) * 100),
		255,
	}
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

func renderScene(r *Render, scene Scene, cam Camera) {
	img := r.img
	height := img.Bounds().Dy()
	width := img.Bounds().Dx()
	right := cam.up.Cross(cam.fw)

	inv_aspect_ratio := float64(height) / float64(width)

	topleft := cam.origin.
		Add(cam.fw).
		Add(cam.up.Scaled(inv_aspect_ratio * 0.5)).
		Add(right.Scaled(-0.5))

	dx := right.Scaled(1 / float64(width))
	dy := cam.up.Scaled(-inv_aspect_ratio / float64(height))

	doneChan := make(chan bool, 1000)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			go func() {
				target := topleft.
					Add(dx.Scaled(float64(x) + 0.5)).
					Add(dy.Scaled(float64(y) + 0.5))
				ray := Ray{
					ori: cam.origin,
					dir: target.Sub(cam.origin).Normalized(),
				}
				color := trace(ray, scene)
				img.Set(x, y, color)
				doneChan <- true
			}()
		}
	}

	totalCount := width * height
	doneCount := 0
	for <-doneChan {
		doneCount++
		r.progress = float32(doneCount) / float32(totalCount)
	}
}
