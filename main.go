package main

import (
	"fmt"
	"image"
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
	"golang.org/x/term"
	// "github.com/muesli/termenv"
)

type Render = struct {
	img *image.RGBA
	// From 0 to 1
	progress float32
}

// Type to indicate there is a new frame that should be rendered
type NewFrameMsg struct{}

type Model struct {
	// Data of the render of the current image
	render *Render
	scene  Scene
	cam    Camera
}

func initialModel() *Model {
	return &Model{
		render: nil,
		scene: Scene{
			backgroundColor: RGB{0, 127, 127},
			light: DirLight{
				dir: Vec3{-0.5, -1, 1}.Normalized(),
				col: RGB{255, 31, 31},
			},
			spheres: []Sphere{
				{
					center: Vec3{0, 0, 9},
					radius: 1,
					color:  RGB{255, 255, 255},
				},
			},
		},
		cam: Camera{
			origin: Vec3{0, 0, 0},
			up:     Vec3Up,
			fw:     Vec3Forward,
		},
	}
}

type retraceImageMsg struct{}

func RetraceImage() tea.Msg {
	return retraceImageMsg{}
}

type checkRedrawMsg time.Time

func checkRedrawCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*10, func(t time.Time) tea.Msg {
		return checkRedrawMsg(t)
	})
}

func (m Model) Init() tea.Cmd {
	return RetraceImage
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg, retraceImageMsg:
		return m, m.RetraceImage()
	case checkRedrawMsg:
		if m.render.progress < 1.0 {
			return m, checkRedrawCmd()
		}
	}

	return m, nil
}

func (m *Model) View() string {
	if m.render == nil {
		return ""
	}

	s := ""
	s += fmt.Sprintf("Size: %dx%d (%d pixels). Progress: %.2f%%\n", m.render.img.Rect.Dx(), m.render.img.Rect.Dy(), m.render.img.Rect.Dx()*m.render.img.Rect.Dy(), m.render.progress*100)

	for y := range m.render.img.Bounds().Dy() {
		for x := range m.render.img.Bounds().Dx() {
			s += fmt.Sprintf(
				"%s%sm",
				termenv.CSI,
				termenv.TrueColor.FromColor(m.render.img.At(x, y)).Sequence(true),
			)
			s += " "
		}
		s += fmt.Sprintf(
			"%s%sm ",
			termenv.CSI,
			termenv.ResetSeq,
		)
		s += "\n"
	}
	return s
}

func (m *Model) RetraceImage() tea.Cmd {

	// TODO: if there is an ongoing tracing, interrupt it. Current approach
	// replaces ongoing tracing, but it keeps going in the background until
	// it's done, which is a waste of resources

	width, height, err := term.GetSize(0)
	if err != nil {
		log.Fatalln("Error getting terminal size: ", err)
	}
	height -= 2

	render := Render{
		img:      image.NewRGBA(image.Rect(0, 0, width, height)),
		progress: 0,
	}
	m.render = &render

	go renderScene(m.render, &m.scene, m.cam)
	return checkRedrawCmd()
}

func main() {
	m := initialModel()
	p := tea.NewProgram(m)
	_, err := p.Run()
	if err != nil {
		log.Fatalf("Alas, there's been an error: %v", err)
	}
}
