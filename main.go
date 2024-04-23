package main

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func makeEditor(gt *GoTracer) fyne.CanvasObject {
	return container.NewVScroll(container.NewVBox(
		canvas.NewText("first line of editor", color.Black),
		canvas.NewText("second line of editor", color.Black),
	))
}

func makeViewport(gt *GoTracer) fyne.CanvasObject {
	gt.render.img = image.NewRGBA(image.Rect(0, 0, 640, 480))
	render := canvas.NewImageFromImage(gt.render.img)
	render.FillMode = canvas.ImageFillOriginal

	go func() {
		for {
			<-gt.render.update
			render.Refresh()
		}
	}()

	return render
}

func main() {
	gt := NewGoTracer()

	a := app.New()
	w := a.NewWindow("Viewport")

	content := container.NewHSplit(
		makeEditor(gt),
		makeViewport(gt),
	)
	content.Offset = 0.3

	w.SetContent(content)

	go gt.Run()

	w.ShowAndRun()
}
