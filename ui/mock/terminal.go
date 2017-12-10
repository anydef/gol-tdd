package mock

import (
	"github.com/anydef/gol-tdd/ui"
	"github.com/anydef/gol-tdd/engine"
)

type TerminalMock struct{}

func (t TerminalMock) Start() error {
	return nil
}
func (t TerminalMock) Stop()                        {}
func (t TerminalMock) Draw(engine.Coordinate)       {}
func (t TerminalMock) DrawGrid([]engine.Coordinate) {}

func NewScreenMock() ui.Terminal {
	return TerminalMock{}
}
