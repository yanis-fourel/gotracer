package main

import (
	"image"
	"image/color"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func makeEditor() fyne.CanvasObject {
	return container.NewVScroll(container.NewVBox(
		canvas.NewText("first line of editor", color.Black),
		canvas.NewText("second line of editor", color.Black),
	))
}

func makeViewport() fyne.CanvasObject {
	img := image.NewRGBA(image.Rect(0, 0, 640, 480))
	res := canvas.NewImageFromImage(img)

	go func() {
		for {
			for x := 0; x < 640; x++ {
				for y := 0; y < 480; y++ {
					img.Set(x, y, color.RGBA{
						R: uint8(rand.Uint32() % 256),
						G: uint8(rand.Uint32() % 256),
						B: uint8(rand.Uint32() % 256),
						A: 255,
					})
				}
			}
			res.Refresh()
		}
	}()

	return res
}

func main() {
	a := app.New()
	w := a.NewWindow("Viewport")

	content := container.NewHSplit(
		makeEditor(),
		makeViewport(),
	)
	content.Offset = 0.3

	w.SetContent(content)
	w.ShowAndRun()
}
