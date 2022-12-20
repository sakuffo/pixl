package apptype

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type BrushType = int

type PxCanvasConfig struct {
	DrawingArea fyne.Size
	CanvasOffset fyne.Position
	PxRows, PxCols int
	PxSize int
}

// This looks like all the state of the application that should be tracked

type State struct {
	BrushColor color.Color
	BrushType int
	SwatchSelected int
	FilePath string
}

func (state *State) SetFilePath(path string) {
	state.FilePath = path
}