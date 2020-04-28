package service

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)


func TestNewHandler(t *testing.T) {
	handler := NewHandler()
	assert.Equal(t, handler.mode, "draw")
	assert.Zero(t, handler.direction)
	assert.Equal(t, len(handler.canvas), 30)
	assert.Equal(t, len(handler.canvas[0]), 30)
}

func TestDrawingSwitchingModes(t *testing.T) {
	handler := NewHandler()
	assert.Equal(t, "(15,15)", strings.TrimSpace(handler.Handle("coord")))
	handler.Handle("steps")
	assert.True(t, handler.canvas[15][15])
	assert.Equal(t, "(15,14)", strings.TrimSpace(handler.Handle("coord")))
	handler.Handle("hover")
	handler.Handle("steps 2")
	assert.False(t, handler.canvas[14][15])
	assert.False(t, handler.canvas[13][15])
	assert.Equal(t, "(15,12)", strings.TrimSpace(handler.Handle("coord")))
	handler.Handle("draw")
	handler.Handle("steps 3")
	assert.True(t, handler.canvas[12][15])
	assert.True(t, handler.canvas[11][15])
	assert.True(t, handler.canvas[10][15])
	assert.Equal(t, "(15,9)", strings.TrimSpace(handler.Handle("coord")))
	handler.Handle("right 4")
	assert.Equal(t, handler.direction, 4)
	handler.Handle("eraser")
	assert.Equal(t, handler.mode, "eraser")
	handler.Handle("steps 7")
	assert.False(t, handler.canvas[15][15])
	assert.False(t, handler.canvas[14][15])
	assert.False(t, handler.canvas[13][15])
	assert.False(t, handler.canvas[12][15])
	assert.False(t, handler.canvas[11][15])
	assert.False(t, handler.canvas[10][15])
	assert.Equal(t, "(15,16)", strings.TrimSpace(handler.Handle("coord")))
}

func TestChangeDirection(t *testing.T) {
	handler := NewHandler()
	handler.Handle("right 7")
	assert.Equal(t, handler.direction, 7)
	handler.Handle("right")
	assert.Equal(t, handler.direction, 0)
	handler.Handle("right 10")
	assert.Equal(t, handler.direction, 2)
	handler.Handle("left")
	assert.Equal(t, handler.direction, 1)
	handler.Handle("left 9")
	assert.Equal(t, handler.direction, 0)
	handler.Handle("left 2")
	assert.Equal(t, handler.direction, 6)
	handler.Handle("steps 2")
	assert.Equal(t, "(13,15)", strings.TrimSpace(handler.Handle("coord")))
}

func TestClear(t *testing.T) {
	handler := NewHandler()
	handlerNew := NewHandler()
	handler.Handle("steps 10")
	handler.Handle("right 2")
	handler.Handle("steps 5")
	assert.Equal(t, "(20,5)", strings.TrimSpace(handler.Handle("coord")))
	assert.NotEqual(t, handler.render(), handlerNew.render())
	handler.Handle("clear")
	assert.Equal(t, handler.render(), handlerNew.render())
	assert.Equal(t, "(20,5)", strings.TrimSpace(handler.Handle("coord")))
	assert.Equal(t, handler.direction, 2)
}