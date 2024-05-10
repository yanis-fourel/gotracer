package main

type Ray struct {
	ori, dir Vec3
}

type Scene struct {
	backgroundColor RGB
	ambientLight    RGB
	dirLight        DirLight
	spheres         []Sphere
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
