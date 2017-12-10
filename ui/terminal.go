package ui

import (
	"github.com/gdamore/tcell"
	"github.com/anydef/gol-tdd/engine"
)

type Terminal interface {
	Start() error
	Stop()
	Draw(engine.Coordinate)
	RedrawGrid([]engine.Coordinate)
	Poll() tcell.Event
	Size() (int, int)
	Sync()
	Clear()
	SetStyle(tcell.Style)
}

func (t TerminalImpl) SetStyle(style tcell.Style) {
	t.Screen.SetStyle(style)
}

func (t TerminalImpl) Clear() {
	t.Screen.Clear()
}

type TerminalImpl struct {
	Screen tcell.Screen
}

func (t TerminalImpl) Draw(coord engine.Coordinate) {
	style := tcell.StyleDefault.
		Background(tcell.NewHexColor(0x000000)).
		Foreground(tcell.NewHexColor(0xffffff))
	t.Screen.SetCell(coord.X, coord.Y, style, 'A')
}

func (t TerminalImpl) RedrawGrid(grid []engine.Coordinate) {
	for _, c := range grid {
		t.Draw(c)
	}
	t.Screen.Show()
}

func (t TerminalImpl) Start() error {
	return t.Screen.Init()
}

func (t TerminalImpl) Poll() tcell.Event {
	return t.Screen.PollEvent()
}

func (t TerminalImpl) Stop() {
	t.Screen.Fini()
}

func (t TerminalImpl) Size() (int, int) {
	width, height := t.Screen.Size()
	return width, height
}
func (t TerminalImpl) Sync() {
	//t.Screen.Sync()
	t.Screen.Show()
}
