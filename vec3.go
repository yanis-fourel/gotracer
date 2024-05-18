package main

import "math"

type Vec3 struct {
	x, y, z float64
}

var Vec3Up = Vec3{0, 1, 0}
var Vec3Right = Vec3{1, 0, 0}
var Vec3Forward = Vec3{0, 0, 1}

func (a Vec3) Add(b Vec3) Vec3 {
	a.x += b.x
	a.y += b.y
	a.z += b.z
	return a
}

func (a Vec3) Sub(b Vec3) Vec3 {
	a.x -= b.x
	a.y -= b.y
	a.z -= b.z
	return a
}

func (a Vec3) Dot(b Vec3) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

func (a Vec3) Cross(b Vec3) Vec3 {
	// Stolen from https://github.com/ungerik/go3d/blob/55ced4bcb3347d37e613f054daa9211ea1a06a3b/vec3/vec3.go#L252
	return Vec3{
		x: a.y*b.z - a.z*b.y,
		y: a.z*b.x - a.x*b.z,
		z: a.x*b.y - a.y*b.x,
	}
}

func (v Vec3) Scaled(f float64) Vec3 {
	v.x *= f
	v.y *= f
	v.z *= f
	return v
}

func (v Vec3) Normalized() Vec3 {
	return v.Scaled(1 / v.Length())
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSqr())
}

func (v Vec3) LengthSqr() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

func (v Vec3) RotatedAroundX(rad float64) Vec3 {
	return Vec3{
		v.x,
		v.y*math.Cos(rad) - v.z*math.Sin(rad),
		v.y*math.Sin(rad) + v.z*math.Cos(rad),
	}
}

func (v Vec3) RotatedAroundY(rad float64) Vec3 {
	return Vec3{
		v.z*math.Sin(rad) + v.x*math.Cos(rad),
		v.y,
		v.z*math.Cos(rad) - v.x*math.Sin(rad),
	}
}
