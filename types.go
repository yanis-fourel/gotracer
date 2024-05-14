package main

type Ray struct {
	ori, dir Vec3
}

type Scene struct {
	backgroundColor RGB
	ambientLight    RGB
	dirLight        DirLight
	spheres         []Sphere
	planes          []Plane
}

type Camera struct {
	origin Vec3
	up     Vec3
	fw     Vec3
}

type DirLight struct {
	dir Vec3
	col RGB
}

type Sphere struct {
	center Vec3
	radius float64
	color  RGB
}

type Plane struct {
	point Vec3
	norm  Vec3
	color RGB
}
