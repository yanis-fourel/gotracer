package main

import (
	"fmt"
	"image"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
)

type Render = struct {
	img *image.RGBA
}

// Type to indicate there is a new frame that should be rendered
type NewFrameMsg struct{}

type Model struct {
	render Render
	scene  Scene
	cam    Camera
}

func initialModel() *Model {
	spheres := [...]Sphere{
		{
			center: Vec3{0, 0, 10},
			radius: 1,
		},
	}

	return &Model{
		render: Render{
			img: image.NewRGBA(image.Rect(0, 0, 40, 64)),
		},
		scene: Scene{
			spheres: spheres[:],
		},
		cam: Camera{
			origin: Vec3{0, 0, 0},
			up:     Vec3Up,
			fw:     Vec3Forward,
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
	return m, nil
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

func (m *Model) Run(p *tea.Program) {
	for {
		// TODO: Run that only when the scene or camera changed
		renderScene(m.render.img, m.scene, m.cam)
		p.Send(NewFrameMsg{})
		return
	}
}

func main() {
	m := initialModel()
	p := tea.NewProgram(m)
	go m.Run(p)
	_, err := p.Run()
	if err != nil {
		log.Fatalf("Alas, there's been an error: %v", err)
	}
}
