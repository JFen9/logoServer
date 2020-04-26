package service

import (
	"fmt"
	"strings"
)

const (
	topLeft = '\u2552'
	horizontal = '\u2550'
	topRight = '\u2557'
	vertical = '\u2551'
	bottomLeft = '\u255A'
	bottomRight = '\u255B'
)

type Handler struct {
	canvas 		[30][30]bool
	direction	int
	x 			int
	y 			int
}

var ValidCommands = map[string]bool{
	"hover": true,
	"draw": true,
	"eraser": true,
	"coord": true,
	"render": true,
	"clear": true,
}

func (h Handler) Handle(cmd string) string {
	switch cmd {
	case "render": return h.render()
	case "coord": return h.coord()
	}
	return ""
}

func (h Handler) render() string {
	var b strings.Builder
	b.WriteRune(topLeft)
	b.WriteString( strings.Repeat(string(horizontal), 30))
	b.WriteRune(topRight)
	b.WriteRune('\n')
	for _, row := range h.canvas {
		b.WriteRune(vertical)
		for _, c := range row {
			if c {
				b.WriteRune('*')
			} else {
				b.WriteRune(' ')
			}
		}
		b.WriteRune(vertical)
		b.WriteRune('\n')
	}
	b.WriteRune(bottomLeft)
	b.WriteString( strings.Repeat(string(horizontal), 30))
	b.WriteRune(bottomRight)
	b.WriteRune('\n')
	return b.String()
}

func (h Handler) coord() string {
	return fmt.Sprintf("(%d,%d)\r\n", h.x, h.y)
}

func NewHandler() *Handler {
	handler := &Handler{}
	handler.x = 15
	handler.y = 15
	return handler
}