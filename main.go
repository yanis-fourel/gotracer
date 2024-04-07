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

func rgbGradient(x, y, w, h int) color.Color {
	g := int(float32(x) / float32(w) * float32(255))
	b := int(float32(y) / float32(h) * float32(255))

	return color.NRGBA{uint8(255 - b), uint8(g), uint8(b), 0xff}
}

func getRandomImage() *canvas.Image {
	img := image.NewRGBA(image.Rect(0, 0, 640, 480))

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

	return canvas.NewImageFromImage(img)
}

func loop(w fyne.Window) {
	for {
		w.SetContent(getRandomImage())
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("Viewport")

	editor := container.NewVBox(
		canvas.NewText("first line of editor", color.Black),
		canvas.NewText("second line of editor", color.Black),
	)
	viewport := canvas.NewCircle(color.RGBA{R: 255, G: 126, B: 126, A: 255})

	content := container.NewHSplit(
		container.NewVScroll(editor),
		viewport,
	)
	content.Offset = 0.3

	w.SetContent(content)
	w.ShowAndRun()
}
