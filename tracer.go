package main

// Returns the perceived color by a ray shot in the scene
func trace(r Ray, scene *Scene) RGB {
	impact := raycast(r, scene)

	if impact == nil {
		return scene.backgroundColor
	}

	// return debug_distanceToColor(impact.dist).MixSub(debug_normToColor(impact.n))
	// return debug_distanceToColor(impact.dist)
	// return debug_normToColor(impact.n)

	lightcol := getLight(impact, scene)

	color := impact.mat.color.MixSub(lightcol)

	if impact.mat.reflection > 0.01 {
		reflectRay := Ray{
			ori: impact.p,
			dir: r.dir.Sub(impact.n.Scaled(2 * r.dir.Dot(impact.n))),
		}
		reflectCol := trace(reflectRay, scene)

		color = color.Scaled(1 - impact.mat.reflection)
		color = color.MixAdd(reflectCol.Scaled(impact.mat.reflection))
	}

	return color
}

// The closer, the whiter. The further, the darker
func debug_distanceToColor(d float64) RGB {
	f := d * 32
	if f > 200 {
		return RGB{200, 200, 200}
	}
	u := uint8(f)
	return RGB{u, u, u}
}

func debug_normToColor(n Vec3) RGB {
	return RGB{
		uint8(127 + n.x*127),
		uint8(127 + n.y*127),
		uint8(127 + n.z*127),
	}
}

// Returns the color of the light on the impact
func getLight(impact *Impact, scene *Scene) RGB {
	dirLightCol := getDirLight(impact, scene)

	res := scene.ambientLight.MixAdd(dirLightCol)
	return res
}

// Returns the color of the global directional light on the impact
func getDirLight(impact *Impact, scene *Scene) RGB {
	if impact.n.Dot(scene.dirLight.dir) >= 0 {
		return RGB{}
	}
	lightray := Ray{
		ori: impact.p,
		dir: scene.dirLight.dir.Scaled(-1),
	}
	if raycast(lightray, scene) != nil {
		return RGB{}
	}
	return scene.dirLight.col.Scaled(impact.n.Dot(lightray.dir))
}

func raycast(r Ray, scene *Scene) *Impact {
	var res *Impact

	for _, sphere := range scene.spheres {
		r := RaycastSphere(r, sphere)
		if r != nil && (res == nil || r.dist < res.dist) {
			res = r
		}
	}
	for _, plane := range scene.planes {
		r := RaycastPlane(r, plane)
		if r != nil && (res == nil || r.dist < res.dist) {
			res = r
		}
	}

	return res
}

func renderScene(r *Render, scene *Scene, cam Camera) {
	img := r.img
	height := img.height
	width := img.width
	right := cam.up.Cross(cam.fw)

	// We assume each cell is twice as high as it is wide
	termCellAspectRatio := float64(2)
	invAspectRatio := float64(height) / float64(width) * termCellAspectRatio

	topleft := cam.origin.
		Add(cam.fw).
		Add(cam.up.Scaled(invAspectRatio * 0.5)).
		Add(right.Scaled(-0.5))

	dx := right.Scaled(1 / float64(width))
	dy := cam.up.Scaled(-invAspectRatio / float64(height))

	doneChan := make(chan bool, 1000)
	for x := 0; x < width; x++ {
		go func() {
			for y := 0; y < height; y++ {
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
			}
		}()
	}

	totalCount := width * height
	doneCount := 0
	for <-doneChan {
		doneCount++
		r.progress = float32(doneCount) / float32(totalCount)
	}
}
