package main

type Ray struct {
	ori, dir Vec3
}

type Scene struct {
	spheres []Sphere
}

type Camera struct {
	origin Vec3
	up     Vec3
	fw     Vec3
}

type Sphere struct {
	center Vec3
	radius float32
}
