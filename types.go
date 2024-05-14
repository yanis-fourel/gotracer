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

type Material struct {
	// Surface color
	color RGB

	// 0 to 1
	reflection float64
}

type Sphere struct {
	center Vec3
	radius float64
	mat    Material
}

type Plane struct {
	point Vec3
	norm  Vec3
	mat   Material
}
