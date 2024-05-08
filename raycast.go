package main

import "math"

type Impact struct {
	dist float64
	p    Vec3
	n    Vec3
}

func RaycastSphere(r Ray, s Sphere) *Impact {
	delta := r.ori.Sub(s.center)
	toto := math.Pow(r.dir.Dot(delta), 2) -
		delta.LengthSqr() + s.radius*s.radius

	if toto < 0 {
		return nil
	}

	p1_d := -(r.dir.Dot(delta)) - math.Sqrt(toto)
	p2_d := -(r.dir.Dot(delta)) + math.Sqrt(toto)

	dist := min(p1_d, p2_d)

	if dist < 0 {
		return nil
	}

	p := r.ori.Add(r.dir.Scaled(dist))
	n := p.Sub(s.center).Normalized()
	return &Impact{dist, p, n}
}