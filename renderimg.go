package main

import (
	"bytes"
	"fmt"

	"github.com/muesli/termenv"
)

// Minimum length of a cell to be able to write any truecolor
const CellLength = 20

type RenderImg struct {
	width, height int
	buff          []byte
}

func (img *RenderImg) String() string {
	return string(img.buff)
}

func NewRenderImg(width, height int, initcol RGB) *RenderImg {
	res := RenderImg{
		width,
		height,
		nil,
	}
	line := bytes.Repeat([]byte(padded(colorToStr(initcol))), width+1)
	copy(line[len(line)-CellLength:], newLine())
	res.buff = bytes.Repeat(line, height)
	return &res
}

func (img *RenderImg) offsetAt(x, y int) int {
	px_count := x + y*img.width
	nl_count := y

	return CellLength * (px_count + nl_count)
}

func (img *RenderImg) Set(x, y int, color RGB) {
	img.SetCellStr(x, y, colorToStr(color))
}

func (img *RenderImg) SetCellStr(x, y int, str string) {
	offset := img.offsetAt(x, y)
	copy(img.buff[offset:], str)
}

// / Returns the string needed to print a cell of given color
func colorToStr(color RGB) string {
	return padded(fmt.Sprintf(
		"%s%sm ",
		termenv.CSI,
		termenv.TrueColor.FromColor(color).Sequence(true),
	))
}

func newLine() string {
	return padded(fmt.Sprintf(
		"%s%sm\n",
		termenv.CSI,
		termenv.ResetSeq,
	))
}

func padded(str string) string {
	if len(str) > CellLength {
		panic(fmt.Sprintln("Too long: ", str))
	}

	prefix := bytes.Repeat([]byte{termenv.ESC}, CellLength-len(str))
	return string(prefix) + str
}
