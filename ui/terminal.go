package ui

import (
	"github.com/gdamore/tcell"
	"github.com/anydef/gol-tdd/engine"
)

type Terminal interface {
	Start() error
	Stop()
	Draw(engine.Coordinate)
	DrawGrid([]engine.Coordinate)
}

func (t TerminalImpl) DrawGrid(grid []engine.Coordinate) {
	for _, c := range grid {
		t.Draw(c)
	}
}

type TerminalImpl struct {
	screen tcell.Screen
}

func (t TerminalImpl) Draw(coord engine.Coordinate) {
	style := tcell.StyleDefault.Background(tcell.NewHexColor(0xffffff))
	t.screen.SetCell(coord.X, coord.Y, style, ' ')
}

func (t TerminalImpl) Start() error {
	return t.screen.Init()
}
func (t TerminalImpl) Stop() {
	t.screen.Fini()
}

func NewScreen() (Terminal, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	return TerminalImpl{screen: screen}, nil
}
