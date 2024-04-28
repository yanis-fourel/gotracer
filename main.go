package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math/rand"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
)

type Render = struct {
	img    *image.RGBA
	update chan bool
}

type Model struct {
	render Render
	// scene
	// camera
}

func initialModel() *Model {
	return &Model{
		render: Render{
			img:    image.NewRGBA(image.Rect(0, 0, 40, 64)),
			update: make(chan bool, 1),
		},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, tea.ClearScreen
}

func (m Model) View() string {
	s := ""
	for x := range m.render.img.Bounds().Dx() {
		for y := range m.render.img.Bounds().Dy() {
			s += fmt.Sprintf(
				"%s%sm ",
				termenv.CSI,
				termenv.ANSI256.FromColor(m.render.img.At(x, y)).Sequence(true),
			)
			s += " "
		}
		s += "\n"
	}
	s += fmt.Sprintf(
		"%s%sm ",
		termenv.CSI,
		termenv.ResetSeq,
	)
	return s
}

func (m *Model) Run() {
	for {
		for x := 0; x < 640; x++ {
			for y := 0; y < 480; y++ {
				m.render.img.Set(x, y, color.RGBA{
					R: uint8(rand.Uint32() % 256),
					G: uint8(rand.Uint32() % 256),
					B: uint8(rand.Uint32() % 256),
					A: 255,
				})
			}
		}
		select {
		case m.render.update <- true:
		default:
		}
	}
}

func main() {
	m := initialModel()
	go m.Run()
	p := tea.NewProgram(m)
	_, err := p.Run()
	if err != nil {
		log.Fatalf("Alas, there's been an error: %v", err)
	}
}
