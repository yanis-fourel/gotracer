package main

import (
	"image"
	"image/color"
	"math/rand"
)

type Render = struct {
	img    *image.RGBA
	update chan bool
}

type GoTracer struct {
	render Render
	// scene
	// camera
}

func NewGoTracer() *GoTracer {
	return &GoTracer{
		render: Render{
			img:    nil,
			update: make(chan bool, 1),
		},
	}
}

func (gt *GoTracer) Run() {
	for {
		for x := 0; x < 640; x++ {
			for y := 0; y < 480; y++ {
				gt.render.img.Set(x, y, color.RGBA{
					R: uint8(rand.Uint32() % 256),
					G: uint8(rand.Uint32() % 256),
					B: uint8(rand.Uint32() % 256),
					A: 255,
				})
			}
		}
		select {
		case gt.render.update <- true:
		default:
		}
	}
}
