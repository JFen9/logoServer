package service

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	topLeft = '\u2554'
	horizontal = '\u2550'
	topRight = '\u2557'
	vertical = '\u2551'
	bottomLeft = '\u255A'
	bottomRight = '\u255D'
)

type Handler struct {
	canvas 		[30][30]bool
	direction	int
	mode		string
	x 			int
	y 			int
}

func NewHandler() *Handler {
	handler := &Handler{}
	handler.x = 15
	handler.y = 15
	handler.mode = "draw"
	return handler
}

func (h *Handler) Handle(cmd string) string {
	if strings.HasPrefix(cmd, "steps") {
		return h.steps(getN(cmd, "steps "))
	}
	if strings.HasPrefix(cmd, "right") {
		return h.right(getN(cmd, "right "))
	}
	if strings.HasPrefix(cmd, "left") {
		return h.left(getN(cmd, "left "))
	}
	switch cmd {
	case "hover": h.mode = "hover"
	case "draw": h.mode = "draw"
	case "eraser": h.mode = "eraser"
	case "render": return h.render()
	case "coord": return h.coord()
	}
	return ""
}

func getN(s string, prefix string) int {
	s = strings.TrimLeft(s, prefix)
	if len(s) == 0 {
		return 1
	} else {
		n, _ := strconv.Atoi(s)
		return n
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
		if h.mode == "draw" { h.canvas[h.x][h.y] = true }
		if h.mode == "eraser" { h.canvas[h.x][h.y] = false }
		switch h.direction {
		case 0: if h.x > 0 { h.x-- }
		case 1: if h.x > 0 && h.y < 29 { h.x--; h.y++ }
		case 2: if h.y < 29 { h.y++ }
		case 3: if h.y < 29 && h.x < 29 { h.y++; h.x++ }
		case 4: if h.x < 29 { h.x++ }
		case 5: if h.x < 29 && h.y > 0 { h.x++; h.y-- }
		case 6: if h.y > 0 { h.y-- }
		case 7: if h.y > 0 && h.x > 0 { h.y--; h.x-- }
		}
	}
	return ""
}

func (h Handler) render() string {
	var b strings.Builder
	b.WriteRune(topLeft)
	b.WriteString( strings.Repeat(string(horizontal), 30))
	b.WriteRune(topRight)
	b.WriteString("\r\n")
	for _, row := range h.canvas {
		b.WriteRune(vertical)
		for _, c := range row {
			if c { b.WriteRune('*') } else { b.WriteRune(' ') }
		}
		b.WriteRune(vertical)
		b.WriteString("\r\n")
	}
	b.WriteRune(bottomLeft)
	b.WriteString( strings.Repeat(string(horizontal), 30))
	b.WriteRune(bottomRight)
	b.WriteString("\r\n\r\n")
	return b.String()
}

func (h Handler) coord() string {
	return fmt.Sprintf("(%d,%d)\r\n", h.x, h.y)
}
