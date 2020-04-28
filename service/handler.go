package service

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	// Runes for drawing the box
	Vertical    = '\u2551'
	Horizontal  = '\u2550'
	TopLeft     = '\u2554'
	TopRight    = '\u2557'
	BottomLeft  = '\u255A'
	BottomRight = '\u255D'

	// defaults for the canvas
	Width 		= 30
	Length 		= 30
	StartX		= 15
	StartY		= 15
)

type Handler struct {
	canvas    [Length][Width]bool
	direction int
	mode      string
	x         int
	y         int
}

func NewHandler() *Handler {
	handler := &Handler{}
	handler.x = StartX
	handler.y = StartY
	handler.mode = "draw"
	return handler
}

func (h *Handler) Handle(cmd string) string {
	if strings.HasPrefix(cmd, "steps") {
		return h.steps(getN(cmd, "steps"))
	}
	if strings.HasPrefix(cmd, "right") {
		return h.right(getN(cmd, "right"))
	}
	if strings.HasPrefix(cmd, "left") {
		return h.left(getN(cmd, "left"))
	}
	switch cmd {
	case "hover": h.mode = "hover"
	case "draw": h.mode = "draw"
	case "eraser": h.mode = "eraser"
	case "render": return h.render()
	case "coord": return h.coord()
	case "clear": h.canvas = [Length][Width]bool{}
	}
	return ""
}

func getN(s string, prefix string) int {
	s = strings.TrimPrefix(s, prefix)
	if len(s) == 0 {
		return 1
	} else {
		n, _ := strconv.Atoi(strings.TrimSpace(s))
		if n > 0 { return n } else { return 0 }
	}
}

func (h *Handler) right(n int) string {
	h.direction = (h.direction + n) % 8
	return ""
}

func (h *Handler) left(n int) string {
	h.direction = (h.direction + 8 - n % 8) % 8
	return ""
}

func (h *Handler) steps(n int) string {
	for i := 0; i < n; i++ {
		if h.mode == "draw" { h.canvas[h.y][h.x] = true }
		if h.mode == "eraser" { h.canvas[h.y][h.x] = false }
		switch h.direction {
		case 0: if h.y > 0 { h.y-- }
		case 1: if h.y > 0 && h.x < Width - 1 { h.y--; h.x++ }
		case 2: if h.x < Width - 1 { h.x++ }
		case 3: if h.x < Width - 1 && h.y < Length - 1 { h.x++; h.y++ }
		case 4: if h.y < Length - 1 { h.y++ }
		case 5: if h.y < Length - 1 && h.x > 0 { h.y++; h.x-- }
		case 6: if h.x > 0 { h.x-- }
		case 7: if h.x > 0 && h.y > 0 { h.x--; h.y-- }
		}
	}
	return ""
}

func (h *Handler) render() string {
	var b strings.Builder
	b.WriteRune(TopLeft)
	b.WriteString( strings.Repeat(string(Horizontal), Width))
	b.WriteRune(TopRight)
	b.WriteString("\r\n")
	for _, row := range h.canvas {
		b.WriteRune(Vertical)
		for _, c := range row {
			if c { b.WriteRune('*') } else { b.WriteRune(' ') }
		}
		b.WriteRune(Vertical)
		b.WriteString("\r\n")
	}
	b.WriteRune(BottomLeft)
	b.WriteString( strings.Repeat(string(Horizontal), Width))
	b.WriteRune(BottomRight)
	b.WriteString("\r\n\r\n")
	return b.String()
}

func (h *Handler) coord() string {
	return fmt.Sprintf("(%d,%d)\r\n", h.x, h.y)
}
