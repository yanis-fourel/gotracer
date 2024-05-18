package main

import (
	"fmt"
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
)

type Render = struct {
	img RenderImg
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
			ambientLight:    RGB{100, 100, 100},
			backgroundColor: RGB{},
			dirLight: DirLight{
				dir: Vec3{-2, -1, 2}.Normalized(),
				col: RGB{255, 255, 255},
			},
			spheres: []Sphere{
				{
					center: Vec3{-1.5, 1.2, 7},
					radius: 1,
					mat: Material{
						color:      RGB{255, 63, 63},
						reflection: 0,
					},
				},
				{
					center: Vec3{1.5, 1.2, 5},
					radius: 1,
					mat: Material{
						color:      RGB{255, 63, 63},
						reflection: 0.2,
					},
				},
			},
			planes: []Plane{
				{
					point: Vec3{0, 0, 0},
					norm:  Vec3Up,
					mat: Material{
						color:      RGBFrom("#9090BA"),
						reflection: 0.7,
					},
				},
			},
		},
		cam: Camera{
			origin: Vec3{0, 1.2, -1},
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
	return fmt.Sprintf(
		"Size: %dx%d (%d pixels). Progress: %.2f%%\n%s",
		m.render.img.width,
		m.render.img.height,
		m.render.img.width*m.render.img.height,
		m.render.progress*100,
		m.render.img.String(),
	)
}

func (m *Model) RetraceImage() tea.Cmd {

	// TODO: if there is an ongoing tracing, interrupt it. Current approach
	// replaces ongoing tracing, but it keeps going in the background until
	// it's done, which is a waste of resources and prevents from re-using
	// the same RenderImg

	width, height, err := term.GetSize(0)
	if err != nil {
		log.Fatalln("Error getting terminal size: ", err)
	}
	height -= 2

	render := Render{
		img:      NewRenderImg(width, height, m.scene.backgroundColor),
		progress: 0,
	}
	m.render = &render

	go renderScene(m.render, &m.scene, m.cam)
	return checkRedrawCmd()
}

func main() {
	// img := NewRenderImg(10, 5, RGB{})
	// img.Set(5, 0, RGB{255, 0, 0})
	// img.Set(2, 3, RGB{0, 255, 0})
	// img.Set(9, 4, RGB{0, 0, 255})
	// img.Set(0, 0, RGB{255, 255, 255})
	// fmt.Println(img.String())

	m := initialModel()
	p := tea.NewProgram(m)
	_, err := p.Run()
	if err != nil {
		log.Fatalf("Alas, there's been an error: %v", err)
	}
}
